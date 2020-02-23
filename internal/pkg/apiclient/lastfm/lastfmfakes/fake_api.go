// Code generated by counterfeiter. DO NOT EDIT.
package lastfmfakes

import (
	"context"
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/apiclient/lastfm"
)

type FakeAPI struct {
	RecentTracksStub        func(context.Context, string, int) ([]lastfm.Track, error)
	recentTracksMutex       sync.RWMutex
	recentTracksArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 int
	}
	recentTracksReturns struct {
		result1 []lastfm.Track
		result2 error
	}
	recentTracksReturnsOnCall map[int]struct {
		result1 []lastfm.Track
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) RecentTracks(arg1 context.Context, arg2 string, arg3 int) ([]lastfm.Track, error) {
	fake.recentTracksMutex.Lock()
	ret, specificReturn := fake.recentTracksReturnsOnCall[len(fake.recentTracksArgsForCall)]
	fake.recentTracksArgsForCall = append(fake.recentTracksArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("RecentTracks", []interface{}{arg1, arg2, arg3})
	fake.recentTracksMutex.Unlock()
	if fake.RecentTracksStub != nil {
		return fake.RecentTracksStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.recentTracksReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) RecentTracksCallCount() int {
	fake.recentTracksMutex.RLock()
	defer fake.recentTracksMutex.RUnlock()
	return len(fake.recentTracksArgsForCall)
}

func (fake *FakeAPI) RecentTracksCalls(stub func(context.Context, string, int) ([]lastfm.Track, error)) {
	fake.recentTracksMutex.Lock()
	defer fake.recentTracksMutex.Unlock()
	fake.RecentTracksStub = stub
}

func (fake *FakeAPI) RecentTracksArgsForCall(i int) (context.Context, string, int) {
	fake.recentTracksMutex.RLock()
	defer fake.recentTracksMutex.RUnlock()
	argsForCall := fake.recentTracksArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAPI) RecentTracksReturns(result1 []lastfm.Track, result2 error) {
	fake.recentTracksMutex.Lock()
	defer fake.recentTracksMutex.Unlock()
	fake.RecentTracksStub = nil
	fake.recentTracksReturns = struct {
		result1 []lastfm.Track
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) RecentTracksReturnsOnCall(i int, result1 []lastfm.Track, result2 error) {
	fake.recentTracksMutex.Lock()
	defer fake.recentTracksMutex.Unlock()
	fake.RecentTracksStub = nil
	if fake.recentTracksReturnsOnCall == nil {
		fake.recentTracksReturnsOnCall = make(map[int]struct {
			result1 []lastfm.Track
			result2 error
		})
	}
	fake.recentTracksReturnsOnCall[i] = struct {
		result1 []lastfm.Track
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.recentTracksMutex.RLock()
	defer fake.recentTracksMutex.RUnlock()
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

var _ lastfm.API = new(FakeAPI)