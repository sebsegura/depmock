// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	context "context"
	client "sebsegura/sample-lambda/pkg/client"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: ctx, r
func (_m *Client) Authenticate(ctx context.Context, r *client.AuthRequest) (*client.AuthResponse, error) {
	ret := _m.Called(ctx, r)

	var r0 *client.AuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.AuthRequest) (*client.AuthResponse, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *client.AuthRequest) *client.AuthResponse); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.AuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *client.AuthRequest) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOperations provides a mock function with given fields: ctx, r
func (_m *Client) GetOperations(ctx context.Context, r *client.GetOperationsRequest) (*client.GetOperationsResponse, error) {
	ret := _m.Called(ctx, r)

	var r0 *client.GetOperationsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.GetOperationsRequest) (*client.GetOperationsResponse, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *client.GetOperationsRequest) *client.GetOperationsResponse); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.GetOperationsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *client.GetOperationsRequest) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
