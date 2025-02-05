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
			var config = config.PrometheusConfig{Port: 0}
			r = NewMetricsRegistry(config)
			m = NewRouteRegistryMetrics(r, perRequestMetricsReporting)
			endpoint = new(route.Endpoint)
		})

		It("sends number of nats messages received from each component", func() {
			endpoint.Tags = map[string]string{}
			m.CaptureRegistryMessageWithLabel(endpoint.Component(), route.ADDED.String())

			Expect(getMetrics(r.Port())).To(ContainSubstring("registry_message{component_name=\"\",update_type=\"added\"}"))
		})

		It("sends number of nats messages received from each component", func() {
			endpoint.Tags = map[string]string{"component": "uaa"}
			m.CaptureRegistryMessageWithLabel(endpoint.Component(), route.UPDATED.String())
			Expect(getMetrics(r.Port())).To(ContainSubstring("registry_message{component_name=\"uaa\",update_type=\"updated\"} 1"))

			endpoint.Tags = map[string]string{"component": "route-emitter"}
			m.CaptureRegistryMessageWithLabel(endpoint.Component(), route.UPDATED.String())
			Expect(getMetrics(r.Port())).To(ContainSubstring("registry_message{component_name=\"route-emitter\",update_type=\"updated\"} 1"))

		})

		It("sends the total routes", func() {
			m.CaptureTotalRoutes(12)
			time.Sleep(100 * time.Millisecond)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_routes 12"))
		})

		It("sends the time since last update", func() {
			m.CaptureTimeSinceLastRegistryUpdate(int64(100))

			Expect(getMetrics(r.Port())).To(ContainSubstring("ms_since_last_registry_update 100"))
		})

		It("sends the lookup time for routing table", func() {
			m.CaptureLookupTime(time.Duration(955) * time.Millisecond)

			Expect(getMetrics(r.Port())).To(ContainSubstring("total_routes"))
		})
	})

	It("increments the routes_pruned metric", func() {
		m.CaptureRoutesPruned(50)
		Expect(getMetrics(r.Port())).To(ContainSubstring(`routes_pruned 50`))
	})

	Describe("CaptureRouteRegistrationLatency", func() {
		It("is muzzled by default", func() {
			m.CaptureRouteRegistrationLatency(2 * time.Second)

			Expect(getMetrics).To(ContainSubstring("route_registration_latency 2000"))
		})
		It("sends router registration latency when unmuzzled", func() {
			m.UnmuzzleRouteRegistrationLatency()
			m.CaptureRouteRegistrationLatency(2 * time.Second)
			Expect(getMetrics(r.Port())).To(ContainSubstring("route_registration_latency 2000"))
		})
	})

})

func getMetrics(port string) string {
	addr := fmt.Sprintf("http://127.0.0.1:%s/metrics", port)
	fmt.Print(addr)
	resp, err := http.Get(addr) //nolint:gosec
	if err != nil {
		return ""
	}

	respBytes, err := io.ReadAll(resp.Body)
	Expect(err).ToNot(HaveOccurred())

	return string(respBytes)
}
