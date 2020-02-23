// Code generated by counterfeiter. DO NOT EDIT.
package simplefakes

import (
	"context"
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/apiclient/simple"
)

type FakeAPI struct {
	PlaintextStub        func(context.Context, string) (string, error)
	plaintextMutex       sync.RWMutex
	plaintextArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	plaintextReturns struct {
		result1 string
		result2 error
	}
	plaintextReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) Plaintext(arg1 context.Context, arg2 string) (string, error) {
	fake.plaintextMutex.Lock()
	ret, specificReturn := fake.plaintextReturnsOnCall[len(fake.plaintextArgsForCall)]
	fake.plaintextArgsForCall = append(fake.plaintextArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Plaintext", []interface{}{arg1, arg2})
	fake.plaintextMutex.Unlock()
	if fake.PlaintextStub != nil {
		return fake.PlaintextStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.plaintextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) PlaintextCallCount() int {
	fake.plaintextMutex.RLock()
	defer fake.plaintextMutex.RUnlock()
	return len(fake.plaintextArgsForCall)
}

func (fake *FakeAPI) PlaintextCalls(stub func(context.Context, string) (string, error)) {
	fake.plaintextMutex.Lock()
	defer fake.plaintextMutex.Unlock()
	fake.PlaintextStub = stub
}

func (fake *FakeAPI) PlaintextArgsForCall(i int) (context.Context, string) {
	fake.plaintextMutex.RLock()
	defer fake.plaintextMutex.RUnlock()
	argsForCall := fake.plaintextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) PlaintextReturns(result1 string, result2 error) {
	fake.plaintextMutex.Lock()
	defer fake.plaintextMutex.Unlock()
	fake.PlaintextStub = nil
	fake.plaintextReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) PlaintextReturnsOnCall(i int, result1 string, result2 error) {
	fake.plaintextMutex.Lock()
	defer fake.plaintextMutex.Unlock()
	fake.PlaintextStub = nil
	if fake.plaintextReturnsOnCall == nil {
		fake.plaintextReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.plaintextReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.plaintextMutex.RLock()
	defer fake.plaintextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAPI) recordInvocation(key string, args []interface{}) {
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

var _ simple.API = new(FakeAPI)