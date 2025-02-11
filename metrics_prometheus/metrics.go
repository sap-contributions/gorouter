package metrics_prometheus

import (
	mr "code.cloudfoundry.org/go-metric-registry"
	"code.cloudfoundry.org/gorouter/config"
	"code.cloudfoundry.org/gorouter/metrics"
	"code.cloudfoundry.org/gorouter/route"
	"log"
	"net/http"
	"time"
)

// Metrics represents a prometheus metrics endpoint.
type Metrics struct {
	RouteRegistration           mr.CounterVec
	RouteUnregistration         mr.CounterVec
	RoutesPruned                mr.Counter
	TotalRoutes                 mr.Gauge
	TimeSinceLastRegistryUpdate mr.Gauge
	RouteLookupTime             mr.Histogram
	RouteRegistrationLatency    mr.Histogram
	BadRequest                  mr.Counter
	BadGateway                  mr.Counter // TODO from b1tamara: rename to BackendBadGateway?
	BackendInvalidID            mr.Counter
	BackendInvalidTLSCert       mr.Counter
	BackendTLSHandshakeFailed   mr.Counter
	BackendExhaustedConns       mr.Counter
	// lookup metrics
	// proxy round tripper metrics
	// reporter metrics
	perRequestMetricsReporting bool
	unmuzzled                  uint64
}

func NewMetricsRegistry(config config.PrometheusConfig) *mr.Registry {
	var metricsRegistry *mr.Registry
	if config.Port != 0 && config.CertPath != "" {
		metricsRegistry = mr.NewRegistry(log.Default(),
			// the server starts in background. Endpoint: 127.0.0.1:port/metrics
			mr.WithTLSServer(int(config.Port), config.CertPath, config.KeyPath, config.CAPath))
	} else {
		metricsRegistry = mr.NewRegistry(log.Default(),
			// the server starts in background. Endpoint: 127.0.0.1:port/metrics
			mr.WithServer(int(config.Port)))
	}
	return metricsRegistry
}

var _ interface {
	metrics.ProxyReporter
	metrics.RouteRegistryReporter
} = &Metrics{}

func NewMetrics(registry *mr.Registry, perRequestMetricsReporting bool) *Metrics {
	return &Metrics{
		RouteRegistration:           registry.NewCounterVec("registry_message", "number of route registration messages", []string{"component", "action"}),
		RouteUnregistration:         registry.NewCounterVec("unregistry_message", "number of unregister messages", []string{"component"}),
		RoutesPruned:                registry.NewCounter("routes_pruned", "number of pruned routes"),
		TotalRoutes:                 registry.NewGauge("total_routes", "number of total routes"),
		TimeSinceLastRegistryUpdate: registry.NewGauge("ms_since_last_registry_update", "time since last registry update in ms"),
		RouteLookupTime:             registry.NewHistogram("route_lookup_time", "route lookup time per request in ns", []float64{10_000, 20_000, 30_000, 40_000, 50_000, 60_000, 70_000, 80_000, 90_000, 100_000}),
		RouteRegistrationLatency:    registry.NewHistogram("route_registration_latency", "route registration latency in ns", []float64{0.2, 0.4, 0.6, 0.8, 1, 1.2, 1.4, 1.6, 1.8, 2}), // TODO: validate
		BadRequest:                  registry.NewCounter("rejected_requests", "number of rejected requests"),
		BadGateway:                  registry.NewCounter("bad_gateways", "number of bad gateway errors received from backends"),
		BackendInvalidID:            registry.NewCounter("backend_invalid_id", "number of bad backend id errors received from backends"),
		BackendInvalidTLSCert:       registry.NewCounter("backend_invalid_tls_cert", "number of tls certificate errors received from backends"),
		BackendTLSHandshakeFailed:   registry.NewCounter("backend_tls_handshake_failed", "number of backend handshake errors"),
		BackendExhaustedConns:       registry.NewCounter("backend_exhausted_conns", "number of errors related to backend connection limit reached"),
		perRequestMetricsReporting:  perRequestMetricsReporting,
	}
}

func (metrics *Metrics) CaptureRouteStats(totalRoutes int, msSinceLastUpdate int64) {
	metrics.TotalRoutes.Set(float64(totalRoutes))
	metrics.TimeSinceLastRegistryUpdate.Set(float64(msSinceLastUpdate))
}

func (metrics *Metrics) CaptureRegistryMessage(msg metrics.ComponentTagged, action string) {
	metrics.RouteRegistration.Add(1, []string{msg.Component(), action})
}

func (metrics *Metrics) CaptureUnregistryMessage(msg metrics.ComponentTagged) {
	metrics.RouteUnregistration.Add(1, []string{msg.Component()})
}

func (metrics *Metrics) CaptureRoutesPruned(routesPruned uint64) {
	metrics.RoutesPruned.Add(float64(routesPruned))
}

func (metrics *Metrics) CaptureTotalRoutes(totalRoutes int) {
	metrics.TotalRoutes.Set(float64(totalRoutes))
}

func (metrics *Metrics) CaptureTimeSinceLastRegistryUpdate(msSinceLastUpdate int64) {
	metrics.TimeSinceLastRegistryUpdate.Set(float64(msSinceLastUpdate))
}

func (metrics *Metrics) CaptureLookupTime(t time.Duration) {

	// TODO: a histogram would be better.
	metrics.RouteLookupTime.Observe(float64(t.Nanoseconds()))
}

func (metrics *Metrics) CaptureRouteRegistrationLatency(t time.Duration) {
	metrics.RouteRegistrationLatency.Observe(float64(t) / float64(time.Millisecond))
}

// TODO: explain
func (metrics *Metrics) UnmuzzleRouteRegistrationLatency() {} // needed to fulfil interface

func (metrics *Metrics) CaptureBackendExhaustedConns() {
	metrics.BackendExhaustedConns.Add(1)
}

func (metrics *Metrics) CaptureBadGateway() {
	metrics.BadGateway.Add(1)
}

func (metrics *Metrics) CaptureBackendInvalidID() {
	metrics.BackendInvalidID.Add(1)
}

func (metrics *Metrics) CaptureBackendInvalidTLSCert() {
	metrics.BackendInvalidTLSCert.Add(1)
}

func (metrics *Metrics) CaptureBackendTLSHandshakeFailed() {
	metrics.BackendTLSHandshakeFailed.Add(1)
}

func (metrics *Metrics) CaptureBadRequest() {

	metrics.BadRequest.Add(1)
}

func (metrics *Metrics) CaptureEmptyContentLengthHeader() {
}

// TODO: check if function is used at all
func (metrics *Metrics) CaptureRoutingRequest(b *route.Endpoint) {
}
func (metrics *Metrics) CaptureRoutingResponse(statusCode int) {
}

func (metrics *Metrics) CaptureRoutingResponseLatency(b *route.Endpoint, statusCode int, t time.Time, d time.Duration) {
}

func (metrics *Metrics) CaptureRouteServiceResponse(res *http.Response) {
}

func (metrics *Metrics) CaptureWebSocketUpdate() {
}

func (metrics *Metrics) CaptureWebSocketFailure() {
}
