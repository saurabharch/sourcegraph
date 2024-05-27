// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package productsubscriptiontest

import (
	"context"
	"sync"

	productsubscription "github.com/sourcegraph/sourcegraph/cmd/cody-gateway/internal/actor/productsubscription"
	v1 "github.com/sourcegraph/sourcegraph/lib/enterpriseportal/codyaccess/v1"
	v11 "github.com/sourcegraph/sourcegraph/lib/enterpriseportal/subscriptions/v1"
	grpc "google.golang.org/grpc"
)

// MockEnterprisePortalClient is a mock implementation of the
// EnterprisePortalClient interface (from the package
// github.com/sourcegraph/sourcegraph/cmd/cody-gateway/internal/actor/productsubscription)
// used for unit testing.
type MockEnterprisePortalClient struct {
	// GetCodyGatewayAccessFunc is an instance of a mock function object
	// controlling the behavior of the method GetCodyGatewayAccess.
	GetCodyGatewayAccessFunc *EnterprisePortalClientGetCodyGatewayAccessFunc
	// ListCodyGatewayAccessesFunc is an instance of a mock function object
	// controlling the behavior of the method ListCodyGatewayAccesses.
	ListCodyGatewayAccessesFunc *EnterprisePortalClientListCodyGatewayAccessesFunc
	// ListEnterpriseSubscriptionLicensesFunc is an instance of a mock
	// function object controlling the behavior of the method
	// ListEnterpriseSubscriptionLicenses.
	ListEnterpriseSubscriptionLicensesFunc *EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc
}

// NewMockEnterprisePortalClient creates a new mock of the
// EnterprisePortalClient interface. All methods return zero values for all
// results, unless overwritten.
func NewMockEnterprisePortalClient() *MockEnterprisePortalClient {
	return &MockEnterprisePortalClient{
		GetCodyGatewayAccessFunc: &EnterprisePortalClientGetCodyGatewayAccessFunc{
			defaultHook: func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (r0 *v1.GetCodyGatewayAccessResponse, r1 error) {
				return
			},
		},
		ListCodyGatewayAccessesFunc: &EnterprisePortalClientListCodyGatewayAccessesFunc{
			defaultHook: func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (r0 *v1.ListCodyGatewayAccessesResponse, r1 error) {
				return
			},
		},
		ListEnterpriseSubscriptionLicensesFunc: &EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc{
			defaultHook: func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (r0 *v11.ListEnterpriseSubscriptionLicensesResponse, r1 error) {
				return
			},
		},
	}
}

// NewStrictMockEnterprisePortalClient creates a new mock of the
// EnterprisePortalClient interface. All methods panic on invocation, unless
// overwritten.
func NewStrictMockEnterprisePortalClient() *MockEnterprisePortalClient {
	return &MockEnterprisePortalClient{
		GetCodyGatewayAccessFunc: &EnterprisePortalClientGetCodyGatewayAccessFunc{
			defaultHook: func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error) {
				panic("unexpected invocation of MockEnterprisePortalClient.GetCodyGatewayAccess")
			},
		},
		ListCodyGatewayAccessesFunc: &EnterprisePortalClientListCodyGatewayAccessesFunc{
			defaultHook: func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error) {
				panic("unexpected invocation of MockEnterprisePortalClient.ListCodyGatewayAccesses")
			},
		},
		ListEnterpriseSubscriptionLicensesFunc: &EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc{
			defaultHook: func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error) {
				panic("unexpected invocation of MockEnterprisePortalClient.ListEnterpriseSubscriptionLicenses")
			},
		},
	}
}

// NewMockEnterprisePortalClientFrom creates a new mock of the
// MockEnterprisePortalClient interface. All methods delegate to the given
// implementation, unless overwritten.
func NewMockEnterprisePortalClientFrom(i productsubscription.EnterprisePortalClient) *MockEnterprisePortalClient {
	return &MockEnterprisePortalClient{
		GetCodyGatewayAccessFunc: &EnterprisePortalClientGetCodyGatewayAccessFunc{
			defaultHook: i.GetCodyGatewayAccess,
		},
		ListCodyGatewayAccessesFunc: &EnterprisePortalClientListCodyGatewayAccessesFunc{
			defaultHook: i.ListCodyGatewayAccesses,
		},
		ListEnterpriseSubscriptionLicensesFunc: &EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc{
			defaultHook: i.ListEnterpriseSubscriptionLicenses,
		},
	}
}

