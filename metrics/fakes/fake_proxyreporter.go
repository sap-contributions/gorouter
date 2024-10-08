// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"
	"time"

	"code.cloudfoundry.org/gorouter/metrics"
	"code.cloudfoundry.org/gorouter/route"
)

type FakeProxyReporter struct {
	CaptureBackendExhaustedConnsStub        func()
	captureBackendExhaustedConnsMutex       sync.RWMutex
	captureBackendExhaustedConnsArgsForCall []struct {
	}
	CaptureBackendInvalidIDStub        func()
	captureBackendInvalidIDMutex       sync.RWMutex
	captureBackendInvalidIDArgsForCall []struct {
	}
	CaptureBackendInvalidTLSCertStub        func()
	captureBackendInvalidTLSCertMutex       sync.RWMutex
	captureBackendInvalidTLSCertArgsForCall []struct {
	}
	CaptureBackendTLSHandshakeFailedStub        func()
	captureBackendTLSHandshakeFailedMutex       sync.RWMutex
	captureBackendTLSHandshakeFailedArgsForCall []struct {
	}
	CaptureBadGatewayStub        func()
	captureBadGatewayMutex       sync.RWMutex
	captureBadGatewayArgsForCall []struct {
	}
	CaptureBadRequestStub        func()
	captureBadRequestMutex       sync.RWMutex
	captureBadRequestArgsForCall []struct {
	}
	CaptureEmptyContentLengthHeaderStub        func()
	captureEmptyContentLengthHeaderMutex       sync.RWMutex
	captureEmptyContentLengthHeaderArgsForCall []struct {
	}
	CaptureRouteServiceResponseStub        func(*http.Response)
	captureRouteServiceResponseMutex       sync.RWMutex
	captureRouteServiceResponseArgsForCall []struct {
		arg1 *http.Response
	}
	CaptureRoutingRequestStub        func(*route.Endpoint)
	captureRoutingRequestMutex       sync.RWMutex
	captureRoutingRequestArgsForCall []struct {
		arg1 *route.Endpoint
	}
	CaptureRoutingResponseStub        func(int)
	captureRoutingResponseMutex       sync.RWMutex
	captureRoutingResponseArgsForCall []struct {
		arg1 int
	}
	CaptureRoutingResponseLatencyStub        func(*route.Endpoint, int, time.Time, time.Duration)
	captureRoutingResponseLatencyMutex       sync.RWMutex
	captureRoutingResponseLatencyArgsForCall []struct {
		arg1 *route.Endpoint
		arg2 int
		arg3 time.Time
		arg4 time.Duration
	}
	CaptureWebSocketFailureStub        func()
	captureWebSocketFailureMutex       sync.RWMutex
	captureWebSocketFailureArgsForCall []struct {
	}
	CaptureWebSocketUpdateStub        func()
	captureWebSocketUpdateMutex       sync.RWMutex
	captureWebSocketUpdateArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProxyReporter) CaptureBackendExhaustedConns() {
	fake.captureBackendExhaustedConnsMutex.Lock()
	fake.captureBackendExhaustedConnsArgsForCall = append(fake.captureBackendExhaustedConnsArgsForCall, struct {
	}{})
	stub := fake.CaptureBackendExhaustedConnsStub
	fake.recordInvocation("CaptureBackendExhaustedConns", []interface{}{})
	fake.captureBackendExhaustedConnsMutex.Unlock()
	if stub != nil {
		fake.CaptureBackendExhaustedConnsStub()
	}
}

func (fake *FakeProxyReporter) CaptureBackendExhaustedConnsCallCount() int {
	fake.captureBackendExhaustedConnsMutex.RLock()
	defer fake.captureBackendExhaustedConnsMutex.RUnlock()
	return len(fake.captureBackendExhaustedConnsArgsForCall)
}

func (fake *FakeProxyReporter) CaptureBackendExhaustedConnsCalls(stub func()) {
	fake.captureBackendExhaustedConnsMutex.Lock()
	defer fake.captureBackendExhaustedConnsMutex.Unlock()
	fake.CaptureBackendExhaustedConnsStub = stub
}

