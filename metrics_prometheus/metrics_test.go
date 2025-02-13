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
		AfterEach(func() {
			m.perRequestMetricsReporting = true
		})

		It("sends number of nats messages received from each component", func() {
			endpoint.Tags = map[string]string{}
			m.CaptureRegistryMessage(endpoint, route.UPDATED.String())
			expected := fmt.Sprintf("registry_message{action=\"%s\",component=\"\"} 1", route.UPDATED.String())
			Expect(getMetrics(r.Port())).To(ContainSubstring(expected))

			m.CaptureRegistryMessage(endpoint, route.UPDATED.String())
			expected = fmt.Sprintf("registry_message{action=\"%s\",component=\"\"} 2", route.UPDATED.String())
			Expect(getMetrics(r.Port())).To(ContainSubstring(expected))
		})

		It("sends number of nats messages received from each component", func() {
			endpoint.Tags = map[string]string{"component": "uaa"}
			m.CaptureRegistryMessage(endpoint, route.ADDED.String())
			expected := fmt.Sprintf("registry_message{action=\"%s\",component=\"uaa\"} 1", route.ADDED.String())
			Expect(getMetrics(r.Port())).To(ContainSubstring(expected))

			endpoint.Tags = map[string]string{"component": "route-emitter"}
			m.CaptureRegistryMessage(endpoint, route.ADDED.String())
			expected = fmt.Sprintf("registry_message{action=\"%s\",component=\"route-emitter\"} 1", route.ADDED.String())
			Expect(getMetrics(r.Port())).To(ContainSubstring(expected))
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
			m.CaptureLookupTime(time.Duration(95) * time.Microsecond)
			Expect(getMetrics(r.Port())).To(ContainSubstring("route_lookup_time_bucket{le=\"100000\"} 1"))

			m.perRequestMetricsReporting = false
			m.CaptureLookupTime(time.Duration(95) * time.Microsecond)
			Expect(getMetrics(r.Port())).To(ContainSubstring("route_lookup_time_bucket{le=\"100000\"} 1"))
		})

		It("increments the routes pruned metric", func() {
			m.CaptureRoutesPruned(50)
			Expect(getMetrics(r.Port())).To(ContainSubstring(`routes_pruned 50`))
		})

		Describe("captures route registration latency", func() {
			It("properly splits the latencies apart", func() {
				m.CaptureRouteRegistrationLatency(1234 * time.Microsecond)
				m.CaptureRouteRegistrationLatency(134 * time.Microsecond)

				Expect(getMetrics(r.Port())).To(ContainSubstring("route_registration_latency_bucket{le=\"1.4\"} 2"))
				Expect(getMetrics(r.Port())).To(ContainSubstring("route_registration_latency_bucket{le=\"0.2\"} 1"))
			})
		})
	})
	Context("sends backend errors metrics", func() {
		BeforeEach(func() {
			var perRequestMetricsReporting = true
			var config = config.PrometheusConfig{Port: 0}
			r = NewMetricsRegistry(config)
			m = NewMetrics(r, perRequestMetricsReporting)
		})
		AfterEach(func() {
			m.perRequestMetricsReporting = true
		})

		It("increments the bad gateway to backend metric", func() {
			m.CaptureBadGateway()
			Expect(getMetrics(r.Port())).To(ContainSubstring("bad_gateways 1"))

			m.CaptureBadGateway()
			Expect(getMetrics(r.Port())).To(ContainSubstring("bad_gateways 2"))
		})

		It("increments the backend invalid id metric", func() {
			m.CaptureBackendInvalidID()
			Expect(getMetrics(r.Port())).To(ContainSubstring("backend_invalid_id 1"))

			m.CaptureBackendInvalidID()
			Expect(getMetrics(r.Port())).To(ContainSubstring("backend_invalid_id 2"))
		})

		It("increments the backend invalid tls cert metric", func() {
			m.CaptureBackendInvalidTLSCert()
			Expect(getMetrics(r.Port())).To(ContainSubstring("backend_invalid_tls_cert 1"))

			m.CaptureBackendInvalidTLSCert()
			Expect(getMetrics(r.Port())).To(ContainSubstring("backend_invalid_tls_cert 2"))
		})

		It("increments the backend tls handshake failed metric", func() {
			m.CaptureBackendTLSHandshakeFailed()
			Expect(getMetrics(r.Port())).To(ContainSubstring("backend_tls_handshake_failed 1"))

			m.CaptureBackendTLSHandshakeFailed()
			Expect(getMetrics(r.Port())).To(ContainSubstring("backend_tls_handshake_failed 2"))
		})
	})
	Context("sends lookup error metrics", func() {
		BeforeEach(func() {
			var perRequestMetricsReporting = true
			var config = config.PrometheusConfig{Port: 0}
			r = NewMetricsRegistry(config)
			m = NewMetrics(r, perRequestMetricsReporting)
		})
		AfterEach(func() {
			m.perRequestMetricsReporting = true
		})

		It("increments the bad requests metric", func() {
			m.CaptureBadRequest()
			Expect(getMetrics(r.Port())).To(ContainSubstring("rejected_requests 1"))

			m.CaptureBadRequest()
			Expect(getMetrics(r.Port())).To(ContainSubstring("rejected_requests 2"))
		})

		It("increments the empty content length header metric", func() {
			m.CaptureEmptyContentLengthHeader()
			Expect(getMetrics(r.Port())).To(ContainSubstring("empty_content_length_header 1"))
		})

		It("increments the backend exhausted conns metric", func() {
			m.CaptureBackendExhaustedConns()
			Expect(getMetrics(r.Port())).To(ContainSubstring("backend_exhausted_conns 1"))

			m.CaptureBackendExhaustedConns()
			Expect(getMetrics(r.Port())).To(ContainSubstring("backend_exhausted_conns 2"))
		})
	})
	Context("websocket metrics", func() {
		It("increments the websocket upgrades metric", func() {
			m.CaptureWebSocketUpdate()
			Expect(getMetrics(r.Port())).To(ContainSubstring("websocket_upgrades 1"))
		})

		It("increments the websocket failures metric", func() {
			m.CaptureWebSocketFailure()
			Expect(getMetrics(r.Port())).To(ContainSubstring("websocket_failures 1"))
		})
	})
	Context("increments the round trip metrics", func() {
		var endpoint *route.Endpoint

		BeforeEach(func() {
			var perRequestMetricsReporting = true
			var config = config.PrometheusConfig{Port: 0}
			r = NewMetricsRegistry(config)
			m = NewMetrics(r, perRequestMetricsReporting)
			endpoint = new(route.Endpoint)
		})
		AfterEach(func() {
			m.perRequestMetricsReporting = true
		})

		It("increments the total requests metric", func() {
			endpoint.Tags = map[string]string{}
			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"\",is_routed_app=\"no\"} 1"))

			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"\",is_routed_app=\"no\"} 2"))
		})

		It("increments the requests metric for the given component", func() {
			endpoint.Tags = map[string]string{"component": "CloudController"}
			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"CloudController\",is_routed_app=\"no\"} 1"))

			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"CloudController\",is_routed_app=\"no\"} 2"))

			endpoint.Tags["component"] = "UAA"
			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"CloudController\",is_routed_app=\"no\"} 2"))
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"UAA\",is_routed_app=\"no\"} 1"))

			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"UAA\",is_routed_app=\"no\"} 2"))
		})

		It("increments the request metric for the routed app", func() {
			endpoint.Tags = map[string]string{"component": "dea-1"}
			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"dea-*\",is_routed_app=\"yes\"} 1"))

			endpoint.Tags["component"] = "dea-3"
			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"dea-*\",is_routed_app=\"yes\"} 2"))

			endpoint.Tags["component"] = "CustomController"
			m.CaptureRoutingRequest(endpoint)
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"dea-*\",is_routed_app=\"yes\"} 2"))
			Expect(getMetrics(r.Port())).To(ContainSubstring("total_requests{component=\"CustomController\",is_routed_app=\"no\"} 1"))
		})
	})

	Context("increments the response metrics", func() {
		It("increments the 2XX response metrics", func() {
			m.CaptureRoutingResponse(200)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"2xx\"} 1"))

			m.CaptureRoutingResponse(200)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"2xx\"} 2"))
		})

		It("increments the 3XX response metrics", func() {
			m.CaptureRoutingResponse(304)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"3xx\"} 1"))

			m.CaptureRoutingResponse(300)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"3xx\"} 2"))
		})

		It("increments the 4XX response metrics", func() {
			m.CaptureRoutingResponse(401)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"4xx\"} 1"))

			m.CaptureRoutingResponse(401)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"4xx\"} 2"))
		})

		It("increments the 5XX response metrics", func() {
			m.CaptureRoutingResponse(500)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"5xx\"} 1"))

			m.CaptureRoutingResponse(504)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"5xx\"} 2"))
		})

		It("increments the XXX response metrics", func() {
			m.CaptureRoutingResponse(100)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"xxx\"} 1"))

			m.CaptureRoutingResponse(100)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"xxx\"} 2"))
		})

		It("increments the XXX response metrics with null response", func() {
			m.CaptureRoutingResponse(0)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"xxx\"} 3")) // from the prev It node

			m.CaptureRoutingResponse(0)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses{status_group=\"xxx\"} 4"))
		})
	})

	Context("increments the response metrics for route services", func() {
		var response http.Response

		BeforeEach(func() {
			var perRequestMetricsReporting = true
			var config = config.PrometheusConfig{Port: 0}
			r = NewMetricsRegistry(config)
			m = NewMetrics(r, perRequestMetricsReporting)
			response = http.Response{}
		})
		AfterEach(func() {
			m.perRequestMetricsReporting = true
		})

		It("increments the 2XX route services response metrics", func() {
			response.StatusCode = 200
			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"2xx\"} 1"))

			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"2xx\"} 2"))
		})

		It("increments the 3XX response metrics", func() {
			response.StatusCode = 300
			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"3xx\"} 1"))

			response.StatusCode = 304
			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"3xx\"} 2"))
		})

		It("increments the 4XX response metrics", func() {
			response.StatusCode = 401
			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"4xx\"} 1"))

			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"4xx\"} 2"))
		})

		It("increments the 5XX response metrics", func() {
			response.StatusCode = 500
			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"5xx\"} 1"))

			response.StatusCode = 504
			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"5xx\"} 2"))
		})

		It("increments the XXX response metrics", func() {
			response.StatusCode = 100
			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"xxx\"} 1"))

			m.CaptureRouteServiceResponse(&response)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"xxx\"} 2"))
		})

		It("increments the XXX response metrics with null response", func() {
			m.CaptureRouteServiceResponse(nil)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"xxx\"} 1"))

			m.CaptureRouteServiceResponse(nil)
			Expect(getMetrics(r.Port())).To(ContainSubstring("responses_route_services{status_group=\"xxx\"} 2"))
		})
	})

	Context("increments the route response metrics", func() {
		var endpoint *route.Endpoint

		BeforeEach(func() {
			var perRequestMetricsReporting = true
			var config = config.PrometheusConfig{Port: 0}
			r = NewMetricsRegistry(config)
			m = NewMetrics(r, perRequestMetricsReporting)
			endpoint = new(route.Endpoint)
		})
		AfterEach(func() {
			m.perRequestMetricsReporting = true
		})

		It("sends the latency", func() {
			m.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 2*time.Millisecond)
			m.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 500*time.Microsecond)
			Expect(getMetrics(r.Port())).To(ContainSubstring("latency_bucket{component=\"\",le=\"0.6\"} 1"))
			Expect(getMetrics(r.Port())).To(ContainSubstring("latency_bucket{component=\"\",le=\"2\"} 2"))
		})

		It("does not send the latency if switched off", func() {
			m.perRequestMetricsReporting = false
			m.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 2*time.Millisecond)
			Expect(getMetrics(r.Port())).NotTo(ContainSubstring("\nlatency_bucket"))
		})

		It("sends the latency for the given component", func() {
			endpoint.Tags = map[string]string{"component": "CloudController"}
			m.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 2*time.Millisecond)
			Expect(getMetrics(r.Port())).To(ContainSubstring("latency_bucket{component=\"CloudController\",le=\"2\"} 1"))
		})

		It("does not send the latency for the given component if switched off", func() {
			m.perRequestMetricsReporting = false
			endpoint.Tags = map[string]string{"component": "CloudController"}
			m.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 2*time.Millisecond)
			Expect(getMetrics(r.Port())).NotTo(ContainSubstring("\nlatency_bucket"))
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
