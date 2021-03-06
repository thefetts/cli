// This file was generated by counterfeiter
package netfakes

import (
	_ "crypto/sha512"
	"net/http"
	"sync"

	"code.cloudfoundry.org/cli/cf/net"
)

type FakeHTTPClientInterface struct {
	DumpRequestStub        func(*http.Request)
	dumpRequestMutex       sync.RWMutex
	dumpRequestArgsForCall []struct {
		arg1 *http.Request
	}
	DumpResponseStub        func(*http.Response)
	dumpResponseMutex       sync.RWMutex
	dumpResponseArgsForCall []struct {
		arg1 *http.Response
	}
	DoStub        func(*http.Request) (*http.Response, error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		arg1 *http.Request
	}
	doReturns struct {
		result1 *http.Response
		result2 error
	}
	ExecuteCheckRedirectStub        func(req *http.Request, via []*http.Request) error
	executeCheckRedirectMutex       sync.RWMutex
	executeCheckRedirectArgsForCall []struct {
		req *http.Request
		via []*http.Request
	}
	executeCheckRedirectReturns struct {
		result1 error
	}
}

func (fake *FakeHTTPClientInterface) DumpRequest(arg1 *http.Request) {
	fake.dumpRequestMutex.Lock()
	fake.dumpRequestArgsForCall = append(fake.dumpRequestArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.dumpRequestMutex.Unlock()
	if fake.DumpRequestStub != nil {
		fake.DumpRequestStub(arg1)
	}
}

func (fake *FakeHTTPClientInterface) DumpRequestCallCount() int {
	fake.dumpRequestMutex.RLock()
	defer fake.dumpRequestMutex.RUnlock()
	return len(fake.dumpRequestArgsForCall)
}

func (fake *FakeHTTPClientInterface) DumpRequestArgsForCall(i int) *http.Request {
	fake.dumpRequestMutex.RLock()
	defer fake.dumpRequestMutex.RUnlock()
	return fake.dumpRequestArgsForCall[i].arg1
}

func (fake *FakeHTTPClientInterface) DumpResponse(arg1 *http.Response) {
	fake.dumpResponseMutex.Lock()
	fake.dumpResponseArgsForCall = append(fake.dumpResponseArgsForCall, struct {
		arg1 *http.Response
	}{arg1})
	fake.dumpResponseMutex.Unlock()
	if fake.DumpResponseStub != nil {
		fake.DumpResponseStub(arg1)
	}
}

func (fake *FakeHTTPClientInterface) DumpResponseCallCount() int {
	fake.dumpResponseMutex.RLock()
	defer fake.dumpResponseMutex.RUnlock()
	return len(fake.dumpResponseArgsForCall)
}

func (fake *FakeHTTPClientInterface) DumpResponseArgsForCall(i int) *http.Response {
	fake.dumpResponseMutex.RLock()
	defer fake.dumpResponseMutex.RUnlock()
	return fake.dumpResponseArgsForCall[i].arg1
}

func (fake *FakeHTTPClientInterface) Do(arg1 *http.Request) (*http.Response, error) {
	fake.doMutex.Lock()
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(arg1)
	} else {
		return fake.doReturns.result1, fake.doReturns.result2
	}
}

func (fake *FakeHTTPClientInterface) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeHTTPClientInterface) DoArgsForCall(i int) *http.Request {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return fake.doArgsForCall[i].arg1
}

func (fake *FakeHTTPClientInterface) DoReturns(result1 *http.Response, result2 error) {
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirect(req *http.Request, via []*http.Request) error {
	fake.executeCheckRedirectMutex.Lock()
	fake.executeCheckRedirectArgsForCall = append(fake.executeCheckRedirectArgsForCall, struct {
		req *http.Request
		via []*http.Request
	}{req, via})
	fake.executeCheckRedirectMutex.Unlock()
	if fake.ExecuteCheckRedirectStub != nil {
		return fake.ExecuteCheckRedirectStub(req, via)
	} else {
		return fake.executeCheckRedirectReturns.result1
	}
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirectCallCount() int {
	fake.executeCheckRedirectMutex.RLock()
	defer fake.executeCheckRedirectMutex.RUnlock()
	return len(fake.executeCheckRedirectArgsForCall)
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirectArgsForCall(i int) (*http.Request, []*http.Request) {
	fake.executeCheckRedirectMutex.RLock()
	defer fake.executeCheckRedirectMutex.RUnlock()
	return fake.executeCheckRedirectArgsForCall[i].req, fake.executeCheckRedirectArgsForCall[i].via
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirectReturns(result1 error) {
	fake.ExecuteCheckRedirectStub = nil
	fake.executeCheckRedirectReturns = struct {
		result1 error
	}{result1}
}

var _ net.HTTPClientInterface = new(FakeHTTPClientInterface)