func (fake *FakeProxyReporter) CaptureBackendInvalidID() {
	fake.captureBackendInvalidIDMutex.Lock()
	fake.captureBackendInvalidIDArgsForCall = append(fake.captureBackendInvalidIDArgsForCall, struct {
	}{})
	stub := fake.CaptureBackendInvalidIDStub
	fake.recordInvocation("CaptureBackendInvalidID", []interface{}{})
	fake.captureBackendInvalidIDMutex.Unlock()
	if stub != nil {
		fake.CaptureBackendInvalidIDStub()
	}
}

func (fake *FakeProxyReporter) CaptureBackendInvalidIDCallCount() int {
	fake.captureBackendInvalidIDMutex.RLock()
	defer fake.captureBackendInvalidIDMutex.RUnlock()
	return len(fake.captureBackendInvalidIDArgsForCall)
}

func (fake *FakeProxyReporter) CaptureBackendInvalidIDCalls(stub func()) {
	fake.captureBackendInvalidIDMutex.Lock()
	defer fake.captureBackendInvalidIDMutex.Unlock()
	fake.CaptureBackendInvalidIDStub = stub
}

func (fake *FakeProxyReporter) CaptureBackendInvalidTLSCert() {
	fake.captureBackendInvalidTLSCertMutex.Lock()
	fake.captureBackendInvalidTLSCertArgsForCall = append(fake.captureBackendInvalidTLSCertArgsForCall, struct {
	}{})
	stub := fake.CaptureBackendInvalidTLSCertStub
	fake.recordInvocation("CaptureBackendInvalidTLSCert", []interface{}{})
	fake.captureBackendInvalidTLSCertMutex.Unlock()
	if stub != nil {
		fake.CaptureBackendInvalidTLSCertStub()
	}
}

func (fake *FakeProxyReporter) CaptureBackendInvalidTLSCertCallCount() int {
	fake.captureBackendInvalidTLSCertMutex.RLock()
	defer fake.captureBackendInvalidTLSCertMutex.RUnlock()
	return len(fake.captureBackendInvalidTLSCertArgsForCall)
}

func (fake *FakeProxyReporter) CaptureBackendInvalidTLSCertCalls(stub func()) {
	fake.captureBackendInvalidTLSCertMutex.Lock()
	defer fake.captureBackendInvalidTLSCertMutex.Unlock()
	fake.CaptureBackendInvalidTLSCertStub = stub
}

func (fake *FakeProxyReporter) CaptureBackendTLSHandshakeFailed() {
	fake.captureBackendTLSHandshakeFailedMutex.Lock()
	fake.captureBackendTLSHandshakeFailedArgsForCall = append(fake.captureBackendTLSHandshakeFailedArgsForCall, struct {
	}{})
	stub := fake.CaptureBackendTLSHandshakeFailedStub
	fake.recordInvocation("CaptureBackendTLSHandshakeFailed", []interface{}{})
	fake.captureBackendTLSHandshakeFailedMutex.Unlock()
	if stub != nil {
		fake.CaptureBackendTLSHandshakeFailedStub()
	}
}

func (fake *FakeProxyReporter) CaptureBackendTLSHandshakeFailedCallCount() int {
	fake.captureBackendTLSHandshakeFailedMutex.RLock()
	defer fake.captureBackendTLSHandshakeFailedMutex.RUnlock()
	return len(fake.captureBackendTLSHandshakeFailedArgsForCall)
}

func (fake *FakeProxyReporter) CaptureBackendTLSHandshakeFailedCalls(stub func()) {
	fake.captureBackendTLSHandshakeFailedMutex.Lock()
	defer fake.captureBackendTLSHandshakeFailedMutex.Unlock()
	fake.CaptureBackendTLSHandshakeFailedStub = stub
}

func (fake *FakeProxyReporter) CaptureBadGateway() {
	fake.captureBadGatewayMutex.Lock()
	fake.captureBadGatewayArgsForCall = append(fake.captureBadGatewayArgsForCall, struct {
	}{})
	stub := fake.CaptureBadGatewayStub
	fake.recordInvocation("CaptureBadGateway", []interface{}{})
	fake.captureBadGatewayMutex.Unlock()
	if stub != nil {
		fake.CaptureBadGatewayStub()
	}
}

func (fake *FakeProxyReporter) CaptureBadGatewayCallCount() int {
	fake.captureBadGatewayMutex.RLock()
	defer fake.captureBadGatewayMutex.RUnlock()
	return len(fake.captureBadGatewayArgsForCall)
}

