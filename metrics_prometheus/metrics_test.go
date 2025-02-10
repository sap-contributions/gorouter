package metrics_prometheus

import (
	metrics "code.cloudfoundry.org/go-metric-registry"
	"code.cloudfoundry.org/gorouter/config"
	"code.cloudfoundry.org/gorouter/route"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"time"
)

var m *Metrics
var r *metrics.Registry

var _ = Describe("Metrics", func() {

	Context("sends route metrics", func() {
		var endpoint *route.Endpoint

		BeforeEach(func() {
			var perRequestMetricsReporting = true
			var config = config.PrometheusConfig{Port: 0}
			r = NewMetricsRegistry(config)
			m = NewMetrics(r, perRequestMetricsReporting)
			endpoint = new(route.Endpoint)
		})

		It("sends number of nats messages received from each component", func() {
			endpoint.Tags = map[string]string{}
			m.CaptureRegistryMessage(endpoint)

			Expect(getMetrics(r.Port())).To(ContainSubstring("registry_message{component=\"\"} 1"))
		})

		It("sends number of nats messages received from each component", func() {
			endpoint.Tags = map[string]string{"component": "uaa"}
			m.CaptureRegistryMessage(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("registry_message{component=\"uaa\"} 1"))

			endpoint.Tags = map[string]string{"component": "route-emitter"}
			m.CaptureRegistryMessage(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("registry_message{component=\"route-emitter\"} 1"))
		})

		It("sends the total routes", func() {
			m.CaptureTotalRoutes(12)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_routes 12"))
		})

		It("sends the time since last update", func() {
			m.CaptureTimeSinceLastRegistryUpdate(int64(100))

			Expect(getMetrics(r.Port())).To(ContainSubstring("ms_since_last_registry_update 100"))
		})

		It("sends the lookup time for routing table", func() {
			m.CaptureLookupTime(time.Duration(955) * time.Millisecond)

			Expect(getMetrics(r.Port())).To(ContainSubstring("route_lookup_time 9.55e+08"))
		})

		It("increments the routes_pruned metric", func() {
			m.CaptureRoutesPruned(50)
			Expect(getMetrics(r.Port())).To(ContainSubstring(`routes_pruned 50`))
		})

		Describe("CaptureRouteRegistrationLatency", func() {
			It("properly splits the latencies apart", func() {
				m.CaptureRouteRegistrationLatency(1234 * time.Microsecond)
				m.CaptureRouteRegistrationLatency(134 * time.Microsecond)

				Expect(getMetrics(r.Port())).To(ContainSubstring("route_registration_latency_bucket{le=\"1.4\"} 2"))
				Expect(getMetrics(r.Port())).To(ContainSubstring("route_registration_latency_bucket{le=\"0.2\"} 1"))
			})
		})
	})
})

func getMetrics(port string) string {
	addr := fmt.Sprintf("http://127.0.0.1:%s/metrics", port)
	resp, err := http.Get(addr) //nolint:gosec
	if err != nil {
		return ""
	}

	respBytes, err := io.ReadAll(resp.Body)
	Expect(err).ToNot(HaveOccurred())

	return string(respBytes)
}
