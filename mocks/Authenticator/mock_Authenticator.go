// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocksauth

import (
	dynamic "k8s.io/client-go/dynamic"
	kubernetes "k8s.io/client-go/kubernetes"

	mock "github.com/stretchr/testify/mock"
)

// MockAuthenticator is an autogenerated mock type for the Authenticator type
type MockAuthenticator struct {
	mock.Mock
}

type MockAuthenticator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthenticator) EXPECT() *MockAuthenticator_Expecter {
	return &MockAuthenticator_Expecter{mock: &_m.Mock}
}

// DynamicAPI provides a mock function with no fields
func (_m *MockAuthenticator) DynamicAPI() (dynamic.Interface, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DynamicAPI")
	}

	var r0 dynamic.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func() (dynamic.Interface, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() dynamic.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthenticator_DynamicAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DynamicAPI'
type MockAuthenticator_DynamicAPI_Call struct {
	*mock.Call
}

// DynamicAPI is a helper method to define mock.On call
func (_e *MockAuthenticator_Expecter) DynamicAPI() *MockAuthenticator_DynamicAPI_Call {
	return &MockAuthenticator_DynamicAPI_Call{Call: _e.mock.On("DynamicAPI")}
}

func (_c *MockAuthenticator_DynamicAPI_Call) Run(run func()) *MockAuthenticator_DynamicAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAuthenticator_DynamicAPI_Call) Return(_a0 dynamic.Interface, _a1 error) *MockAuthenticator_DynamicAPI_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthenticator_DynamicAPI_Call) RunAndReturn(run func() (dynamic.Interface, error)) *MockAuthenticator_DynamicAPI_Call {
	_c.Call.Return(run)
	return _c
}

// NativeAPI provides a mock function with no fields
func (_m *MockAuthenticator) NativeAPI() (kubernetes.Interface, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NativeAPI")
	}

	var r0 kubernetes.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func() (kubernetes.Interface, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() kubernetes.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kubernetes.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthenticator_NativeAPI_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NativeAPI'
type MockAuthenticator_NativeAPI_Call struct {
	*mock.Call
}

// NativeAPI is a helper method to define mock.On call
func (_e *MockAuthenticator_Expecter) NativeAPI() *MockAuthenticator_NativeAPI_Call {
	return &MockAuthenticator_NativeAPI_Call{Call: _e.mock.On("NativeAPI")}
}

func (_c *MockAuthenticator_NativeAPI_Call) Run(run func()) *MockAuthenticator_NativeAPI_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAuthenticator_NativeAPI_Call) Return(_a0 kubernetes.Interface, _a1 error) *MockAuthenticator_NativeAPI_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthenticator_NativeAPI_Call) RunAndReturn(run func() (kubernetes.Interface, error)) *MockAuthenticator_NativeAPI_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAuthenticator creates a new instance of MockAuthenticator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthenticator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthenticator {
	mock := &MockAuthenticator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
