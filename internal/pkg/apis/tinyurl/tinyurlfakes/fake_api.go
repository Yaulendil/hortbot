// Code generated by counterfeiter. DO NOT EDIT.
package tinyurlfakes

import (
	"context"
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/apis/tinyurl"
)

type FakeAPI struct {
	ShortenStub        func(context.Context, string) (string, error)
	shortenMutex       sync.RWMutex
	shortenArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	shortenReturns struct {
		result1 string
		result2 error
	}
	shortenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) Shorten(arg1 context.Context, arg2 string) (string, error) {
	fake.shortenMutex.Lock()
	ret, specificReturn := fake.shortenReturnsOnCall[len(fake.shortenArgsForCall)]
	fake.shortenArgsForCall = append(fake.shortenArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Shorten", []interface{}{arg1, arg2})
	fake.shortenMutex.Unlock()
	if fake.ShortenStub != nil {
		return fake.ShortenStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.shortenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) ShortenCallCount() int {
	fake.shortenMutex.RLock()
	defer fake.shortenMutex.RUnlock()
	return len(fake.shortenArgsForCall)
}

func (fake *FakeAPI) ShortenCalls(stub func(context.Context, string) (string, error)) {
	fake.shortenMutex.Lock()
	defer fake.shortenMutex.Unlock()
	fake.ShortenStub = stub
}

func (fake *FakeAPI) ShortenArgsForCall(i int) (context.Context, string) {
	fake.shortenMutex.RLock()
	defer fake.shortenMutex.RUnlock()
	argsForCall := fake.shortenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) ShortenReturns(result1 string, result2 error) {
	fake.shortenMutex.Lock()
	defer fake.shortenMutex.Unlock()
	fake.ShortenStub = nil
	fake.shortenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) ShortenReturnsOnCall(i int, result1 string, result2 error) {
	fake.shortenMutex.Lock()
	defer fake.shortenMutex.Unlock()
	fake.ShortenStub = nil
	if fake.shortenReturnsOnCall == nil {
		fake.shortenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.shortenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.shortenMutex.RLock()
	defer fake.shortenMutex.RUnlock()
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

var _ tinyurl.API = new(FakeAPI)