func (fake *FakeProxyReporter) CaptureBadGatewayCalls(stub func()) {
	fake.captureBadGatewayMutex.Lock()
	defer fake.captureBadGatewayMutex.Unlock()
	fake.CaptureBadGatewayStub = stub
}

func (fake *FakeProxyReporter) CaptureBadRequest() {
	fake.captureBadRequestMutex.Lock()
	fake.captureBadRequestArgsForCall = append(fake.captureBadRequestArgsForCall, struct {
	}{})
	stub := fake.CaptureBadRequestStub
	fake.recordInvocation("CaptureBadRequest", []interface{}{})
	fake.captureBadRequestMutex.Unlock()
	if stub != nil {
		fake.CaptureBadRequestStub()
	}
}

func (fake *FakeProxyReporter) CaptureBadRequestCallCount() int {
	fake.captureBadRequestMutex.RLock()
	defer fake.captureBadRequestMutex.RUnlock()
	return len(fake.captureBadRequestArgsForCall)
}

func (fake *FakeProxyReporter) CaptureBadRequestCalls(stub func()) {
	fake.captureBadRequestMutex.Lock()
	defer fake.captureBadRequestMutex.Unlock()
	fake.CaptureBadRequestStub = stub
}

func (fake *FakeProxyReporter) CaptureEmptyContentLengthHeader() {
	fake.captureEmptyContentLengthHeaderMutex.Lock()
	fake.captureEmptyContentLengthHeaderArgsForCall = append(fake.captureEmptyContentLengthHeaderArgsForCall, struct {
	}{})
	stub := fake.CaptureEmptyContentLengthHeaderStub
	fake.recordInvocation("CaptureEmptyContentLengthHeader", []interface{}{})
	fake.captureEmptyContentLengthHeaderMutex.Unlock()
	if stub != nil {
		fake.CaptureEmptyContentLengthHeaderStub()
	}
}

func (fake *FakeProxyReporter) CaptureEmptyContentLengthHeaderCallCount() int {
	fake.captureEmptyContentLengthHeaderMutex.RLock()
	defer fake.captureEmptyContentLengthHeaderMutex.RUnlock()
	return len(fake.captureEmptyContentLengthHeaderArgsForCall)
}

func (fake *FakeProxyReporter) CaptureEmptyContentLengthHeaderCalls(stub func()) {
	fake.captureEmptyContentLengthHeaderMutex.Lock()
	defer fake.captureEmptyContentLengthHeaderMutex.Unlock()
	fake.CaptureEmptyContentLengthHeaderStub = stub
}

func (fake *FakeProxyReporter) CaptureRouteServiceResponse(arg1 *http.Response) {
	fake.captureRouteServiceResponseMutex.Lock()
	fake.captureRouteServiceResponseArgsForCall = append(fake.captureRouteServiceResponseArgsForCall, struct {
		arg1 *http.Response
	}{arg1})
	stub := fake.CaptureRouteServiceResponseStub
	fake.recordInvocation("CaptureRouteServiceResponse", []interface{}{arg1})
	fake.captureRouteServiceResponseMutex.Unlock()
	if stub != nil {
		fake.CaptureRouteServiceResponseStub(arg1)
	}
}

func (fake *FakeProxyReporter) CaptureRouteServiceResponseCallCount() int {
	fake.captureRouteServiceResponseMutex.RLock()
	defer fake.captureRouteServiceResponseMutex.RUnlock()
	return len(fake.captureRouteServiceResponseArgsForCall)
}

func (fake *FakeProxyReporter) CaptureRouteServiceResponseCalls(stub func(*http.Response)) {
	fake.captureRouteServiceResponseMutex.Lock()
	defer fake.captureRouteServiceResponseMutex.Unlock()
	fake.CaptureRouteServiceResponseStub = stub
}