// EnterprisePortalClientGetCodyGatewayAccessFunc describes the behavior
// when the GetCodyGatewayAccess method of the parent
// MockEnterprisePortalClient instance is invoked.
type EnterprisePortalClientGetCodyGatewayAccessFunc struct {
	defaultHook func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error)
	hooks       []func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error)
	history     []EnterprisePortalClientGetCodyGatewayAccessFuncCall
	mutex       sync.Mutex
}

// GetCodyGatewayAccess delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockEnterprisePortalClient) GetCodyGatewayAccess(v0 context.Context, v1 *v1.GetCodyGatewayAccessRequest, v2 ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error) {
	r0, r1 := m.GetCodyGatewayAccessFunc.nextHook()(v0, v1, v2...)
	m.GetCodyGatewayAccessFunc.appendCall(EnterprisePortalClientGetCodyGatewayAccessFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetCodyGatewayAccess
// method of the parent MockEnterprisePortalClient instance is invoked and
// the hook queue is empty.
func (f *EnterprisePortalClientGetCodyGatewayAccessFunc) SetDefaultHook(hook func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetCodyGatewayAccess method of the parent MockEnterprisePortalClient
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *EnterprisePortalClientGetCodyGatewayAccessFunc) PushHook(hook func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *EnterprisePortalClientGetCodyGatewayAccessFunc) SetDefaultReturn(r0 *v1.GetCodyGatewayAccessResponse, r1 error) {
	f.SetDefaultHook(func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *EnterprisePortalClientGetCodyGatewayAccessFunc) PushReturn(r0 *v1.GetCodyGatewayAccessResponse, r1 error) {
	f.PushHook(func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error) {
		return r0, r1
	})
}

func (f *EnterprisePortalClientGetCodyGatewayAccessFunc) nextHook() func(context.Context, *v1.GetCodyGatewayAccessRequest, ...grpc.CallOption) (*v1.GetCodyGatewayAccessResponse, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *EnterprisePortalClientGetCodyGatewayAccessFunc) appendCall(r0 EnterprisePortalClientGetCodyGatewayAccessFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// EnterprisePortalClientGetCodyGatewayAccessFuncCall objects describing the
// invocations of this function.
func (f *EnterprisePortalClientGetCodyGatewayAccessFunc) History() []EnterprisePortalClientGetCodyGatewayAccessFuncCall {
	f.mutex.Lock()
	history := make([]EnterprisePortalClientGetCodyGatewayAccessFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// EnterprisePortalClientGetCodyGatewayAccessFuncCall is an object that
// describes an invocation of method GetCodyGatewayAccess on an instance of
// MockEnterprisePortalClient.
type EnterprisePortalClientGetCodyGatewayAccessFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *v1.GetCodyGatewayAccessRequest
	// Arg2 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg2 []grpc.CallOption
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *v1.GetCodyGatewayAccessResponse
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c EnterprisePortalClientGetCodyGatewayAccessFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg2 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0, c.Arg1}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c EnterprisePortalClientGetCodyGatewayAccessFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// EnterprisePortalClientListCodyGatewayAccessesFunc describes the behavior
// when the ListCodyGatewayAccesses method of the parent
// MockEnterprisePortalClient instance is invoked.
type EnterprisePortalClientListCodyGatewayAccessesFunc struct {
	defaultHook func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error)
	hooks       []func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error)
	history     []EnterprisePortalClientListCodyGatewayAccessesFuncCall
	mutex       sync.Mutex
}

// ListCodyGatewayAccesses delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockEnterprisePortalClient) ListCodyGatewayAccesses(v0 context.Context, v1 *v1.ListCodyGatewayAccessesRequest, v2 ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error) {
	r0, r1 := m.ListCodyGatewayAccessesFunc.nextHook()(v0, v1, v2...)
	m.ListCodyGatewayAccessesFunc.appendCall(EnterprisePortalClientListCodyGatewayAccessesFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// ListCodyGatewayAccesses method of the parent MockEnterprisePortalClient
// instance is invoked and the hook queue is empty.
func (f *EnterprisePortalClientListCodyGatewayAccessesFunc) SetDefaultHook(hook func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ListCodyGatewayAccesses method of the parent MockEnterprisePortalClient
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *EnterprisePortalClientListCodyGatewayAccessesFunc) PushHook(hook func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *EnterprisePortalClientListCodyGatewayAccessesFunc) SetDefaultReturn(r0 *v1.ListCodyGatewayAccessesResponse, r1 error) {
	f.SetDefaultHook(func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *EnterprisePortalClientListCodyGatewayAccessesFunc) PushReturn(r0 *v1.ListCodyGatewayAccessesResponse, r1 error) {
	f.PushHook(func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error) {
		return r0, r1
	})
}

func (f *EnterprisePortalClientListCodyGatewayAccessesFunc) nextHook() func(context.Context, *v1.ListCodyGatewayAccessesRequest, ...grpc.CallOption) (*v1.ListCodyGatewayAccessesResponse, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *EnterprisePortalClientListCodyGatewayAccessesFunc) appendCall(r0 EnterprisePortalClientListCodyGatewayAccessesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// EnterprisePortalClientListCodyGatewayAccessesFuncCall objects describing
// the invocations of this function.
func (f *EnterprisePortalClientListCodyGatewayAccessesFunc) History() []EnterprisePortalClientListCodyGatewayAccessesFuncCall {
	f.mutex.Lock()
	history := make([]EnterprisePortalClientListCodyGatewayAccessesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// EnterprisePortalClientListCodyGatewayAccessesFuncCall is an object that
// describes an invocation of method ListCodyGatewayAccesses on an instance
// of MockEnterprisePortalClient.
type EnterprisePortalClientListCodyGatewayAccessesFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *v1.ListCodyGatewayAccessesRequest
	// Arg2 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg2 []grpc.CallOption
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *v1.ListCodyGatewayAccessesResponse
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c EnterprisePortalClientListCodyGatewayAccessesFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg2 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0, c.Arg1}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c EnterprisePortalClientListCodyGatewayAccessesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc describes
// the behavior when the ListEnterpriseSubscriptionLicenses method of the
// parent MockEnterprisePortalClient instance is invoked.
type EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc struct {
	defaultHook func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error)
	hooks       []func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error)
	history     []EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall
	mutex       sync.Mutex
}

// ListEnterpriseSubscriptionLicenses delegates to the next hook function in
// the queue and stores the parameter and result values of this invocation.
func (m *MockEnterprisePortalClient) ListEnterpriseSubscriptionLicenses(v0 context.Context, v1 *v11.ListEnterpriseSubscriptionLicensesRequest, v2 ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error) {
	r0, r1 := m.ListEnterpriseSubscriptionLicensesFunc.nextHook()(v0, v1, v2...)
	m.ListEnterpriseSubscriptionLicensesFunc.appendCall(EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// ListEnterpriseSubscriptionLicenses method of the parent
// MockEnterprisePortalClient instance is invoked and the hook queue is
// empty.
func (f *EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc) SetDefaultHook(hook func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ListEnterpriseSubscriptionLicenses method of the parent
// MockEnterprisePortalClient instance invokes the hook at the front of the
// queue and discards it. After the queue is empty, the default hook
// function is invoked for any future action.
func (f *EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc) PushHook(hook func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc) SetDefaultReturn(r0 *v11.ListEnterpriseSubscriptionLicensesResponse, r1 error) {
	f.SetDefaultHook(func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc) PushReturn(r0 *v11.ListEnterpriseSubscriptionLicensesResponse, r1 error) {
	f.PushHook(func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error) {
		return r0, r1
	})
}

func (f *EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc) nextHook() func(context.Context, *v11.ListEnterpriseSubscriptionLicensesRequest, ...grpc.CallOption) (*v11.ListEnterpriseSubscriptionLicensesResponse, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc) appendCall(r0 EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall objects
// describing the invocations of this function.
func (f *EnterprisePortalClientListEnterpriseSubscriptionLicensesFunc) History() []EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall {
	f.mutex.Lock()
	history := make([]EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall is an
// object that describes an invocation of method
// ListEnterpriseSubscriptionLicenses on an instance of
// MockEnterprisePortalClient.
type EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *v11.ListEnterpriseSubscriptionLicensesRequest
	// Arg2 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg2 []grpc.CallOption
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *v11.ListEnterpriseSubscriptionLicensesResponse
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg2 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0, c.Arg1}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c EnterprisePortalClientListEnterpriseSubscriptionLicensesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
