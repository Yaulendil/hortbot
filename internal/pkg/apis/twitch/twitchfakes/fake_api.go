// Code generated by counterfeiter. DO NOT EDIT.
package twitchfakes

import (
	"context"
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/apis/twitch"
	"golang.org/x/oauth2"
)

type FakeAPI struct {
	AuthCodeURLStub        func(string, ...string) string
	authCodeURLMutex       sync.RWMutex
	authCodeURLArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	authCodeURLReturns struct {
		result1 string
	}
	authCodeURLReturnsOnCall map[int]struct {
		result1 string
	}
	ExchangeStub        func(context.Context, string) (*oauth2.Token, error)
	exchangeMutex       sync.RWMutex
	exchangeArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	exchangeReturns struct {
		result1 *oauth2.Token
		result2 error
	}
	exchangeReturnsOnCall map[int]struct {
		result1 *oauth2.Token
		result2 error
	}
	FollowChannelStub        func(context.Context, int64, *oauth2.Token, int64) (*oauth2.Token, error)
	followChannelMutex       sync.RWMutex
	followChannelArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 int64
	}
	followChannelReturns struct {
		result1 *oauth2.Token
		result2 error
	}
	followChannelReturnsOnCall map[int]struct {
		result1 *oauth2.Token
		result2 error
	}
	GetChannelByIDStub        func(context.Context, int64) (*twitch.Channel, error)
	getChannelByIDMutex       sync.RWMutex
	getChannelByIDArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getChannelByIDReturns struct {
		result1 *twitch.Channel
		result2 error
	}
	getChannelByIDReturnsOnCall map[int]struct {
		result1 *twitch.Channel
		result2 error
	}
	GetChattersStub        func(context.Context, string) (*twitch.Chatters, error)
	getChattersMutex       sync.RWMutex
	getChattersArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getChattersReturns struct {
		result1 *twitch.Chatters
		result2 error
	}
	getChattersReturnsOnCall map[int]struct {
		result1 *twitch.Chatters
		result2 error
	}
	GetCurrentStreamStub        func(context.Context, int64) (*twitch.Stream, error)
	getCurrentStreamMutex       sync.RWMutex
	getCurrentStreamArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getCurrentStreamReturns struct {
		result1 *twitch.Stream
		result2 error
	}
	getCurrentStreamReturnsOnCall map[int]struct {
		result1 *twitch.Stream
		result2 error
	}
	GetUserForTokenStub        func(context.Context, *oauth2.Token) (*twitch.User, *oauth2.Token, error)
	getUserForTokenMutex       sync.RWMutex
	getUserForTokenArgsForCall []struct {
		arg1 context.Context
		arg2 *oauth2.Token
	}
	getUserForTokenReturns struct {
		result1 *twitch.User
		result2 *oauth2.Token
		result3 error
	}
	getUserForTokenReturnsOnCall map[int]struct {
		result1 *twitch.User
		result2 *oauth2.Token
		result3 error
	}
	GetUserForUsernameStub        func(context.Context, string) (*twitch.User, error)
	getUserForUsernameMutex       sync.RWMutex
	getUserForUsernameArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getUserForUsernameReturns struct {
		result1 *twitch.User
		result2 error
	}
	getUserForUsernameReturnsOnCall map[int]struct {
		result1 *twitch.User
		result2 error
	}
	SetChannelGameStub        func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error)
	setChannelGameMutex       sync.RWMutex
	setChannelGameArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}
	setChannelGameReturns struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}
	setChannelGameReturnsOnCall map[int]struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}
	SetChannelStatusStub        func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error)
	setChannelStatusMutex       sync.RWMutex
	setChannelStatusArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}
	setChannelStatusReturns struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}
	setChannelStatusReturnsOnCall map[int]struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) AuthCodeURL(arg1 string, arg2 ...string) string {
	fake.authCodeURLMutex.Lock()
	ret, specificReturn := fake.authCodeURLReturnsOnCall[len(fake.authCodeURLArgsForCall)]
	fake.authCodeURLArgsForCall = append(fake.authCodeURLArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	fake.recordInvocation("AuthCodeURL", []interface{}{arg1, arg2})
	fake.authCodeURLMutex.Unlock()
	if fake.AuthCodeURLStub != nil {
		return fake.AuthCodeURLStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.authCodeURLReturns
	return fakeReturns.result1
}

func (fake *FakeAPI) AuthCodeURLCallCount() int {
	fake.authCodeURLMutex.RLock()
	defer fake.authCodeURLMutex.RUnlock()
	return len(fake.authCodeURLArgsForCall)
}

func (fake *FakeAPI) AuthCodeURLCalls(stub func(string, ...string) string) {
	fake.authCodeURLMutex.Lock()
	defer fake.authCodeURLMutex.Unlock()
	fake.AuthCodeURLStub = stub
}

func (fake *FakeAPI) AuthCodeURLArgsForCall(i int) (string, []string) {
	fake.authCodeURLMutex.RLock()
	defer fake.authCodeURLMutex.RUnlock()
	argsForCall := fake.authCodeURLArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) AuthCodeURLReturns(result1 string) {
	fake.authCodeURLMutex.Lock()
	defer fake.authCodeURLMutex.Unlock()
	fake.AuthCodeURLStub = nil
	fake.authCodeURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAPI) AuthCodeURLReturnsOnCall(i int, result1 string) {
	fake.authCodeURLMutex.Lock()
	defer fake.authCodeURLMutex.Unlock()
	fake.AuthCodeURLStub = nil
	if fake.authCodeURLReturnsOnCall == nil {
		fake.authCodeURLReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.authCodeURLReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAPI) Exchange(arg1 context.Context, arg2 string) (*oauth2.Token, error) {
	fake.exchangeMutex.Lock()
	ret, specificReturn := fake.exchangeReturnsOnCall[len(fake.exchangeArgsForCall)]
	fake.exchangeArgsForCall = append(fake.exchangeArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Exchange", []interface{}{arg1, arg2})
	fake.exchangeMutex.Unlock()
	if fake.ExchangeStub != nil {
		return fake.ExchangeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.exchangeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) ExchangeCallCount() int {
	fake.exchangeMutex.RLock()
	defer fake.exchangeMutex.RUnlock()
	return len(fake.exchangeArgsForCall)
}

func (fake *FakeAPI) ExchangeCalls(stub func(context.Context, string) (*oauth2.Token, error)) {
	fake.exchangeMutex.Lock()
	defer fake.exchangeMutex.Unlock()
	fake.ExchangeStub = stub
}

func (fake *FakeAPI) ExchangeArgsForCall(i int) (context.Context, string) {
	fake.exchangeMutex.RLock()
	defer fake.exchangeMutex.RUnlock()
	argsForCall := fake.exchangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) ExchangeReturns(result1 *oauth2.Token, result2 error) {
	fake.exchangeMutex.Lock()
	defer fake.exchangeMutex.Unlock()
	fake.ExchangeStub = nil
	fake.exchangeReturns = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) ExchangeReturnsOnCall(i int, result1 *oauth2.Token, result2 error) {
	fake.exchangeMutex.Lock()
	defer fake.exchangeMutex.Unlock()
	fake.ExchangeStub = nil
	if fake.exchangeReturnsOnCall == nil {
		fake.exchangeReturnsOnCall = make(map[int]struct {
			result1 *oauth2.Token
			result2 error
		})
	}
	fake.exchangeReturnsOnCall[i] = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) FollowChannel(arg1 context.Context, arg2 int64, arg3 *oauth2.Token, arg4 int64) (*oauth2.Token, error) {
	fake.followChannelMutex.Lock()
	ret, specificReturn := fake.followChannelReturnsOnCall[len(fake.followChannelArgsForCall)]
	fake.followChannelArgsForCall = append(fake.followChannelArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 int64
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("FollowChannel", []interface{}{arg1, arg2, arg3, arg4})
	fake.followChannelMutex.Unlock()
	if fake.FollowChannelStub != nil {
		return fake.FollowChannelStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.followChannelReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) FollowChannelCallCount() int {
	fake.followChannelMutex.RLock()
	defer fake.followChannelMutex.RUnlock()
	return len(fake.followChannelArgsForCall)
}

func (fake *FakeAPI) FollowChannelCalls(stub func(context.Context, int64, *oauth2.Token, int64) (*oauth2.Token, error)) {
	fake.followChannelMutex.Lock()
	defer fake.followChannelMutex.Unlock()
	fake.FollowChannelStub = stub
}

func (fake *FakeAPI) FollowChannelArgsForCall(i int) (context.Context, int64, *oauth2.Token, int64) {
	fake.followChannelMutex.RLock()
	defer fake.followChannelMutex.RUnlock()
	argsForCall := fake.followChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeAPI) FollowChannelReturns(result1 *oauth2.Token, result2 error) {
	fake.followChannelMutex.Lock()
	defer fake.followChannelMutex.Unlock()
	fake.FollowChannelStub = nil
	fake.followChannelReturns = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) FollowChannelReturnsOnCall(i int, result1 *oauth2.Token, result2 error) {
	fake.followChannelMutex.Lock()
	defer fake.followChannelMutex.Unlock()
	fake.FollowChannelStub = nil
	if fake.followChannelReturnsOnCall == nil {
		fake.followChannelReturnsOnCall = make(map[int]struct {
			result1 *oauth2.Token
			result2 error
		})
	}
	fake.followChannelReturnsOnCall[i] = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetChannelByID(arg1 context.Context, arg2 int64) (*twitch.Channel, error) {
	fake.getChannelByIDMutex.Lock()
	ret, specificReturn := fake.getChannelByIDReturnsOnCall[len(fake.getChannelByIDArgsForCall)]
	fake.getChannelByIDArgsForCall = append(fake.getChannelByIDArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	fake.recordInvocation("GetChannelByID", []interface{}{arg1, arg2})
	fake.getChannelByIDMutex.Unlock()
	if fake.GetChannelByIDStub != nil {
		return fake.GetChannelByIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChannelByIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetChannelByIDCallCount() int {
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	return len(fake.getChannelByIDArgsForCall)
}

func (fake *FakeAPI) GetChannelByIDCalls(stub func(context.Context, int64) (*twitch.Channel, error)) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = stub
}

func (fake *FakeAPI) GetChannelByIDArgsForCall(i int) (context.Context, int64) {
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	argsForCall := fake.getChannelByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetChannelByIDReturns(result1 *twitch.Channel, result2 error) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = nil
	fake.getChannelByIDReturns = struct {
		result1 *twitch.Channel
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetChannelByIDReturnsOnCall(i int, result1 *twitch.Channel, result2 error) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = nil
	if fake.getChannelByIDReturnsOnCall == nil {
		fake.getChannelByIDReturnsOnCall = make(map[int]struct {
			result1 *twitch.Channel
			result2 error
		})
	}
	fake.getChannelByIDReturnsOnCall[i] = struct {
		result1 *twitch.Channel
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetChatters(arg1 context.Context, arg2 string) (*twitch.Chatters, error) {
	fake.getChattersMutex.Lock()
	ret, specificReturn := fake.getChattersReturnsOnCall[len(fake.getChattersArgsForCall)]
	fake.getChattersArgsForCall = append(fake.getChattersArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetChatters", []interface{}{arg1, arg2})
	fake.getChattersMutex.Unlock()
	if fake.GetChattersStub != nil {
		return fake.GetChattersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChattersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetChattersCallCount() int {
	fake.getChattersMutex.RLock()
	defer fake.getChattersMutex.RUnlock()
	return len(fake.getChattersArgsForCall)
}

func (fake *FakeAPI) GetChattersCalls(stub func(context.Context, string) (*twitch.Chatters, error)) {
	fake.getChattersMutex.Lock()
	defer fake.getChattersMutex.Unlock()
	fake.GetChattersStub = stub
}

func (fake *FakeAPI) GetChattersArgsForCall(i int) (context.Context, string) {
	fake.getChattersMutex.RLock()
	defer fake.getChattersMutex.RUnlock()
	argsForCall := fake.getChattersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetChattersReturns(result1 *twitch.Chatters, result2 error) {
	fake.getChattersMutex.Lock()
	defer fake.getChattersMutex.Unlock()
	fake.GetChattersStub = nil
	fake.getChattersReturns = struct {
		result1 *twitch.Chatters
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetChattersReturnsOnCall(i int, result1 *twitch.Chatters, result2 error) {
	fake.getChattersMutex.Lock()
	defer fake.getChattersMutex.Unlock()
	fake.GetChattersStub = nil
	if fake.getChattersReturnsOnCall == nil {
		fake.getChattersReturnsOnCall = make(map[int]struct {
			result1 *twitch.Chatters
			result2 error
		})
	}
	fake.getChattersReturnsOnCall[i] = struct {
		result1 *twitch.Chatters
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetCurrentStream(arg1 context.Context, arg2 int64) (*twitch.Stream, error) {
	fake.getCurrentStreamMutex.Lock()
	ret, specificReturn := fake.getCurrentStreamReturnsOnCall[len(fake.getCurrentStreamArgsForCall)]
	fake.getCurrentStreamArgsForCall = append(fake.getCurrentStreamArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	fake.recordInvocation("GetCurrentStream", []interface{}{arg1, arg2})
	fake.getCurrentStreamMutex.Unlock()
	if fake.GetCurrentStreamStub != nil {
		return fake.GetCurrentStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getCurrentStreamReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetCurrentStreamCallCount() int {
	fake.getCurrentStreamMutex.RLock()
	defer fake.getCurrentStreamMutex.RUnlock()
	return len(fake.getCurrentStreamArgsForCall)
}

func (fake *FakeAPI) GetCurrentStreamCalls(stub func(context.Context, int64) (*twitch.Stream, error)) {
	fake.getCurrentStreamMutex.Lock()
	defer fake.getCurrentStreamMutex.Unlock()
	fake.GetCurrentStreamStub = stub
}

func (fake *FakeAPI) GetCurrentStreamArgsForCall(i int) (context.Context, int64) {
	fake.getCurrentStreamMutex.RLock()
	defer fake.getCurrentStreamMutex.RUnlock()
	argsForCall := fake.getCurrentStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetCurrentStreamReturns(result1 *twitch.Stream, result2 error) {
	fake.getCurrentStreamMutex.Lock()
	defer fake.getCurrentStreamMutex.Unlock()
	fake.GetCurrentStreamStub = nil
	fake.getCurrentStreamReturns = struct {
		result1 *twitch.Stream
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetCurrentStreamReturnsOnCall(i int, result1 *twitch.Stream, result2 error) {
	fake.getCurrentStreamMutex.Lock()
	defer fake.getCurrentStreamMutex.Unlock()
	fake.GetCurrentStreamStub = nil
	if fake.getCurrentStreamReturnsOnCall == nil {
		fake.getCurrentStreamReturnsOnCall = make(map[int]struct {
			result1 *twitch.Stream
			result2 error
		})
	}
	fake.getCurrentStreamReturnsOnCall[i] = struct {
		result1 *twitch.Stream
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetUserForToken(arg1 context.Context, arg2 *oauth2.Token) (*twitch.User, *oauth2.Token, error) {
	fake.getUserForTokenMutex.Lock()
	ret, specificReturn := fake.getUserForTokenReturnsOnCall[len(fake.getUserForTokenArgsForCall)]
	fake.getUserForTokenArgsForCall = append(fake.getUserForTokenArgsForCall, struct {
		arg1 context.Context
		arg2 *oauth2.Token
	}{arg1, arg2})
	fake.recordInvocation("GetUserForToken", []interface{}{arg1, arg2})
	fake.getUserForTokenMutex.Unlock()
	if fake.GetUserForTokenStub != nil {
		return fake.GetUserForTokenStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getUserForTokenReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAPI) GetUserForTokenCallCount() int {
	fake.getUserForTokenMutex.RLock()
	defer fake.getUserForTokenMutex.RUnlock()
	return len(fake.getUserForTokenArgsForCall)
}

func (fake *FakeAPI) GetUserForTokenCalls(stub func(context.Context, *oauth2.Token) (*twitch.User, *oauth2.Token, error)) {
	fake.getUserForTokenMutex.Lock()
	defer fake.getUserForTokenMutex.Unlock()
	fake.GetUserForTokenStub = stub
}

func (fake *FakeAPI) GetUserForTokenArgsForCall(i int) (context.Context, *oauth2.Token) {
	fake.getUserForTokenMutex.RLock()
	defer fake.getUserForTokenMutex.RUnlock()
	argsForCall := fake.getUserForTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetUserForTokenReturns(result1 *twitch.User, result2 *oauth2.Token, result3 error) {
	fake.getUserForTokenMutex.Lock()
	defer fake.getUserForTokenMutex.Unlock()
	fake.GetUserForTokenStub = nil
	fake.getUserForTokenReturns = struct {
		result1 *twitch.User
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) GetUserForTokenReturnsOnCall(i int, result1 *twitch.User, result2 *oauth2.Token, result3 error) {
	fake.getUserForTokenMutex.Lock()
	defer fake.getUserForTokenMutex.Unlock()
	fake.GetUserForTokenStub = nil
	if fake.getUserForTokenReturnsOnCall == nil {
		fake.getUserForTokenReturnsOnCall = make(map[int]struct {
			result1 *twitch.User
			result2 *oauth2.Token
			result3 error
		})
	}
	fake.getUserForTokenReturnsOnCall[i] = struct {
		result1 *twitch.User
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) GetUserForUsername(arg1 context.Context, arg2 string) (*twitch.User, error) {
	fake.getUserForUsernameMutex.Lock()
	ret, specificReturn := fake.getUserForUsernameReturnsOnCall[len(fake.getUserForUsernameArgsForCall)]
	fake.getUserForUsernameArgsForCall = append(fake.getUserForUsernameArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetUserForUsername", []interface{}{arg1, arg2})
	fake.getUserForUsernameMutex.Unlock()
	if fake.GetUserForUsernameStub != nil {
		return fake.GetUserForUsernameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserForUsernameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetUserForUsernameCallCount() int {
	fake.getUserForUsernameMutex.RLock()
	defer fake.getUserForUsernameMutex.RUnlock()
	return len(fake.getUserForUsernameArgsForCall)
}

func (fake *FakeAPI) GetUserForUsernameCalls(stub func(context.Context, string) (*twitch.User, error)) {
	fake.getUserForUsernameMutex.Lock()
	defer fake.getUserForUsernameMutex.Unlock()
	fake.GetUserForUsernameStub = stub
}

func (fake *FakeAPI) GetUserForUsernameArgsForCall(i int) (context.Context, string) {
	fake.getUserForUsernameMutex.RLock()
	defer fake.getUserForUsernameMutex.RUnlock()
	argsForCall := fake.getUserForUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetUserForUsernameReturns(result1 *twitch.User, result2 error) {
	fake.getUserForUsernameMutex.Lock()
	defer fake.getUserForUsernameMutex.Unlock()
	fake.GetUserForUsernameStub = nil
	fake.getUserForUsernameReturns = struct {
		result1 *twitch.User
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetUserForUsernameReturnsOnCall(i int, result1 *twitch.User, result2 error) {
	fake.getUserForUsernameMutex.Lock()
	defer fake.getUserForUsernameMutex.Unlock()
	fake.GetUserForUsernameStub = nil
	if fake.getUserForUsernameReturnsOnCall == nil {
		fake.getUserForUsernameReturnsOnCall = make(map[int]struct {
			result1 *twitch.User
			result2 error
		})
	}
	fake.getUserForUsernameReturnsOnCall[i] = struct {
		result1 *twitch.User
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) SetChannelGame(arg1 context.Context, arg2 int64, arg3 *oauth2.Token, arg4 string) (string, *oauth2.Token, error) {
	fake.setChannelGameMutex.Lock()
	ret, specificReturn := fake.setChannelGameReturnsOnCall[len(fake.setChannelGameArgsForCall)]
	fake.setChannelGameArgsForCall = append(fake.setChannelGameArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetChannelGame", []interface{}{arg1, arg2, arg3, arg4})
	fake.setChannelGameMutex.Unlock()
	if fake.SetChannelGameStub != nil {
		return fake.SetChannelGameStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.setChannelGameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAPI) SetChannelGameCallCount() int {
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	return len(fake.setChannelGameArgsForCall)
}

func (fake *FakeAPI) SetChannelGameCalls(stub func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error)) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = stub
}

func (fake *FakeAPI) SetChannelGameArgsForCall(i int) (context.Context, int64, *oauth2.Token, string) {
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	argsForCall := fake.setChannelGameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeAPI) SetChannelGameReturns(result1 string, result2 *oauth2.Token, result3 error) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = nil
	fake.setChannelGameReturns = struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) SetChannelGameReturnsOnCall(i int, result1 string, result2 *oauth2.Token, result3 error) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = nil
	if fake.setChannelGameReturnsOnCall == nil {
		fake.setChannelGameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 *oauth2.Token
			result3 error
		})
	}
	fake.setChannelGameReturnsOnCall[i] = struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) SetChannelStatus(arg1 context.Context, arg2 int64, arg3 *oauth2.Token, arg4 string) (string, *oauth2.Token, error) {
	fake.setChannelStatusMutex.Lock()
	ret, specificReturn := fake.setChannelStatusReturnsOnCall[len(fake.setChannelStatusArgsForCall)]
	fake.setChannelStatusArgsForCall = append(fake.setChannelStatusArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetChannelStatus", []interface{}{arg1, arg2, arg3, arg4})
	fake.setChannelStatusMutex.Unlock()
	if fake.SetChannelStatusStub != nil {
		return fake.SetChannelStatusStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.setChannelStatusReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAPI) SetChannelStatusCallCount() int {
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
	return len(fake.setChannelStatusArgsForCall)
}

func (fake *FakeAPI) SetChannelStatusCalls(stub func(context.Context, int64, *oauth2.Token, string) (string, *oauth2.Token, error)) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = stub
}

func (fake *FakeAPI) SetChannelStatusArgsForCall(i int) (context.Context, int64, *oauth2.Token, string) {
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
	argsForCall := fake.setChannelStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeAPI) SetChannelStatusReturns(result1 string, result2 *oauth2.Token, result3 error) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = nil
	fake.setChannelStatusReturns = struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) SetChannelStatusReturnsOnCall(i int, result1 string, result2 *oauth2.Token, result3 error) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = nil
	if fake.setChannelStatusReturnsOnCall == nil {
		fake.setChannelStatusReturnsOnCall = make(map[int]struct {
			result1 string
			result2 *oauth2.Token
			result3 error
		})
	}
	fake.setChannelStatusReturnsOnCall[i] = struct {
		result1 string
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authCodeURLMutex.RLock()
	defer fake.authCodeURLMutex.RUnlock()
	fake.exchangeMutex.RLock()
	defer fake.exchangeMutex.RUnlock()
	fake.followChannelMutex.RLock()
	defer fake.followChannelMutex.RUnlock()
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	fake.getChattersMutex.RLock()
	defer fake.getChattersMutex.RUnlock()
	fake.getCurrentStreamMutex.RLock()
	defer fake.getCurrentStreamMutex.RUnlock()
	fake.getUserForTokenMutex.RLock()
	defer fake.getUserForTokenMutex.RUnlock()
	fake.getUserForUsernameMutex.RLock()
	defer fake.getUserForUsernameMutex.RUnlock()
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
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

var _ twitch.API = new(FakeAPI)
