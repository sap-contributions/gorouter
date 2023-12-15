// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/gorouter/route"
)

type FakeEndpointIterator struct {
	EndpointFailedStub        func(error)
	endpointFailedMutex       sync.RWMutex
	endpointFailedArgsForCall []struct {
		arg1 error
	}
	NextStub        func(int) *route.Endpoint
	nextMutex       sync.RWMutex
	nextArgsForCall []struct {
		arg1 int
	}
	nextReturns struct {
		result1 *route.Endpoint
	}
	nextReturnsOnCall map[int]struct {
		result1 *route.Endpoint
	}
	PostRequestStub        func(*route.Endpoint)
	postRequestMutex       sync.RWMutex
	postRequestArgsForCall []struct {
		arg1 *route.Endpoint
	}
	PreRequestStub        func(*route.Endpoint)
	preRequestMutex       sync.RWMutex
	preRequestArgsForCall []struct {
		arg1 *route.Endpoint
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEndpointIterator) EndpointFailed(arg1 error) {
	fake.endpointFailedMutex.Lock()
	fake.endpointFailedArgsForCall = append(fake.endpointFailedArgsForCall, struct {
		arg1 error
	}{arg1})
	stub := fake.EndpointFailedStub
	fake.recordInvocation("EndpointFailed", []interface{}{arg1})
	fake.endpointFailedMutex.Unlock()
	if stub != nil {
		fake.EndpointFailedStub(arg1)
	}
}

func (fake *FakeEndpointIterator) EndpointFailedCallCount() int {
	fake.endpointFailedMutex.RLock()
	defer fake.endpointFailedMutex.RUnlock()
	return len(fake.endpointFailedArgsForCall)
}

func (fake *FakeEndpointIterator) EndpointFailedCalls(stub func(error)) {
	fake.endpointFailedMutex.Lock()
	defer fake.endpointFailedMutex.Unlock()
	fake.EndpointFailedStub = stub
}

func (fake *FakeEndpointIterator) EndpointFailedArgsForCall(i int) error {
	fake.endpointFailedMutex.RLock()
	defer fake.endpointFailedMutex.RUnlock()
	argsForCall := fake.endpointFailedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEndpointIterator) Next(arg1 int) *route.Endpoint {
	fake.nextMutex.Lock()
	ret, specificReturn := fake.nextReturnsOnCall[len(fake.nextArgsForCall)]
	fake.nextArgsForCall = append(fake.nextArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.NextStub
	fakeReturns := fake.nextReturns
	fake.recordInvocation("Next", []interface{}{arg1})
	fake.nextMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEndpointIterator) NextCallCount() int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	return len(fake.nextArgsForCall)
}

func (fake *FakeEndpointIterator) NextCalls(stub func(int) *route.Endpoint) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = stub
}

func (fake *FakeEndpointIterator) NextArgsForCall(i int) int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	argsForCall := fake.nextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEndpointIterator) NextReturns(result1 *route.Endpoint) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	fake.nextReturns = struct {
		result1 *route.Endpoint
	}{result1}
}

func (fake *FakeEndpointIterator) NextReturnsOnCall(i int, result1 *route.Endpoint) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	if fake.nextReturnsOnCall == nil {
		fake.nextReturnsOnCall = make(map[int]struct {
			result1 *route.Endpoint
		})
	}
	fake.nextReturnsOnCall[i] = struct {
		result1 *route.Endpoint
	}{result1}
}

func (fake *FakeEndpointIterator) PostRequest(arg1 *route.Endpoint) {
	fake.postRequestMutex.Lock()
	fake.postRequestArgsForCall = append(fake.postRequestArgsForCall, struct {
		arg1 *route.Endpoint
	}{arg1})
	stub := fake.PostRequestStub
	fake.recordInvocation("PostRequest", []interface{}{arg1})
	fake.postRequestMutex.Unlock()
	if stub != nil {
		fake.PostRequestStub(arg1)
	}
}

func (fake *FakeEndpointIterator) PostRequestCallCount() int {
	fake.postRequestMutex.RLock()
	defer fake.postRequestMutex.RUnlock()
	return len(fake.postRequestArgsForCall)
}

func (fake *FakeEndpointIterator) PostRequestCalls(stub func(*route.Endpoint)) {
	fake.postRequestMutex.Lock()
	defer fake.postRequestMutex.Unlock()
	fake.PostRequestStub = stub
}

func (fake *FakeEndpointIterator) PostRequestArgsForCall(i int) *route.Endpoint {
	fake.postRequestMutex.RLock()
	defer fake.postRequestMutex.RUnlock()
	argsForCall := fake.postRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEndpointIterator) PreRequest(arg1 *route.Endpoint) {
	fake.preRequestMutex.Lock()
	fake.preRequestArgsForCall = append(fake.preRequestArgsForCall, struct {
		arg1 *route.Endpoint
	}{arg1})
	stub := fake.PreRequestStub
	fake.recordInvocation("PreRequest", []interface{}{arg1})
	fake.preRequestMutex.Unlock()
	if stub != nil {
		fake.PreRequestStub(arg1)
	}
}

func (fake *FakeEndpointIterator) PreRequestCallCount() int {
	fake.preRequestMutex.RLock()
	defer fake.preRequestMutex.RUnlock()
	return len(fake.preRequestArgsForCall)
}

func (fake *FakeEndpointIterator) PreRequestCalls(stub func(*route.Endpoint)) {
	fake.preRequestMutex.Lock()
	defer fake.preRequestMutex.Unlock()
	fake.PreRequestStub = stub
}

func (fake *FakeEndpointIterator) PreRequestArgsForCall(i int) *route.Endpoint {
	fake.preRequestMutex.RLock()
	defer fake.preRequestMutex.RUnlock()
	argsForCall := fake.preRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEndpointIterator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.endpointFailedMutex.RLock()
	defer fake.endpointFailedMutex.RUnlock()
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	fake.postRequestMutex.RLock()
	defer fake.postRequestMutex.RUnlock()
	fake.preRequestMutex.RLock()
	defer fake.preRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEndpointIterator) recordInvocation(key string, args []interface{}) {
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

var _ route.EndpointIterator = new(FakeEndpointIterator)
