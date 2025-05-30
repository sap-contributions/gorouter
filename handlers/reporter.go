package handlers

import (
	"log/slog"
	"net/http"
	"net/textproto"
	"time"

	"github.com/urfave/negroni/v3"

	log "code.cloudfoundry.org/gorouter/logger"
	"code.cloudfoundry.org/gorouter/metrics"
	"code.cloudfoundry.org/gorouter/proxy/utils"
)

type reporterHandler struct {
	reporter metrics.MetricReporter
	logger   *slog.Logger
}

// NewReporter creates a new handler that handles reporting backend
// responses to metrics and missing Content-Length header
func NewReporter(reporter metrics.MetricReporter, logger *slog.Logger) negroni.Handler {
	return &reporterHandler{
		reporter: reporter,
		logger:   logger,
	}
}

// ServeHTTP handles reporting the response after the request has been completed
func (rh *reporterHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := LoggerWithTraceInfo(rh.logger, r)
	requestInfo, err := ContextRequestInfo(r)
	// logger.Panic does not cause gorouter to exit 1 but rather throw panic with
	// stacktrace in error log
	if err != nil {
		log.Panic(logger, "request-info-err", log.ErrAttr(err))
		return
	}
	if !validContentLength(r.Header) {
		rh.reporter.CaptureEmptyContentLengthHeader()
	}

	next(rw, r)

	requestInfo.FinishedAt = time.Now()
	if requestInfo.RouteEndpoint == nil {
		return
	}

	proxyWriter := rw.(utils.ProxyResponseWriter)
	rh.reporter.CaptureRoutingResponse(proxyWriter.Status())

	if requestInfo.AppRequestFinishedAt.Equal(time.Time{}) {
		return
	}
	rh.reporter.CaptureRoutingResponseLatency(
		requestInfo.RouteEndpoint, proxyWriter.Status(),
		requestInfo.ReceivedAt, requestInfo.AppRequestFinishedAt.Sub(requestInfo.ReceivedAt),
	)
	rh.calculateGorouterTime(requestInfo)
	rh.reporter.CaptureGorouterTime(requestInfo.GorouterTime)
}

// calculateGorouterTime
// calculate the gorouter time by subtracting app response time from the total roundtrip time.
// Parameters:
//   - requestInfo *RequestInfo
func (rh *reporterHandler) calculateGorouterTime(requestInfo *RequestInfo) {
	requestInfo.GorouterTime = -1
	appTime := requestInfo.AppRequestFinishedAt.Sub(requestInfo.AppRequestStartedAt).Seconds()
	rtTime := requestInfo.FinishedAt.Sub(requestInfo.ReceivedAt).Seconds()
	if rtTime >= 0 && appTime >= 0 {
		requestInfo.GorouterTime = rtTime - appTime
	}
}

// validContentLength ensures that if the `Content-Length` header is set, it is not empty.
// Request that don't have a `Content-Length` header are OK.
//
// Based on https://github.com/golang/go/blob/33496c2dd310aad1d56bae9febcbd2f02b4985cb/src/net/http/transfer.go#L1051
// http.Header.Get() will return "" for empty headers, or when the header is not set at all.
func validContentLength(header http.Header) bool {
	clHeaders := header["Content-Length"]

	if len(clHeaders) == 0 {
		return true
	}
	cl := textproto.TrimString(clHeaders[0])

	// The Content-Length must be a valid numeric value.
	// See: https://datatracker.ietf.org/doc/html/rfc2616/#section-14.13
	return cl != ""
}
