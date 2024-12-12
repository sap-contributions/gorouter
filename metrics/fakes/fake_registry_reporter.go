// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/gorouter/metrics"
)

type FakeRouteRegistryReporter struct {
	CaptureLookupTimeStub        func(time.Duration)
	captureLookupTimeMutex       sync.RWMutex
	captureLookupTimeArgsForCall []struct {
		arg1 time.Duration
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
	CaptureUnregistryMessageStub        func(metrics.ComponentTagged, string)
	captureUnregistryMessageMutex       sync.RWMutex
	captureUnregistryMessageArgsForCall []struct {
		arg1 metrics.ComponentTagged
		arg2 string
	}
	UnmuzzleRouteRegistrationLatencyStub        func()
	unmuzzleRouteRegistrationLatencyMutex       sync.RWMutex
	unmuzzleRouteRegistrationLatencyArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouteRegistryReporter) CaptureLookupTime(arg1 time.Duration) {
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

func (fake *FakeRouteRegistryReporter) CaptureLookupTimeCallCount() int {
	fake.captureLookupTimeMutex.RLock()
	defer fake.captureLookupTimeMutex.RUnlock()
	return len(fake.captureLookupTimeArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureLookupTimeCalls(stub func(time.Duration)) {
	fake.captureLookupTimeMutex.Lock()
	defer fake.captureLookupTimeMutex.Unlock()
	fake.CaptureLookupTimeStub = stub
}

func (fake *FakeRouteRegistryReporter) CaptureLookupTimeArgsForCall(i int) time.Duration {
	fake.captureLookupTimeMutex.RLock()
	defer fake.captureLookupTimeMutex.RUnlock()
	argsForCall := fake.captureLookupTimeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouteRegistryReporter) CaptureRegistryMessage(arg1 metrics.ComponentTagged, arg2 string) {
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

func (fake *FakeRouteRegistryReporter) CaptureRegistryMessageCallCount() int {
	fake.captureRegistryMessageMutex.RLock()
	defer fake.captureRegistryMessageMutex.RUnlock()
	return len(fake.captureRegistryMessageArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureRegistryMessageCalls(stub func(metrics.ComponentTagged, string)) {
	fake.captureRegistryMessageMutex.Lock()
	defer fake.captureRegistryMessageMutex.Unlock()
	fake.CaptureRegistryMessageStub = stub
}

func (fake *FakeRouteRegistryReporter) CaptureRegistryMessageArgsForCall(i int) (metrics.ComponentTagged, string) {
	fake.captureRegistryMessageMutex.RLock()
	defer fake.captureRegistryMessageMutex.RUnlock()
	argsForCall := fake.captureRegistryMessageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRouteRegistryReporter) CaptureRouteRegistrationLatency(arg1 time.Duration) {
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

func (fake *FakeRouteRegistryReporter) CaptureRouteRegistrationLatencyCallCount() int {
	fake.captureRouteRegistrationLatencyMutex.RLock()
	defer fake.captureRouteRegistrationLatencyMutex.RUnlock()
	return len(fake.captureRouteRegistrationLatencyArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureRouteRegistrationLatencyCalls(stub func(time.Duration)) {
	fake.captureRouteRegistrationLatencyMutex.Lock()
	defer fake.captureRouteRegistrationLatencyMutex.Unlock()
	fake.CaptureRouteRegistrationLatencyStub = stub
}

func (fake *FakeRouteRegistryReporter) CaptureRouteRegistrationLatencyArgsForCall(i int) time.Duration {
	fake.captureRouteRegistrationLatencyMutex.RLock()
	defer fake.captureRouteRegistrationLatencyMutex.RUnlock()
	argsForCall := fake.captureRouteRegistrationLatencyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouteRegistryReporter) CaptureRouteStats(arg1 int, arg2 int64) {
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

func (fake *FakeRouteRegistryReporter) CaptureRouteStatsCallCount() int {
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	return len(fake.captureRouteStatsArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureRouteStatsCalls(stub func(int, int64)) {
	fake.captureRouteStatsMutex.Lock()
	defer fake.captureRouteStatsMutex.Unlock()
	fake.CaptureRouteStatsStub = stub
}

func (fake *FakeRouteRegistryReporter) CaptureRouteStatsArgsForCall(i int) (int, int64) {
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	argsForCall := fake.captureRouteStatsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRouteRegistryReporter) CaptureRoutesPruned(arg1 uint64) {
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

func (fake *FakeRouteRegistryReporter) CaptureRoutesPrunedCallCount() int {
	fake.captureRoutesPrunedMutex.RLock()
	defer fake.captureRoutesPrunedMutex.RUnlock()
	return len(fake.captureRoutesPrunedArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureRoutesPrunedCalls(stub func(uint64)) {
	fake.captureRoutesPrunedMutex.Lock()
	defer fake.captureRoutesPrunedMutex.Unlock()
	fake.CaptureRoutesPrunedStub = stub
}

func (fake *FakeRouteRegistryReporter) CaptureRoutesPrunedArgsForCall(i int) uint64 {
	fake.captureRoutesPrunedMutex.RLock()
	defer fake.captureRoutesPrunedMutex.RUnlock()
	argsForCall := fake.captureRoutesPrunedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouteRegistryReporter) CaptureUnregistryMessage(arg1 metrics.ComponentTagged, arg2 string) {
	fake.captureUnregistryMessageMutex.Lock()
	fake.captureUnregistryMessageArgsForCall = append(fake.captureUnregistryMessageArgsForCall, struct {
		arg1 metrics.ComponentTagged
		arg2 string
	}{arg1, arg2})
	stub := fake.CaptureUnregistryMessageStub
	fake.recordInvocation("CaptureUnregistryMessage", []interface{}{arg1, arg2})
	fake.captureUnregistryMessageMutex.Unlock()
	if stub != nil {
		fake.CaptureUnregistryMessageStub(arg1, arg2)
	}
}

func (fake *FakeRouteRegistryReporter) CaptureUnregistryMessageCallCount() int {
	fake.captureUnregistryMessageMutex.RLock()
	defer fake.captureUnregistryMessageMutex.RUnlock()
	return len(fake.captureUnregistryMessageArgsForCall)
}

func (fake *FakeRouteRegistryReporter) CaptureUnregistryMessageCalls(stub func(metrics.ComponentTagged, string)) {
	fake.captureUnregistryMessageMutex.Lock()
	defer fake.captureUnregistryMessageMutex.Unlock()
	fake.CaptureUnregistryMessageStub = stub
}

func (fake *FakeRouteRegistryReporter) CaptureUnregistryMessageArgsForCall(i int) (metrics.ComponentTagged, string) {
	fake.captureUnregistryMessageMutex.RLock()
	defer fake.captureUnregistryMessageMutex.RUnlock()
	argsForCall := fake.captureUnregistryMessageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRouteRegistryReporter) UnmuzzleRouteRegistrationLatency() {
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

func (fake *FakeRouteRegistryReporter) UnmuzzleRouteRegistrationLatencyCallCount() int {
	fake.unmuzzleRouteRegistrationLatencyMutex.RLock()
	defer fake.unmuzzleRouteRegistrationLatencyMutex.RUnlock()
	return len(fake.unmuzzleRouteRegistrationLatencyArgsForCall)
}

func (fake *FakeRouteRegistryReporter) UnmuzzleRouteRegistrationLatencyCalls(stub func()) {
	fake.unmuzzleRouteRegistrationLatencyMutex.Lock()
	defer fake.unmuzzleRouteRegistrationLatencyMutex.Unlock()
	fake.UnmuzzleRouteRegistrationLatencyStub = stub
}

func (fake *FakeRouteRegistryReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.captureLookupTimeMutex.RLock()
	defer fake.captureLookupTimeMutex.RUnlock()
	fake.captureRegistryMessageMutex.RLock()
	defer fake.captureRegistryMessageMutex.RUnlock()
	fake.captureRouteRegistrationLatencyMutex.RLock()
	defer fake.captureRouteRegistrationLatencyMutex.RUnlock()
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	fake.captureRoutesPrunedMutex.RLock()
	defer fake.captureRoutesPrunedMutex.RUnlock()
	fake.captureUnregistryMessageMutex.RLock()
	defer fake.captureUnregistryMessageMutex.RUnlock()
	fake.unmuzzleRouteRegistrationLatencyMutex.RLock()
	defer fake.unmuzzleRouteRegistrationLatencyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRouteRegistryReporter) recordInvocation(key string, args []interface{}) {
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

var _ metrics.RouteRegistryReporter = new(FakeRouteRegistryReporter)
