package logger

import (
	"log/slog"
	"net/http"
	"strings"

	"code.cloudfoundry.org/lager/v3"

	"github.com/openzipkin/zipkin-go/idgenerator"
	"github.com/openzipkin/zipkin-go/model"
)

const (
	RequestIdHeader = "X-Vcap-Request-Id"
)

// LagerAdapter satisfies the lager.Logger interface with zap as the
// implementation.
type LagerAdapter struct {
	originalLogger *slog.Logger
	session        string
}

// NewLagerAdapter returns a new lager.Logger that uses zap underneath.
func NewLagerAdapter(slogLogger *slog.Logger) *LagerAdapter {
	return &LagerAdapter{
		originalLogger: slogLogger,
	}
}

// RegisterSink is never used after initialization, so it does nothing.
func (l *LagerAdapter) RegisterSink(_ lager.Sink) {}

// Session returns a new logger with a nested session.
func (l *LagerAdapter) Session(task string, data ...lager.Data) lager.Logger {
	tmpLogger := l.originalLogger.With("session", l.session+"."+task)

	if data != nil {
		tmpLogger = l.originalLogger.With(dataToFields(data)...)
	}

	return &LagerAdapter{
		originalLogger: tmpLogger,
	}
}

// SessionName returns the name of the logger session
func (l *LagerAdapter) SessionName() string {
	return l.session
}

// Debug logs a message at the debug log level.
func (l *LagerAdapter) Debug(action string, data ...lager.Data) {
	l.originalLogger.Debug(action, dataToFields(data)...)
}

// Info logs a message at the info log level.
func (l *LagerAdapter) Info(action string, data ...lager.Data) {
	l.originalLogger.Info(action, dataToFields(data)...)
}

// Error logs a message at the error log level.
func (l *LagerAdapter) Error(action string, err error, data ...lager.Data) {
	l.originalLogger.Error(action, append(dataToFields(data), ErrAttr(err))...)
}

// Fatal logs a message and exits with status 1.
func (l *LagerAdapter) Fatal(action string, err error, data ...lager.Data) {
	l.originalLogger.Error(action, append(dataToFields(data), ErrAttr(err))...)
}

// WithData returns a logger with newly added data.
func (l *LagerAdapter) WithData(data lager.Data) lager.Logger {
	return &LagerAdapter{
		originalLogger: l.originalLogger.With(dataToFields([]lager.Data{data})...),
	}
}

func (l *LagerAdapter) WithTraceInfo(req *http.Request) lager.Logger {
	traceIDHeader := req.Header.Get(RequestIdHeader)
	if traceIDHeader == "" {
		return l.WithData(nil)
	}
	traceHex := strings.Replace(traceIDHeader, "-", "", -1)
	traceID, err := model.TraceIDFromHex(traceHex)
	if err != nil {
		return l.WithData(nil)
	}

	spanID := idgenerator.NewRandom128().SpanID(traceID)
	return l.WithData(lager.Data{"trace-id": traceID.String(), "span-id": spanID.String()})
}

func dataToFields(data []lager.Data) []any {
	var fields []any
	for _, datum := range data {
		for key, value := range datum {
			fields = append(fields, slog.Any(key, value))
		}
	}
	return fields
}