func (fake *FakeProxyReporter) CaptureRouteServiceResponseArgsForCall(i int) *http.Response {
	fake.captureRouteServiceResponseMutex.RLock()
	defer fake.captureRouteServiceResponseMutex.RUnlock()
	argsForCall := fake.captureRouteServiceResponseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProxyReporter) CaptureRoutingRequest(arg1 *route.Endpoint) {
	fake.captureRoutingRequestMutex.Lock()
	fake.captureRoutingRequestArgsForCall = append(fake.captureRoutingRequestArgsForCall, struct {
		arg1 *route.Endpoint
	}{arg1})
	stub := fake.CaptureRoutingRequestStub
	fake.recordInvocation("CaptureRoutingRequest", []interface{}{arg1})
	fake.captureRoutingRequestMutex.Unlock()
	if stub != nil {
		fake.CaptureRoutingRequestStub(arg1)
	}
}

func (fake *FakeProxyReporter) CaptureRoutingRequestCallCount() int {
	fake.captureRoutingRequestMutex.RLock()
	defer fake.captureRoutingRequestMutex.RUnlock()
	return len(fake.captureRoutingRequestArgsForCall)
}

func (fake *FakeProxyReporter) CaptureRoutingRequestCalls(stub func(*route.Endpoint)) {
	fake.captureRoutingRequestMutex.Lock()
	defer fake.captureRoutingRequestMutex.Unlock()
	fake.CaptureRoutingRequestStub = stub
}

