package metrics_test

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/urfave/negroni/v3"
	"go.uber.org/zap/zapcore"

	"code.cloudfoundry.org/gorouter/config"
	"code.cloudfoundry.org/gorouter/handlers"
	log "code.cloudfoundry.org/gorouter/logger"
	"code.cloudfoundry.org/gorouter/metrics"
	"code.cloudfoundry.org/gorouter/metrics/fakes"
	"code.cloudfoundry.org/gorouter/route"
	"code.cloudfoundry.org/gorouter/test_util"
)

var _ = Describe("MetricsReporter", func() {
	var (
		endpoint       *route.Endpoint
		sender         *fakes.MetricSender
		batcher        *fakes.MetricBatcher
		metricReporter *metrics.Metrics
	)

	BeforeEach(func() {
		endpoint = route.NewEndpoint(&route.EndpointOpts{Tags: map[string]string{}})
		sender = new(fakes.MetricSender)
		batcher = new(fakes.MetricBatcher)
		cfg, err := config.DefaultConfig()
		Expect(err).ToNot(HaveOccurred())
		metricReporter = &metrics.Metrics{Sender: sender, Batcher: batcher, PerRequestMetricsReporting: cfg.PerRequestMetricsReporting}
	})

	It("increments the bad_requests metric", func() {
		metricReporter.CaptureBadRequest()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
		Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("rejected_requests"))

		metricReporter.CaptureBadRequest()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
		Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("rejected_requests"))
	})

	It("increments the bad_gateway metric", func() {
		metricReporter.CaptureBadGateway()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
		Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("bad_gateways"))

		metricReporter.CaptureBadGateway()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
		Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("bad_gateways"))
	})

	It("increments the backend_exhausted_conns metric", func() {
		metricReporter.CaptureBackendExhaustedConns()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
		Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("backend_exhausted_conns"))

		metricReporter.CaptureBackendExhaustedConns()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
		Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("backend_exhausted_conns"))
	})

	It("increments the backend_invalid_id metric", func() {
		metricReporter.CaptureBackendInvalidID()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
		Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("backend_invalid_id"))

		metricReporter.CaptureBackendInvalidID()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
		Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("backend_invalid_id"))
	})

	It("increments the backend_invalid_tls_cert metric", func() {
		metricReporter.CaptureBackendInvalidTLSCert()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
		Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("backend_invalid_tls_cert"))

		metricReporter.CaptureBackendInvalidTLSCert()

		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
		Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("backend_invalid_tls_cert"))
	})

	Context("increments the request metrics", func() {
		It("increments the total requests metric", func() {
			metricReporter.CaptureRoutingRequest(&route.Endpoint{})

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("total_requests"))

			metricReporter.CaptureRoutingRequest(&route.Endpoint{})

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("total_requests"))
		})

		It("increments the requests metric for the given component", func() {
			endpoint.Tags["component"] = "CloudController"
			metricReporter.CaptureRoutingRequest(endpoint)

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("requests.CloudController"))

			metricReporter.CaptureRoutingRequest(endpoint)

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(3)).To(Equal("requests.CloudController"))

			endpoint.Tags["component"] = "UAA"
			metricReporter.CaptureRoutingRequest(endpoint)

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(6))
			Expect(batcher.BatchIncrementCounterArgsForCall(5)).To(Equal("requests.UAA"))

			metricReporter.CaptureRoutingRequest(endpoint)

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(8))
			Expect(batcher.BatchIncrementCounterArgsForCall(7)).To(Equal("requests.UAA"))
		})

		It("increments the routed_app_requests metric", func() {
			endpoint.Tags["component"] = "dea-1"
			metricReporter.CaptureRoutingRequest(endpoint)

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(3))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("routed_app_requests"))

			endpoint.Tags["component"] = "dea-3"
			metricReporter.CaptureRoutingRequest(endpoint)

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(6))
			Expect(batcher.BatchIncrementCounterArgsForCall(5)).To(Equal("routed_app_requests"))

			endpoint.Tags["component"] = "CloudController"
			metricReporter.CaptureRoutingRequest(endpoint)

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(8))
		})
	})

	Context("increments the response metrics for route services", func() {
		It("increments the 2XX route services response metrics", func() {
			response := http.Response{
				StatusCode: 200,
			}

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.route_services.2xx"))

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.route_services.2xx"))
		})

		It("increments the 3XX response metrics", func() {
			response := http.Response{
				StatusCode: 304,
			}

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.route_services.3xx"))

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.route_services.3xx"))
		})

		It("increments the 4XX response metrics", func() {
			response := http.Response{
				StatusCode: 401,
			}

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.route_services.4xx"))

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.route_services.4xx"))
		})

		It("increments the 5XX response metrics", func() {
			response := http.Response{
				StatusCode: 504,
			}

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.route_services.5xx"))

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.route_services.5xx"))
		})

		It("increments the XXX response metrics", func() {
			response := http.Response{
				StatusCode: 100,
			}

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.route_services.xxx"))

			metricReporter.CaptureRouteServiceResponse(&response)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.route_services.xxx"))
		})

		It("increments the XXX response metrics with null response", func() {
			metricReporter.CaptureRouteServiceResponse(nil)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.route_services.xxx"))

			metricReporter.CaptureRouteServiceResponse(nil)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.route_services.xxx"))
		})

		It("increments the total responses", func() {
			response2xx := http.Response{
				StatusCode: 205,
			}
			response4xx := http.Response{
				StatusCode: 401,
			}

			metricReporter.CaptureRouteServiceResponse(&response2xx)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("responses.route_services"))

			metricReporter.CaptureRouteServiceResponse(&response4xx)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(3)).To(Equal("responses.route_services"))
		})
	})

	Context("increments the response metrics", func() {
		It("increments the 2XX response metrics", func() {
			metricReporter.CaptureRoutingResponse(200)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.2xx"))

			metricReporter.CaptureRoutingResponse(200)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.2xx"))
		})

		It("increments the 3XX response metrics", func() {
			metricReporter.CaptureRoutingResponse(304)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.3xx"))

			metricReporter.CaptureRoutingResponse(304)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.3xx"))
		})

		It("increments the 4XX response metrics", func() {
			metricReporter.CaptureRoutingResponse(401)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.4xx"))

			metricReporter.CaptureRoutingResponse(401)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.4xx"))
		})

		It("increments the 5XX response metrics", func() {
			metricReporter.CaptureRoutingResponse(504)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.5xx"))

			metricReporter.CaptureRoutingResponse(504)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.5xx"))
		})

		It("increments the XXX response metrics", func() {
			metricReporter.CaptureRoutingResponse(100)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.xxx"))

			metricReporter.CaptureRoutingResponse(100)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.xxx"))
		})

		It("increments the XXX response metrics with null response", func() {
			metricReporter.CaptureRoutingResponse(0)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("responses.xxx"))

			metricReporter.CaptureRoutingResponse(0)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(2)).To(Equal("responses.xxx"))
		})

		It("increments the total responses", func() {
			metricReporter.CaptureRoutingResponse(205)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("responses"))

			metricReporter.CaptureRoutingResponse(401)
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(4))
			Expect(batcher.BatchIncrementCounterArgsForCall(3)).To(Equal("responses"))
		})
	})

	Context("metric empty_content_length_header", func() {
		var (
			testApp  *httptest.Server
			godebug  string
			testSink *test_util.TestSink
			logger   *slog.Logger
		)

		BeforeEach(func() {
			// Ensure we always have httplaxcontentlength=1 set for this test.
			// When httplaxcontentlength=1. is no longer a thing, we should consider
			// removing this test and the metric logic it relates to
			godebug = os.Getenv("GODEBUG")
			os.Setenv("GODEBUG", fmt.Sprintf("%s,httplaxcontentlength=1", godebug))
			logger = log.CreateLogger()
			testSink = &test_util.TestSink{Buffer: gbytes.NewBuffer()}
			log.SetDynamicWriteSyncer(zapcore.NewMultiWriteSyncer(testSink, zapcore.AddSync(GinkgoWriter)))
			log.SetLoggingLevel("Debug")
			negroni := negroni.New()
			negroni.Use(handlers.NewRequestInfo())
			negroni.Use(handlers.NewReporter(metricReporter, logger))

			testApp = httptest.NewServer(negroni)
		})
		AfterEach(func() {
			// Set GODEBUG back to whatever it was prior to this test
			os.Setenv("GODEBUG", godebug)
		})

		It("counts request with empty content-length header correctly", func() {
			u, err := url.Parse(testApp.URL)
			Expect(err).NotTo(HaveOccurred())
			conn, err := net.Dial("tcp", u.Host)
			Expect(err).NotTo(HaveOccurred())
			defer conn.Close()
			fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: sample.com\r\nContent-Length: \r\n\r\n")
			_, err = bufio.NewReader(conn).ReadString('\n')
			Expect(err).NotTo(HaveOccurred())

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("empty_content_length_header"))
		})

		It("does not count request without content-length header", func() {
			u, err := url.Parse(testApp.URL)
			Expect(err).NotTo(HaveOccurred())
			conn, err := net.Dial("tcp", u.Host)
			Expect(err).NotTo(HaveOccurred())
			defer conn.Close()
			fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: sample.com\r\n\r\n")
			_, err = bufio.NewReader(conn).ReadString('\n')
			Expect(err).NotTo(HaveOccurred())

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(0))
		})

		It("does not count request with correct content-length header", func() {
			u, err := url.Parse(testApp.URL)
			Expect(err).NotTo(HaveOccurred())
			conn, err := net.Dial("tcp", u.Host)
			Expect(err).NotTo(HaveOccurred())
			defer conn.Close()
			fmt.Fprintf(conn, "POST / HTTP/1.1\r\nHost: sample.com\r\nContent-Length: 5\r\n\r\n12345")
			_, err = bufio.NewReader(conn).ReadString('\n')
			Expect(err).NotTo(HaveOccurred())

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(0))
		})
	})

	It("sends the latency", func() {
		metricReporter.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 2*time.Second)

		Expect(sender.SendValueCallCount()).To(Equal(1))
		name, value, unit := sender.SendValueArgsForCall(0)
		Expect(name).To(Equal("latency"))
		Expect(value).To(BeEquivalentTo(2000))
		Expect(unit).To(Equal("ms"))

	})

	It("does not send the latency if switched off", func() {
		metricReporter.PerRequestMetricsReporting = false
		metricReporter.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 2*time.Second)

		Expect(sender.SendValueCallCount()).To(Equal(0))

	})

	It("sends the latency for the given component", func() {
		endpoint.Tags["component"] = "CloudController"
		metricReporter.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 2*time.Second)

		Expect(sender.SendValueCallCount()).To(Equal(2))
		name, value, unit := sender.SendValueArgsForCall(1)
		Expect(name).To(Equal("latency.CloudController"))
		Expect(value).To(BeEquivalentTo(2000))
		Expect(unit).To(Equal("ms"))
	})

	It("does not send the latency for the given component if switched off", func() {
		metricReporter.PerRequestMetricsReporting = false
		endpoint.Tags["component"] = "CloudController"
		metricReporter.CaptureRoutingResponseLatency(endpoint, 0, time.Time{}, 2*time.Second)

		Expect(sender.SendValueCallCount()).To(Equal(0))
	})

	It("sends the gorouter time", func() {
		metricReporter.CaptureGorouterTime(3)

		Expect(sender.SendValueCallCount()).To(Equal(1))
		name, value, unit := sender.SendValueArgsForCall(0)
		Expect(name).To(Equal("gorouter_time"))
		Expect(value).To(BeEquivalentTo(3000))
		Expect(unit).To(Equal("ms"))

	})

	It("does not send the goroutertime if switched off", func() {
		metricReporter.PerRequestMetricsReporting = false
		metricReporter.CaptureGorouterTime(3)

		Expect(sender.SendValueCallCount()).To(Equal(0))

	})

	Context("sends route metrics", func() {
		var endpoint *route.Endpoint

		BeforeEach(func() {
			endpoint = new(route.Endpoint)
		})

		It("sends number of nats messages received from each component", func() {
			endpoint.Tags = map[string]string{}
			metricReporter.CaptureRegistryMessage(endpoint, route.ADDED.String())

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("registry_message"))
		})

		It("sends number of nats messages received from each component", func() {
			endpoint.Tags = map[string]string{"component": "uaa"}
			metricReporter.CaptureRegistryMessage(endpoint, route.ADDED.String())

			endpoint.Tags = map[string]string{"component": "route-emitter"}
			metricReporter.CaptureRegistryMessage(endpoint, route.ADDED.String())

			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(2))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("registry_message.uaa"))
			Expect(batcher.BatchIncrementCounterArgsForCall(1)).To(Equal("registry_message.route-emitter"))
		})

		It("sends the total routes", func() {
			metricReporter.CaptureRouteStats(12, 5)

			Expect(sender.SendValueCallCount()).To(Equal(2))
			name, value, unit := sender.SendValueArgsForCall(0)
			Expect(name).To(Equal("total_routes"))
			Expect(value).To(BeEquivalentTo(12))
			Expect(unit).To(Equal(""))
		})

		It("sends the time since last update", func() {
			metricReporter.CaptureRouteStats(12, 5)

			Expect(sender.SendValueCallCount()).To(Equal(2))
			name, value, unit := sender.SendValueArgsForCall(1)
			Expect(name).To(Equal("ms_since_last_registry_update"))
			Expect(value).To(BeEquivalentTo(5))
			Expect(unit).To(Equal("ms"))
		})

		It("sends the lookup time for routing table", func() {
			metricReporter.CaptureLookupTime(time.Duration(9) * time.Second)

			Expect(sender.SendValueCallCount()).To(Equal(1))
			name, value, unit := sender.SendValueArgsForCall(0)
			Expect(name).To(Equal("route_lookup_time"))
			Expect(value).To(BeEquivalentTo(9000000000))
			Expect(unit).To(Equal("ns"))
		})
	})

	It("increments the routes_pruned metric", func() {
		metricReporter.CaptureRoutesPruned(5)
		Expect(batcher.BatchAddCounterCallCount()).To(Equal(1))
		metric, count := batcher.BatchAddCounterArgsForCall(0)
		Expect(metric).To(Equal("routes_pruned"))
		Expect(count).To(Equal(uint64(5)))
	})

	It("increments the backend_tls_handshake_failed metric", func() {
		metricReporter.CaptureBackendTLSHandshakeFailed()
		Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
		Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("backend_tls_handshake_failed"))
	})

	Describe("Unregister messages", func() {
		var endpoint *route.Endpoint
		Context("when unregister msg with component name is incremented", func() {
			BeforeEach(func() {
				endpoint = new(route.Endpoint)
				endpoint.Tags = map[string]string{"component": "oauth-server"}
				metricReporter.CaptureUnregistryMessage(endpoint)
			})

			It("increments the counter metric", func() {
				Expect(sender.IncrementCounterCallCount()).To(Equal(1))
				Expect(sender.IncrementCounterArgsForCall(0)).To(Equal("unregistry_message.oauth-server"))
			})

			It("increments the counter metric for each component unregistered", func() {
				endpointTwo := new(route.Endpoint)
				endpointTwo.Tags = map[string]string{"component": "api-server"}
				metricReporter.CaptureUnregistryMessage(endpointTwo)

				Expect(sender.IncrementCounterCallCount()).To(Equal(2))
				Expect(sender.IncrementCounterArgsForCall(0)).To(Equal("unregistry_message.oauth-server"))
				Expect(sender.IncrementCounterArgsForCall(1)).To(Equal("unregistry_message.api-server"))
			})
		})
		Context("when unregister msg with empty component name is incremented", func() {
			BeforeEach(func() {
				endpoint = new(route.Endpoint)
				endpoint.Tags = map[string]string{}
				metricReporter.CaptureUnregistryMessage(endpoint)
			})
			It("increments the counter metric", func() {
				Expect(sender.IncrementCounterCallCount()).To(Equal(1))
				Expect(sender.IncrementCounterArgsForCall(0)).To(Equal("unregistry_message"))
			})
		})
	})

	Context("websocket metrics", func() {
		It("increments the total responses metric", func() {
			metricReporter.CaptureWebSocketUpdate()
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("websocket_upgrades"))
		})
		It("increments the websocket failures metric", func() {
			metricReporter.CaptureWebSocketFailure()
			Expect(batcher.BatchIncrementCounterCallCount()).To(Equal(1))
			Expect(batcher.BatchIncrementCounterArgsForCall(0)).To(Equal("websocket_failures"))
		})
	})

	Describe("Monitor metrics", func() {
		Context("file descriptor metrics sent", func() {
			It("increments the fd gauge metric", func() {
				metricReporter.CaptureFoundFileDescriptors(10)
				Expect(sender.SendValueCallCount()).To(Equal(1))
				name, value, unit := sender.SendValueArgsForCall(0)
				Expect(name).To(Equal("file_descriptors"))
				Expect(value).To(BeEquivalentTo(10))
				Expect(unit).To(Equal("file"))
			})
		})
		Context("NATS message metrics sent", func() {
			It("increments the buffered messages metric", func() {
				metricReporter.CaptureNATSBufferedMessages(100)
				Expect(sender.SendValueCallCount()).To(Equal(1))
				name, value, unit := sender.SendValueArgsForCall(0)
				Expect(name).To(Equal("buffered_messages"))
				Expect(value).To(BeEquivalentTo(100))
				Expect(unit).To(Equal("message"))
			})
			It("increments the dropped messages metric", func() {
				metricReporter.CaptureNATSDroppedMessages(200)
				Expect(sender.SendValueCallCount()).To(Equal(1))
				name, value, unit := sender.SendValueArgsForCall(0)
				Expect(name).To(Equal("total_dropped_messages"))
				Expect(value).To(BeEquivalentTo(200))
				Expect(unit).To(Equal("message"))
			})
		})
	})

	Describe("CaptureRouteRegistrationLatency", func() {
		It("is muzzled by default", func() {
			metricReporter.CaptureRouteRegistrationLatency(2 * time.Second)
			Expect(sender.SendValueCallCount()).To(Equal(0))
		})
		It("sends router registration latency when unmuzzled", func() {
			metricReporter.UnmuzzleRouteRegistrationLatency()
			metricReporter.CaptureRouteRegistrationLatency(2 * time.Second)
			Expect(sender.SendValueCallCount()).To(Equal(1))
			name, value, unit := sender.SendValueArgsForCall(0)
			Expect(name).To(Equal("route_registration_latency"))
			Expect(value).To(BeEquivalentTo(2000))
			Expect(unit).To(Equal("ms"))
		})
	})

})
