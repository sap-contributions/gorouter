// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/gorouter/proxy/fails"
)

type Classifier struct {
	ClassifyStub        func(error) bool
	classifyMutex       sync.RWMutex
	classifyArgsForCall []struct {
		arg1 error
	}
	classifyReturns struct {
		result1 bool
	}
	classifyReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Classifier) Classify(arg1 error) bool {
	fake.classifyMutex.Lock()
	ret, specificReturn := fake.classifyReturnsOnCall[len(fake.classifyArgsForCall)]
	fake.classifyArgsForCall = append(fake.classifyArgsForCall, struct {
		arg1 error
	}{arg1})
	stub := fake.ClassifyStub
	fakeReturns := fake.classifyReturns
	fake.recordInvocation("Classify", []interface{}{arg1})
	fake.classifyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Classifier) ClassifyCallCount() int {
	fake.classifyMutex.RLock()
	defer fake.classifyMutex.RUnlock()
	return len(fake.classifyArgsForCall)
}

func (fake *Classifier) ClassifyCalls(stub func(error) bool) {
	fake.classifyMutex.Lock()
	defer fake.classifyMutex.Unlock()
	fake.ClassifyStub = stub
}

func (fake *Classifier) ClassifyArgsForCall(i int) error {
	fake.classifyMutex.RLock()
	defer fake.classifyMutex.RUnlock()
	argsForCall := fake.classifyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Classifier) ClassifyReturns(result1 bool) {
	fake.classifyMutex.Lock()
	defer fake.classifyMutex.Unlock()
	fake.ClassifyStub = nil
	fake.classifyReturns = struct {
		result1 bool
	}{result1}
}

func (fake *Classifier) ClassifyReturnsOnCall(i int, result1 bool) {
	fake.classifyMutex.Lock()
	defer fake.classifyMutex.Unlock()
	fake.ClassifyStub = nil
	if fake.classifyReturnsOnCall == nil {
		fake.classifyReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.classifyReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *Classifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.classifyMutex.RLock()
	defer fake.classifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Classifier) recordInvocation(key string, args []interface{}) {
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

var _ fails.Classifier = new(Classifier)