func (fake *FakeProxyReporter) CaptureRoutingRequestArgsForCall(i int) *route.Endpoint {
	fake.captureRoutingRequestMutex.RLock()
	defer fake.captureRoutingRequestMutex.RUnlock()
	argsForCall := fake.captureRoutingRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProxyReporter) CaptureRoutingResponse(arg1 int) {
	fake.captureRoutingResponseMutex.Lock()
	fake.captureRoutingResponseArgsForCall = append(fake.captureRoutingResponseArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.CaptureRoutingResponseStub
	fake.recordInvocation("CaptureRoutingResponse", []interface{}{arg1})
	fake.captureRoutingResponseMutex.Unlock()
	if stub != nil {
		fake.CaptureRoutingResponseStub(arg1)
	}
}

func (fake *FakeProxyReporter) CaptureRoutingResponseCallCount() int {
	fake.captureRoutingResponseMutex.RLock()
	defer fake.captureRoutingResponseMutex.RUnlock()
	return len(fake.captureRoutingResponseArgsForCall)
}

func (fake *FakeProxyReporter) CaptureRoutingResponseCalls(stub func(int)) {
	fake.captureRoutingResponseMutex.Lock()
	defer fake.captureRoutingResponseMutex.Unlock()
	fake.CaptureRoutingResponseStub = stub
}

func (fake *FakeProxyReporter) CaptureRoutingResponseArgsForCall(i int) int {
	fake.captureRoutingResponseMutex.RLock()
	defer fake.captureRoutingResponseMutex.RUnlock()
	argsForCall := fake.captureRoutingResponseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProxyReporter) CaptureRoutingResponseLatency(arg1 *route.Endpoint, arg2 int, arg3 time.Time, arg4 time.Duration) {
	fake.captureRoutingResponseLatencyMutex.Lock()
	fake.captureRoutingResponseLatencyArgsForCall = append(fake.captureRoutingResponseLatencyArgsForCall, struct {
		arg1 *route.Endpoint
		arg2 int
		arg3 time.Time
		arg4 time.Duration
	}{arg1, arg2, arg3, arg4})
	stub := fake.CaptureRoutingResponseLatencyStub
	fake.recordInvocation("CaptureRoutingResponseLatency", []interface{}{arg1, arg2, arg3, arg4})
	fake.captureRoutingResponseLatencyMutex.Unlock()
	if stub != nil {
		fake.CaptureRoutingResponseLatencyStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *FakeProxyReporter) CaptureRoutingResponseLatencyCallCount() int {
	fake.captureRoutingResponseLatencyMutex.RLock()
	defer fake.captureRoutingResponseLatencyMutex.RUnlock()
	return len(fake.captureRoutingResponseLatencyArgsForCall)
}

func (fake *FakeProxyReporter) CaptureRoutingResponseLatencyCalls(stub func(*route.Endpoint, int, time.Time, time.Duration)) {
	fake.captureRoutingResponseLatencyMutex.Lock()
	defer fake.captureRoutingResponseLatencyMutex.Unlock()
	fake.CaptureRoutingResponseLatencyStub = stub
}

func (fake *FakeProxyReporter) CaptureRoutingResponseLatencyArgsForCall(i int) (*route.Endpoint, int, time.Time, time.Duration) {
	fake.captureRoutingResponseLatencyMutex.RLock()
	defer fake.captureRoutingResponseLatencyMutex.RUnlock()
	argsForCall := fake.captureRoutingResponseLatencyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeProxyReporter) CaptureWebSocketFailure() {
	fake.captureWebSocketFailureMutex.Lock()
	fake.captureWebSocketFailureArgsForCall = append(fake.captureWebSocketFailureArgsForCall, struct {
	}{})
	stub := fake.CaptureWebSocketFailureStub
	fake.recordInvocation("CaptureWebSocketFailure", []interface{}{})
	fake.captureWebSocketFailureMutex.Unlock()
	if stub != nil {
		fake.CaptureWebSocketFailureStub()
	}
}

func (fake *FakeProxyReporter) CaptureWebSocketFailureCallCount() int {
	fake.captureWebSocketFailureMutex.RLock()
	defer fake.captureWebSocketFailureMutex.RUnlock()
	return len(fake.captureWebSocketFailureArgsForCall)
}

func (fake *FakeProxyReporter) CaptureWebSocketFailureCalls(stub func()) {
	fake.captureWebSocketFailureMutex.Lock()
	defer fake.captureWebSocketFailureMutex.Unlock()
	fake.CaptureWebSocketFailureStub = stub
}

func (fake *FakeProxyReporter) CaptureWebSocketUpdate() {
	fake.captureWebSocketUpdateMutex.Lock()
	fake.captureWebSocketUpdateArgsForCall = append(fake.captureWebSocketUpdateArgsForCall, struct {
	}{})
	stub := fake.CaptureWebSocketUpdateStub
	fake.recordInvocation("CaptureWebSocketUpdate", []interface{}{})
	fake.captureWebSocketUpdateMutex.Unlock()
	if stub != nil {
		fake.CaptureWebSocketUpdateStub()
	}
}

func (fake *FakeProxyReporter) CaptureWebSocketUpdateCallCount() int {
	fake.captureWebSocketUpdateMutex.RLock()
	defer fake.captureWebSocketUpdateMutex.RUnlock()
	return len(fake.captureWebSocketUpdateArgsForCall)
}

func (fake *FakeProxyReporter) CaptureWebSocketUpdateCalls(stub func()) {
	fake.captureWebSocketUpdateMutex.Lock()
	defer fake.captureWebSocketUpdateMutex.Unlock()
	fake.CaptureWebSocketUpdateStub = stub
}

func (fake *FakeProxyReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.captureBackendExhaustedConnsMutex.RLock()
	defer fake.captureBackendExhaustedConnsMutex.RUnlock()
	fake.captureBackendInvalidIDMutex.RLock()
	defer fake.captureBackendInvalidIDMutex.RUnlock()
	fake.captureBackendInvalidTLSCertMutex.RLock()
	defer fake.captureBackendInvalidTLSCertMutex.RUnlock()
	fake.captureBackendTLSHandshakeFailedMutex.RLock()
	defer fake.captureBackendTLSHandshakeFailedMutex.RUnlock()
	fake.captureBadGatewayMutex.RLock()
	defer fake.captureBadGatewayMutex.RUnlock()
	fake.captureBadRequestMutex.RLock()
	defer fake.captureBadRequestMutex.RUnlock()
	fake.captureEmptyContentLengthHeaderMutex.RLock()
	defer fake.captureEmptyContentLengthHeaderMutex.RUnlock()
	fake.captureRouteServiceResponseMutex.RLock()
	defer fake.captureRouteServiceResponseMutex.RUnlock()
	fake.captureRoutingRequestMutex.RLock()
	defer fake.captureRoutingRequestMutex.RUnlock()
	fake.captureRoutingResponseMutex.RLock()
	defer fake.captureRoutingResponseMutex.RUnlock()
	fake.captureRoutingResponseLatencyMutex.RLock()
	defer fake.captureRoutingResponseLatencyMutex.RUnlock()
	fake.captureWebSocketFailureMutex.RLock()
	defer fake.captureWebSocketFailureMutex.RUnlock()
	fake.captureWebSocketUpdateMutex.RLock()
	defer fake.captureWebSocketUpdateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProxyReporter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.ProxyReporter = new(FakeProxyReporter)
