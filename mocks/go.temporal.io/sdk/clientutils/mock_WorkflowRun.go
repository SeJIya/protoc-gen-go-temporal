// Code generated by mockery v2.38.0. DO NOT EDIT.

package clientutils

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	internal "go.temporal.io/sdk/client"
)

// MockWorkflowRun is an autogenerated mock type for the WorkflowRun type
type MockWorkflowRun struct {
	mock.Mock
}

type MockWorkflowRun_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWorkflowRun) EXPECT() *MockWorkflowRun_Expecter {
	return &MockWorkflowRun_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, valuePtr
func (_m *MockWorkflowRun) Get(ctx context.Context, valuePtr interface{}) error {
	ret := _m.Called(ctx, valuePtr)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, valuePtr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkflowRun_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockWorkflowRun_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - valuePtr interface{}
func (_e *MockWorkflowRun_Expecter) Get(ctx interface{}, valuePtr interface{}) *MockWorkflowRun_Get_Call {
	return &MockWorkflowRun_Get_Call{Call: _e.mock.On("Get", ctx, valuePtr)}
}

func (_c *MockWorkflowRun_Get_Call) Run(run func(ctx context.Context, valuePtr interface{})) *MockWorkflowRun_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *MockWorkflowRun_Get_Call) Return(_a0 error) *MockWorkflowRun_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkflowRun_Get_Call) RunAndReturn(run func(context.Context, interface{}) error) *MockWorkflowRun_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetID provides a mock function with given fields:
func (_m *MockWorkflowRun) GetID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockWorkflowRun_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetID'
type MockWorkflowRun_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
func (_e *MockWorkflowRun_Expecter) GetID() *MockWorkflowRun_GetID_Call {
	return &MockWorkflowRun_GetID_Call{Call: _e.mock.On("GetID")}
}

func (_c *MockWorkflowRun_GetID_Call) Run(run func()) *MockWorkflowRun_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorkflowRun_GetID_Call) Return(_a0 string) *MockWorkflowRun_GetID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkflowRun_GetID_Call) RunAndReturn(run func() string) *MockWorkflowRun_GetID_Call {
	_c.Call.Return(run)
	return _c
}

// GetRunID provides a mock function with given fields:
func (_m *MockWorkflowRun) GetRunID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRunID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockWorkflowRun_GetRunID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRunID'
type MockWorkflowRun_GetRunID_Call struct {
	*mock.Call
}

// GetRunID is a helper method to define mock.On call
func (_e *MockWorkflowRun_Expecter) GetRunID() *MockWorkflowRun_GetRunID_Call {
	return &MockWorkflowRun_GetRunID_Call{Call: _e.mock.On("GetRunID")}
}

func (_c *MockWorkflowRun_GetRunID_Call) Run(run func()) *MockWorkflowRun_GetRunID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWorkflowRun_GetRunID_Call) Return(_a0 string) *MockWorkflowRun_GetRunID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkflowRun_GetRunID_Call) RunAndReturn(run func() string) *MockWorkflowRun_GetRunID_Call {
	_c.Call.Return(run)
	return _c
}

// GetWithOptions provides a mock function with given fields: ctx, valuePtr, options
func (_m *MockWorkflowRun) GetWithOptions(ctx context.Context, valuePtr interface{}, options internal.WorkflowRunGetOptions) error {
	ret := _m.Called(ctx, valuePtr, options)

	if len(ret) == 0 {
		panic("no return value specified for GetWithOptions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, internal.WorkflowRunGetOptions) error); ok {
		r0 = rf(ctx, valuePtr, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWorkflowRun_GetWithOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithOptions'
type MockWorkflowRun_GetWithOptions_Call struct {
	*mock.Call
}

// GetWithOptions is a helper method to define mock.On call
//   - ctx context.Context
//   - valuePtr interface{}
//   - options internal.WorkflowRunGetOptions
func (_e *MockWorkflowRun_Expecter) GetWithOptions(ctx interface{}, valuePtr interface{}, options interface{}) *MockWorkflowRun_GetWithOptions_Call {
	return &MockWorkflowRun_GetWithOptions_Call{Call: _e.mock.On("GetWithOptions", ctx, valuePtr, options)}
}

func (_c *MockWorkflowRun_GetWithOptions_Call) Run(run func(ctx context.Context, valuePtr interface{}, options internal.WorkflowRunGetOptions)) *MockWorkflowRun_GetWithOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(internal.WorkflowRunGetOptions))
	})
	return _c
}

func (_c *MockWorkflowRun_GetWithOptions_Call) Return(_a0 error) *MockWorkflowRun_GetWithOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWorkflowRun_GetWithOptions_Call) RunAndReturn(run func(context.Context, interface{}, internal.WorkflowRunGetOptions) error) *MockWorkflowRun_GetWithOptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWorkflowRun creates a new instance of MockWorkflowRun. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWorkflowRun(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkflowRun {
	mock := &MockWorkflowRun{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
