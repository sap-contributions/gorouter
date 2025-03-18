// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"
	"time"

	"code.cloudfoundry.org/gorouter/metrics"
	"code.cloudfoundry.org/gorouter/route"
)

type FakeMetricReporter struct {
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
	CaptureFoundFileDescriptorsStub        func(int)
	captureFoundFileDescriptorsMutex       sync.RWMutex
	captureFoundFileDescriptorsArgsForCall []struct {
		arg1 int
	}
	CaptureHTTPLatencyStub        func(time.Duration, string)
	captureHTTPLatencyMutex       sync.RWMutex
	captureHTTPLatencyArgsForCall []struct {
		arg1 time.Duration
		arg2 string
	}
	CaptureLookupTimeStub        func(time.Duration)
	captureLookupTimeMutex       sync.RWMutex
	captureLookupTimeArgsForCall []struct {
		arg1 time.Duration
	}
	CaptureNATSBufferedMessagesStub        func(int)
	captureNATSBufferedMessagesMutex       sync.RWMutex
	captureNATSBufferedMessagesArgsForCall []struct {
		arg1 int
	}
	CaptureNATSDroppedMessagesStub        func(int)
	captureNATSDroppedMessagesMutex       sync.RWMutex
	captureNATSDroppedMessagesArgsForCall []struct {
		arg1 int
	}
	CaptureRegistryMessageStub        func(metrics.ComponentTagged, string)
	captureRegistryMessageMutex       sync.RWMutex
	captureRegistryMessageArgsForCall []struct {
		arg1 metrics.ComponentTagged
		arg2 string
	}
	CaptureRouteRegistrationLatencyStub        func(time.Duration)
	captureRouteRegistrationLatencyMutex       sync.RWMutex
	captureRouteRegistrationLatencyArgsForCall []struct {
		arg1 time.Duration
	}
	CaptureRouteServiceResponseStub        func(*http.Response)
	captureRouteServiceResponseMutex       sync.RWMutex
	captureRouteServiceResponseArgsForCall []struct {
		arg1 *http.Response
	}
	CaptureRouteStatsStub        func(int, int64)
	captureRouteStatsMutex       sync.RWMutex
	captureRouteStatsArgsForCall []struct {
		arg1 int
		arg2 int64
	}
	CaptureRoutesPrunedStub        func(uint64)
	captureRoutesPrunedMutex       sync.RWMutex
	captureRoutesPrunedArgsForCall []struct {
		arg1 uint64
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
	CaptureUnregistryMessageStub        func(metrics.ComponentTagged)
	captureUnregistryMessageMutex       sync.RWMutex
	captureUnregistryMessageArgsForCall []struct {
		arg1 metrics.ComponentTagged
	}
	CaptureWebSocketFailureStub        func()
	captureWebSocketFailureMutex       sync.RWMutex
	captureWebSocketFailureArgsForCall []struct {
	}
	CaptureWebSocketUpdateStub        func()
	captureWebSocketUpdateMutex       sync.RWMutex
	captureWebSocketUpdateArgsForCall []struct {
	}
	UnmuzzleRouteRegistrationLatencyStub        func()
	unmuzzleRouteRegistrationLatencyMutex       sync.RWMutex
	unmuzzleRouteRegistrationLatencyArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetricReporter) CaptureBackendExhaustedConns() {
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

func (fake *FakeMetricReporter) CaptureBackendExhaustedConnsCallCount() int {
	fake.captureBackendExhaustedConnsMutex.RLock()
	defer fake.captureBackendExhaustedConnsMutex.RUnlock()
	return len(fake.captureBackendExhaustedConnsArgsForCall)
}

func (fake *FakeMetricReporter) CaptureBackendExhaustedConnsCalls(stub func()) {
	fake.captureBackendExhaustedConnsMutex.Lock()
	defer fake.captureBackendExhaustedConnsMutex.Unlock()
	fake.CaptureBackendExhaustedConnsStub = stub
}

func (fake *FakeMetricReporter) CaptureBackendInvalidID() {
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

func (fake *FakeMetricReporter) CaptureBackendInvalidIDCallCount() int {
	fake.captureBackendInvalidIDMutex.RLock()
	defer fake.captureBackendInvalidIDMutex.RUnlock()
	return len(fake.captureBackendInvalidIDArgsForCall)
}

func (fake *FakeMetricReporter) CaptureBackendInvalidIDCalls(stub func()) {
	fake.captureBackendInvalidIDMutex.Lock()
	defer fake.captureBackendInvalidIDMutex.Unlock()
	fake.CaptureBackendInvalidIDStub = stub
}

func (fake *FakeMetricReporter) CaptureBackendInvalidTLSCert() {
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

func (fake *FakeMetricReporter) CaptureBackendInvalidTLSCertCallCount() int {
	fake.captureBackendInvalidTLSCertMutex.RLock()
	defer fake.captureBackendInvalidTLSCertMutex.RUnlock()
	return len(fake.captureBackendInvalidTLSCertArgsForCall)
}

func (fake *FakeMetricReporter) CaptureBackendInvalidTLSCertCalls(stub func()) {
	fake.captureBackendInvalidTLSCertMutex.Lock()
	defer fake.captureBackendInvalidTLSCertMutex.Unlock()
	fake.CaptureBackendInvalidTLSCertStub = stub
}

func (fake *FakeMetricReporter) CaptureBackendTLSHandshakeFailed() {
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

func (fake *FakeMetricReporter) CaptureBackendTLSHandshakeFailedCallCount() int {
	fake.captureBackendTLSHandshakeFailedMutex.RLock()
	defer fake.captureBackendTLSHandshakeFailedMutex.RUnlock()
	return len(fake.captureBackendTLSHandshakeFailedArgsForCall)
}

func (fake *FakeMetricReporter) CaptureBackendTLSHandshakeFailedCalls(stub func()) {
	fake.captureBackendTLSHandshakeFailedMutex.Lock()
	defer fake.captureBackendTLSHandshakeFailedMutex.Unlock()
	fake.CaptureBackendTLSHandshakeFailedStub = stub
}

func (fake *FakeMetricReporter) CaptureBadGateway() {
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

func (fake *FakeMetricReporter) CaptureBadGatewayCallCount() int {
	fake.captureBadGatewayMutex.RLock()
	defer fake.captureBadGatewayMutex.RUnlock()
	return len(fake.captureBadGatewayArgsForCall)
}

func (fake *FakeMetricReporter) CaptureBadGatewayCalls(stub func()) {
	fake.captureBadGatewayMutex.Lock()
	defer fake.captureBadGatewayMutex.Unlock()
	fake.CaptureBadGatewayStub = stub
}

func (fake *FakeMetricReporter) CaptureBadRequest() {
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

func (fake *FakeMetricReporter) CaptureBadRequestCallCount() int {
	fake.captureBadRequestMutex.RLock()
	defer fake.captureBadRequestMutex.RUnlock()
	return len(fake.captureBadRequestArgsForCall)
}

func (fake *FakeMetricReporter) CaptureBadRequestCalls(stub func()) {
	fake.captureBadRequestMutex.Lock()
	defer fake.captureBadRequestMutex.Unlock()
	fake.CaptureBadRequestStub = stub
}

func (fake *FakeMetricReporter) CaptureEmptyContentLengthHeader() {
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

func (fake *FakeMetricReporter) CaptureEmptyContentLengthHeaderCallCount() int {
	fake.captureEmptyContentLengthHeaderMutex.RLock()
	defer fake.captureEmptyContentLengthHeaderMutex.RUnlock()
	return len(fake.captureEmptyContentLengthHeaderArgsForCall)
}

func (fake *FakeMetricReporter) CaptureEmptyContentLengthHeaderCalls(stub func()) {
	fake.captureEmptyContentLengthHeaderMutex.Lock()
	defer fake.captureEmptyContentLengthHeaderMutex.Unlock()
	fake.CaptureEmptyContentLengthHeaderStub = stub
}

func (fake *FakeMetricReporter) CaptureFoundFileDescriptors(arg1 int) {
	fake.captureFoundFileDescriptorsMutex.Lock()
	fake.captureFoundFileDescriptorsArgsForCall = append(fake.captureFoundFileDescriptorsArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.CaptureFoundFileDescriptorsStub
	fake.recordInvocation("CaptureFoundFileDescriptors", []interface{}{arg1})
	fake.captureFoundFileDescriptorsMutex.Unlock()
	if stub != nil {
		fake.CaptureFoundFileDescriptorsStub(arg1)
	}
}

func (fake *FakeMetricReporter) CaptureFoundFileDescriptorsCallCount() int {
	fake.captureFoundFileDescriptorsMutex.RLock()
	defer fake.captureFoundFileDescriptorsMutex.RUnlock()
	return len(fake.captureFoundFileDescriptorsArgsForCall)
}

func (fake *FakeMetricReporter) CaptureFoundFileDescriptorsCalls(stub func(int)) {
	fake.captureFoundFileDescriptorsMutex.Lock()
	defer fake.captureFoundFileDescriptorsMutex.Unlock()
	fake.CaptureFoundFileDescriptorsStub = stub
}

func (fake *FakeMetricReporter) CaptureFoundFileDescriptorsArgsForCall(i int) int {
	fake.captureFoundFileDescriptorsMutex.RLock()
	defer fake.captureFoundFileDescriptorsMutex.RUnlock()
	argsForCall := fake.captureFoundFileDescriptorsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureHTTPLatency(arg1 time.Duration, arg2 string) {
	fake.captureHTTPLatencyMutex.Lock()
	fake.captureHTTPLatencyArgsForCall = append(fake.captureHTTPLatencyArgsForCall, struct {
		arg1 time.Duration
		arg2 string
	}{arg1, arg2})
	stub := fake.CaptureHTTPLatencyStub
	fake.recordInvocation("CaptureHTTPLatency", []interface{}{arg1, arg2})
	fake.captureHTTPLatencyMutex.Unlock()
	if stub != nil {
		fake.CaptureHTTPLatencyStub(arg1, arg2)
	}
}

func (fake *FakeMetricReporter) CaptureHTTPLatencyCallCount() int {
	fake.captureHTTPLatencyMutex.RLock()
	defer fake.captureHTTPLatencyMutex.RUnlock()
	return len(fake.captureHTTPLatencyArgsForCall)
}

func (fake *FakeMetricReporter) CaptureHTTPLatencyCalls(stub func(time.Duration, string)) {
	fake.captureHTTPLatencyMutex.Lock()
	defer fake.captureHTTPLatencyMutex.Unlock()
	fake.CaptureHTTPLatencyStub = stub
}

func (fake *FakeMetricReporter) CaptureHTTPLatencyArgsForCall(i int) (time.Duration, string) {
	fake.captureHTTPLatencyMutex.RLock()
	defer fake.captureHTTPLatencyMutex.RUnlock()
	argsForCall := fake.captureHTTPLatencyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMetricReporter) CaptureLookupTime(arg1 time.Duration) {
	fake.captureLookupTimeMutex.Lock()
	fake.captureLookupTimeArgsForCall = append(fake.captureLookupTimeArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	stub := fake.CaptureLookupTimeStub
	fake.recordInvocation("CaptureLookupTime", []interface{}{arg1})
	fake.captureLookupTimeMutex.Unlock()
	if stub != nil {
		fake.CaptureLookupTimeStub(arg1)
	}
}

func (fake *FakeMetricReporter) CaptureLookupTimeCallCount() int {
	fake.captureLookupTimeMutex.RLock()
	defer fake.captureLookupTimeMutex.RUnlock()
	return len(fake.captureLookupTimeArgsForCall)
}

func (fake *FakeMetricReporter) CaptureLookupTimeCalls(stub func(time.Duration)) {
	fake.captureLookupTimeMutex.Lock()
	defer fake.captureLookupTimeMutex.Unlock()
	fake.CaptureLookupTimeStub = stub
}

func (fake *FakeMetricReporter) CaptureLookupTimeArgsForCall(i int) time.Duration {
	fake.captureLookupTimeMutex.RLock()
	defer fake.captureLookupTimeMutex.RUnlock()
	argsForCall := fake.captureLookupTimeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureNATSBufferedMessages(arg1 int) {
	fake.captureNATSBufferedMessagesMutex.Lock()
	fake.captureNATSBufferedMessagesArgsForCall = append(fake.captureNATSBufferedMessagesArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.CaptureNATSBufferedMessagesStub
	fake.recordInvocation("CaptureNATSBufferedMessages", []interface{}{arg1})
	fake.captureNATSBufferedMessagesMutex.Unlock()
	if stub != nil {
		fake.CaptureNATSBufferedMessagesStub(arg1)
	}
}

func (fake *FakeMetricReporter) CaptureNATSBufferedMessagesCallCount() int {
	fake.captureNATSBufferedMessagesMutex.RLock()
	defer fake.captureNATSBufferedMessagesMutex.RUnlock()
	return len(fake.captureNATSBufferedMessagesArgsForCall)
}

func (fake *FakeMetricReporter) CaptureNATSBufferedMessagesCalls(stub func(int)) {
	fake.captureNATSBufferedMessagesMutex.Lock()
	defer fake.captureNATSBufferedMessagesMutex.Unlock()
	fake.CaptureNATSBufferedMessagesStub = stub
}

func (fake *FakeMetricReporter) CaptureNATSBufferedMessagesArgsForCall(i int) int {
	fake.captureNATSBufferedMessagesMutex.RLock()
	defer fake.captureNATSBufferedMessagesMutex.RUnlock()
	argsForCall := fake.captureNATSBufferedMessagesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureNATSDroppedMessages(arg1 int) {
	fake.captureNATSDroppedMessagesMutex.Lock()
	fake.captureNATSDroppedMessagesArgsForCall = append(fake.captureNATSDroppedMessagesArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.CaptureNATSDroppedMessagesStub
	fake.recordInvocation("CaptureNATSDroppedMessages", []interface{}{arg1})
	fake.captureNATSDroppedMessagesMutex.Unlock()
	if stub != nil {
		fake.CaptureNATSDroppedMessagesStub(arg1)
	}
}

func (fake *FakeMetricReporter) CaptureNATSDroppedMessagesCallCount() int {
	fake.captureNATSDroppedMessagesMutex.RLock()
	defer fake.captureNATSDroppedMessagesMutex.RUnlock()
	return len(fake.captureNATSDroppedMessagesArgsForCall)
}

func (fake *FakeMetricReporter) CaptureNATSDroppedMessagesCalls(stub func(int)) {
	fake.captureNATSDroppedMessagesMutex.Lock()
	defer fake.captureNATSDroppedMessagesMutex.Unlock()
	fake.CaptureNATSDroppedMessagesStub = stub
}

func (fake *FakeMetricReporter) CaptureNATSDroppedMessagesArgsForCall(i int) int {
	fake.captureNATSDroppedMessagesMutex.RLock()
	defer fake.captureNATSDroppedMessagesMutex.RUnlock()
	argsForCall := fake.captureNATSDroppedMessagesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureRegistryMessage(arg1 metrics.ComponentTagged, arg2 string) {
	fake.captureRegistryMessageMutex.Lock()
	fake.captureRegistryMessageArgsForCall = append(fake.captureRegistryMessageArgsForCall, struct {
		arg1 metrics.ComponentTagged
		arg2 string
	}{arg1, arg2})
	stub := fake.CaptureRegistryMessageStub
	fake.recordInvocation("CaptureRegistryMessage", []interface{}{arg1, arg2})
	fake.captureRegistryMessageMutex.Unlock()
	if stub != nil {
		fake.CaptureRegistryMessageStub(arg1, arg2)
	}
}

func (fake *FakeMetricReporter) CaptureRegistryMessageCallCount() int {
	fake.captureRegistryMessageMutex.RLock()
	defer fake.captureRegistryMessageMutex.RUnlock()
	return len(fake.captureRegistryMessageArgsForCall)
}

func (fake *FakeMetricReporter) CaptureRegistryMessageCalls(stub func(metrics.ComponentTagged, string)) {
	fake.captureRegistryMessageMutex.Lock()
	defer fake.captureRegistryMessageMutex.Unlock()
	fake.CaptureRegistryMessageStub = stub
}

func (fake *FakeMetricReporter) CaptureRegistryMessageArgsForCall(i int) (metrics.ComponentTagged, string) {
	fake.captureRegistryMessageMutex.RLock()
	defer fake.captureRegistryMessageMutex.RUnlock()
	argsForCall := fake.captureRegistryMessageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMetricReporter) CaptureRouteRegistrationLatency(arg1 time.Duration) {
	fake.captureRouteRegistrationLatencyMutex.Lock()
	fake.captureRouteRegistrationLatencyArgsForCall = append(fake.captureRouteRegistrationLatencyArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	stub := fake.CaptureRouteRegistrationLatencyStub
	fake.recordInvocation("CaptureRouteRegistrationLatency", []interface{}{arg1})
	fake.captureRouteRegistrationLatencyMutex.Unlock()
	if stub != nil {
		fake.CaptureRouteRegistrationLatencyStub(arg1)
	}
}

func (fake *FakeMetricReporter) CaptureRouteRegistrationLatencyCallCount() int {
	fake.captureRouteRegistrationLatencyMutex.RLock()
	defer fake.captureRouteRegistrationLatencyMutex.RUnlock()
	return len(fake.captureRouteRegistrationLatencyArgsForCall)
}

func (fake *FakeMetricReporter) CaptureRouteRegistrationLatencyCalls(stub func(time.Duration)) {
	fake.captureRouteRegistrationLatencyMutex.Lock()
	defer fake.captureRouteRegistrationLatencyMutex.Unlock()
	fake.CaptureRouteRegistrationLatencyStub = stub
}

func (fake *FakeMetricReporter) CaptureRouteRegistrationLatencyArgsForCall(i int) time.Duration {
	fake.captureRouteRegistrationLatencyMutex.RLock()
	defer fake.captureRouteRegistrationLatencyMutex.RUnlock()
	argsForCall := fake.captureRouteRegistrationLatencyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureRouteServiceResponse(arg1 *http.Response) {
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

func (fake *FakeMetricReporter) CaptureRouteServiceResponseCallCount() int {
	fake.captureRouteServiceResponseMutex.RLock()
	defer fake.captureRouteServiceResponseMutex.RUnlock()
	return len(fake.captureRouteServiceResponseArgsForCall)
}

func (fake *FakeMetricReporter) CaptureRouteServiceResponseCalls(stub func(*http.Response)) {
	fake.captureRouteServiceResponseMutex.Lock()
	defer fake.captureRouteServiceResponseMutex.Unlock()
	fake.CaptureRouteServiceResponseStub = stub
}

func (fake *FakeMetricReporter) CaptureRouteServiceResponseArgsForCall(i int) *http.Response {
	fake.captureRouteServiceResponseMutex.RLock()
	defer fake.captureRouteServiceResponseMutex.RUnlock()
	argsForCall := fake.captureRouteServiceResponseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureRouteStats(arg1 int, arg2 int64) {
	fake.captureRouteStatsMutex.Lock()
	fake.captureRouteStatsArgsForCall = append(fake.captureRouteStatsArgsForCall, struct {
		arg1 int
		arg2 int64
	}{arg1, arg2})
	stub := fake.CaptureRouteStatsStub
	fake.recordInvocation("CaptureRouteStats", []interface{}{arg1, arg2})
	fake.captureRouteStatsMutex.Unlock()
	if stub != nil {
		fake.CaptureRouteStatsStub(arg1, arg2)
	}
}

func (fake *FakeMetricReporter) CaptureRouteStatsCallCount() int {
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	return len(fake.captureRouteStatsArgsForCall)
}

func (fake *FakeMetricReporter) CaptureRouteStatsCalls(stub func(int, int64)) {
	fake.captureRouteStatsMutex.Lock()
	defer fake.captureRouteStatsMutex.Unlock()
	fake.CaptureRouteStatsStub = stub
}

func (fake *FakeMetricReporter) CaptureRouteStatsArgsForCall(i int) (int, int64) {
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	argsForCall := fake.captureRouteStatsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMetricReporter) CaptureRoutesPruned(arg1 uint64) {
	fake.captureRoutesPrunedMutex.Lock()
	fake.captureRoutesPrunedArgsForCall = append(fake.captureRoutesPrunedArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.CaptureRoutesPrunedStub
	fake.recordInvocation("CaptureRoutesPruned", []interface{}{arg1})
	fake.captureRoutesPrunedMutex.Unlock()
	if stub != nil {
		fake.CaptureRoutesPrunedStub(arg1)
	}
}

func (fake *FakeMetricReporter) CaptureRoutesPrunedCallCount() int {
	fake.captureRoutesPrunedMutex.RLock()
	defer fake.captureRoutesPrunedMutex.RUnlock()
	return len(fake.captureRoutesPrunedArgsForCall)
}

func (fake *FakeMetricReporter) CaptureRoutesPrunedCalls(stub func(uint64)) {
	fake.captureRoutesPrunedMutex.Lock()
	defer fake.captureRoutesPrunedMutex.Unlock()
	fake.CaptureRoutesPrunedStub = stub
}

func (fake *FakeMetricReporter) CaptureRoutesPrunedArgsForCall(i int) uint64 {
	fake.captureRoutesPrunedMutex.RLock()
	defer fake.captureRoutesPrunedMutex.RUnlock()
	argsForCall := fake.captureRoutesPrunedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureRoutingRequest(arg1 *route.Endpoint) {
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

func (fake *FakeMetricReporter) CaptureRoutingRequestCallCount() int {
	fake.captureRoutingRequestMutex.RLock()
	defer fake.captureRoutingRequestMutex.RUnlock()
	return len(fake.captureRoutingRequestArgsForCall)
}

func (fake *FakeMetricReporter) CaptureRoutingRequestCalls(stub func(*route.Endpoint)) {
	fake.captureRoutingRequestMutex.Lock()
	defer fake.captureRoutingRequestMutex.Unlock()
	fake.CaptureRoutingRequestStub = stub
}

func (fake *FakeMetricReporter) CaptureRoutingRequestArgsForCall(i int) *route.Endpoint {
	fake.captureRoutingRequestMutex.RLock()
	defer fake.captureRoutingRequestMutex.RUnlock()
	argsForCall := fake.captureRoutingRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureRoutingResponse(arg1 int) {
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

func (fake *FakeMetricReporter) CaptureRoutingResponseCallCount() int {
	fake.captureRoutingResponseMutex.RLock()
	defer fake.captureRoutingResponseMutex.RUnlock()
	return len(fake.captureRoutingResponseArgsForCall)
}

func (fake *FakeMetricReporter) CaptureRoutingResponseCalls(stub func(int)) {
	fake.captureRoutingResponseMutex.Lock()
	defer fake.captureRoutingResponseMutex.Unlock()
	fake.CaptureRoutingResponseStub = stub
}

func (fake *FakeMetricReporter) CaptureRoutingResponseArgsForCall(i int) int {
	fake.captureRoutingResponseMutex.RLock()
	defer fake.captureRoutingResponseMutex.RUnlock()
	argsForCall := fake.captureRoutingResponseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureRoutingResponseLatency(arg1 *route.Endpoint, arg2 int, arg3 time.Time, arg4 time.Duration) {
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

func (fake *FakeMetricReporter) CaptureRoutingResponseLatencyCallCount() int {
	fake.captureRoutingResponseLatencyMutex.RLock()
	defer fake.captureRoutingResponseLatencyMutex.RUnlock()
	return len(fake.captureRoutingResponseLatencyArgsForCall)
}

func (fake *FakeMetricReporter) CaptureRoutingResponseLatencyCalls(stub func(*route.Endpoint, int, time.Time, time.Duration)) {
	fake.captureRoutingResponseLatencyMutex.Lock()
	defer fake.captureRoutingResponseLatencyMutex.Unlock()
	fake.CaptureRoutingResponseLatencyStub = stub
}

func (fake *FakeMetricReporter) CaptureRoutingResponseLatencyArgsForCall(i int) (*route.Endpoint, int, time.Time, time.Duration) {
	fake.captureRoutingResponseLatencyMutex.RLock()
	defer fake.captureRoutingResponseLatencyMutex.RUnlock()
	argsForCall := fake.captureRoutingResponseLatencyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeMetricReporter) CaptureUnregistryMessage(arg1 metrics.ComponentTagged) {
	fake.captureUnregistryMessageMutex.Lock()
	fake.captureUnregistryMessageArgsForCall = append(fake.captureUnregistryMessageArgsForCall, struct {
		arg1 metrics.ComponentTagged
	}{arg1})
	stub := fake.CaptureUnregistryMessageStub
	fake.recordInvocation("CaptureUnregistryMessage", []interface{}{arg1})
	fake.captureUnregistryMessageMutex.Unlock()
	if stub != nil {
		fake.CaptureUnregistryMessageStub(arg1)
	}
}

func (fake *FakeMetricReporter) CaptureUnregistryMessageCallCount() int {
	fake.captureUnregistryMessageMutex.RLock()
	defer fake.captureUnregistryMessageMutex.RUnlock()
	return len(fake.captureUnregistryMessageArgsForCall)
}

func (fake *FakeMetricReporter) CaptureUnregistryMessageCalls(stub func(metrics.ComponentTagged)) {
	fake.captureUnregistryMessageMutex.Lock()
	defer fake.captureUnregistryMessageMutex.Unlock()
	fake.CaptureUnregistryMessageStub = stub
}

func (fake *FakeMetricReporter) CaptureUnregistryMessageArgsForCall(i int) metrics.ComponentTagged {
	fake.captureUnregistryMessageMutex.RLock()
	defer fake.captureUnregistryMessageMutex.RUnlock()
	argsForCall := fake.captureUnregistryMessageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetricReporter) CaptureWebSocketFailure() {
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

func (fake *FakeMetricReporter) CaptureWebSocketFailureCallCount() int {
	fake.captureWebSocketFailureMutex.RLock()
	defer fake.captureWebSocketFailureMutex.RUnlock()
	return len(fake.captureWebSocketFailureArgsForCall)
}

func (fake *FakeMetricReporter) CaptureWebSocketFailureCalls(stub func()) {
	fake.captureWebSocketFailureMutex.Lock()
	defer fake.captureWebSocketFailureMutex.Unlock()
	fake.CaptureWebSocketFailureStub = stub
}

func (fake *FakeMetricReporter) CaptureWebSocketUpdate() {
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

func (fake *FakeMetricReporter) CaptureWebSocketUpdateCallCount() int {
	fake.captureWebSocketUpdateMutex.RLock()
	defer fake.captureWebSocketUpdateMutex.RUnlock()
	return len(fake.captureWebSocketUpdateArgsForCall)
}

func (fake *FakeMetricReporter) CaptureWebSocketUpdateCalls(stub func()) {
	fake.captureWebSocketUpdateMutex.Lock()
	defer fake.captureWebSocketUpdateMutex.Unlock()
	fake.CaptureWebSocketUpdateStub = stub
}

func (fake *FakeMetricReporter) UnmuzzleRouteRegistrationLatency() {
	fake.unmuzzleRouteRegistrationLatencyMutex.Lock()
	fake.unmuzzleRouteRegistrationLatencyArgsForCall = append(fake.unmuzzleRouteRegistrationLatencyArgsForCall, struct {
	}{})
	stub := fake.UnmuzzleRouteRegistrationLatencyStub
	fake.recordInvocation("UnmuzzleRouteRegistrationLatency", []interface{}{})
	fake.unmuzzleRouteRegistrationLatencyMutex.Unlock()
	if stub != nil {
		fake.UnmuzzleRouteRegistrationLatencyStub()
	}
}

func (fake *FakeMetricReporter) UnmuzzleRouteRegistrationLatencyCallCount() int {
	fake.unmuzzleRouteRegistrationLatencyMutex.RLock()
	defer fake.unmuzzleRouteRegistrationLatencyMutex.RUnlock()
	return len(fake.unmuzzleRouteRegistrationLatencyArgsForCall)
}

func (fake *FakeMetricReporter) UnmuzzleRouteRegistrationLatencyCalls(stub func()) {
	fake.unmuzzleRouteRegistrationLatencyMutex.Lock()
	defer fake.unmuzzleRouteRegistrationLatencyMutex.Unlock()
	fake.UnmuzzleRouteRegistrationLatencyStub = stub
}

func (fake *FakeMetricReporter) Invocations() map[string][][]interface{} {
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
	fake.captureFoundFileDescriptorsMutex.RLock()
	defer fake.captureFoundFileDescriptorsMutex.RUnlock()
	fake.captureHTTPLatencyMutex.RLock()
	defer fake.captureHTTPLatencyMutex.RUnlock()
	fake.captureLookupTimeMutex.RLock()
	defer fake.captureLookupTimeMutex.RUnlock()
	fake.captureNATSBufferedMessagesMutex.RLock()
	defer fake.captureNATSBufferedMessagesMutex.RUnlock()
	fake.captureNATSDroppedMessagesMutex.RLock()
	defer fake.captureNATSDroppedMessagesMutex.RUnlock()
	fake.captureRegistryMessageMutex.RLock()
	defer fake.captureRegistryMessageMutex.RUnlock()
	fake.captureRouteRegistrationLatencyMutex.RLock()
	defer fake.captureRouteRegistrationLatencyMutex.RUnlock()
	fake.captureRouteServiceResponseMutex.RLock()
	defer fake.captureRouteServiceResponseMutex.RUnlock()
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	fake.captureRoutesPrunedMutex.RLock()
	defer fake.captureRoutesPrunedMutex.RUnlock()
	fake.captureRoutingRequestMutex.RLock()
	defer fake.captureRoutingRequestMutex.RUnlock()
	fake.captureRoutingResponseMutex.RLock()
	defer fake.captureRoutingResponseMutex.RUnlock()
	fake.captureRoutingResponseLatencyMutex.RLock()
	defer fake.captureRoutingResponseLatencyMutex.RUnlock()
	fake.captureUnregistryMessageMutex.RLock()
	defer fake.captureUnregistryMessageMutex.RUnlock()
	fake.captureWebSocketFailureMutex.RLock()
	defer fake.captureWebSocketFailureMutex.RUnlock()
	fake.captureWebSocketUpdateMutex.RLock()
	defer fake.captureWebSocketUpdateMutex.RUnlock()
	fake.unmuzzleRouteRegistrationLatencyMutex.RLock()
	defer fake.unmuzzleRouteRegistrationLatencyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMetricReporter) recordInvocation(key string, args []interface{}) {
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

var _ metrics.MetricReporter = new(FakeMetricReporter)
