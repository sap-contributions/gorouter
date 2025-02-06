package metrics_prometheus

import (
	mr "code.cloudfoundry.org/go-metric-registry"
	"code.cloudfoundry.org/gorouter/config"
	"code.cloudfoundry.org/gorouter/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"reflect"
	"sync/atomic"
	"time"
	"unsafe"
)

var _ metrics.RouteRegistryReporter = &Metrics{}

// Metrics represents a prometheus metrics endpoint.
type Metrics struct {
	Config                      config.PrometheusConfig
	Mux                         *http.ServeMux
	server                      *http.Server
	RouteRegistration           *prometheus.CounterVec
	RouteUnregistration         *prometheus.CounterVec
	RoutesPruned                prometheus.Counter
	TotalRoutes                 prometheus.Gauge
	TimeSinceLastRegistryUpdate prometheus.Gauge
	RouteLookupTime             prometheus.Gauge
	RouteRegistrationLatency    prometheus.Gauge
	BadRequest                  prometheus.Counter
	// lookup metrics
	// error handler metrics
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

func NewRouteRegistryMetrics(registry *mr.Registry, perRequestMetricsReporting bool) *Metrics {
	// The interface provided by the library massively restricts the usability of the prometheus
	// library. If we are to implement this we either need to contribute a lot to the library or
	// consume prometheus directly. Either way: what comes next is a no-go and has to be removed.

	promRegV := reflect.ValueOf(registry).Elem().FieldByName("registerer")
	promReg := reflect.NewAt(promRegV.Type(), unsafe.Pointer(promRegV.UnsafeAddr())).Interface().(*prometheus.Registerer)

	m := &Metrics{
		perRequestMetricsReporting: perRequestMetricsReporting,
		unmuzzled:                  uint64(1),
		RouteRegistration: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "registry_message",
			Help: "number of route registration messages",
		}, []string{"component_name"}),
		RouteUnregistration: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "unregistry_message",
			Help: "number of route unregister messages",
		}, []string{"update_type", "component_name"}),
		RoutesPruned: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "routes_pruned",
			Help: "number of pruned routes",
		}),
		TotalRoutes: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "total_routes",
			Help: "number of total routes",
		}),
		TimeSinceLastRegistryUpdate: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "ms_since_last_registry_update",
			Help: "Time since last registry update in ms",
		}),
		RouteLookupTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "route_lookup_time",
			Help: "Route lookup time per request in ns",
		}),
		RouteRegistrationLatency: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "route_registration_latency",
			Help: "Route registration latency in ms",
		}),
	}

	(*promReg).MustRegister(m.RouteRegistration)
	(*promReg).MustRegister(m.RouteUnregistration)
	(*promReg).MustRegister(m.RoutesPruned)
	(*promReg).MustRegister(m.TotalRoutes)
	(*promReg).MustRegister(m.TimeSinceLastRegistryUpdate)
	(*promReg).MustRegister(m.RouteLookupTime)
	(*promReg).MustRegister(m.RouteRegistrationLatency)

	return m
}

func (metrics *Metrics) CaptureRouteStats(totalRoutes int, msSinceLastUpdate int64) {
	if !metrics.isPrometheusEnabled() {
		return
	}
	metrics.TotalRoutes.Set(float64(totalRoutes))
	metrics.TimeSinceLastRegistryUpdate.Set(float64(msSinceLastUpdate))
}

func (metrics *Metrics) CaptureRegistryMessage(msg metrics.ComponentTagged) {
	if !metrics.isPrometheusEnabled() {
		return
	}

	metrics.RouteRegistration.WithLabelValues(msg.Component()).Inc()
}

func (metrics *Metrics) CaptureUnregistryMessage(msg metrics.ComponentTagged) {
	if !metrics.isPrometheusEnabled() {
		return
	}
	metrics.RouteUnregistration.WithLabelValues(msg.Component()).Inc()
}

func (metrics *Metrics) CaptureRoutesPruned(routesPruned uint64) {
	if !metrics.isPrometheusEnabled() {
		return
	}
	metrics.RoutesPruned.Add(float64(routesPruned))
}

func (metrics *Metrics) CaptureTotalRoutes(totalRoutes int) {
	if !metrics.isPrometheusEnabled() {
		return
	}
	metrics.TotalRoutes.Set(float64(totalRoutes))
}

func (metrics *Metrics) CaptureTimeSinceLastRegistryUpdate(msSinceLastUpdate int64) {
	if !metrics.isPrometheusEnabled() {
		return
	}
	metrics.TimeSinceLastRegistryUpdate.Set(float64(msSinceLastUpdate))
}

func (metrics *Metrics) CaptureLookupTime(t time.Duration) {
	if !metrics.isPrometheusEnabled() || !metrics.perRequestMetricsReporting {
		return
	}

	metrics.RouteLookupTime.Set(float64(t.Nanoseconds()))
}

func (metrics *Metrics) CaptureRouteRegistrationLatency(t time.Duration) {
	if !metrics.isPrometheusEnabled() {
		return
	}
	if atomic.LoadUint64(&metrics.unmuzzled) == 1 {
		latency := t / time.Millisecond
		metrics.RouteRegistrationLatency.Set(float64(latency))
	}
}

func (metrics *Metrics) UnmuzzleRouteRegistrationLatency() {
	if !metrics.isPrometheusEnabled() {
		return
	}
	atomic.StoreUint64(&metrics.unmuzzled, 1)
}

func (metrics *Metrics) CaptureBadRequest() {
	if !metrics.isPrometheusEnabled() {
		return
	}

	metrics.BadRequest.Inc()
}

func (metrics *Metrics) isPrometheusEnabled() bool {
	return true
}
