// Code generated by mockery v2.42.2. DO NOT EDIT.

package rpclibmocks

import (
	context "context"

	rpclib "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	mock "github.com/stretchr/testify/mock"
)

// EvmBatchCaller is an autogenerated mock type for the EvmBatchCaller type
type EvmBatchCaller struct {
	mock.Mock
}

// BatchCall provides a mock function with given fields: ctx, blockNumber, calls
func (_m *EvmBatchCaller) BatchCall(ctx context.Context, blockNumber uint64, calls []rpclib.EvmCall) ([]rpclib.DataAndErr, error) {
	ret := _m.Called(ctx, blockNumber, calls)

	if len(ret) == 0 {
		panic("no return value specified for BatchCall")
	}

	var r0 []rpclib.DataAndErr
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []rpclib.EvmCall) ([]rpclib.DataAndErr, error)); ok {
		return rf(ctx, blockNumber, calls)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []rpclib.EvmCall) []rpclib.DataAndErr); ok {
		r0 = rf(ctx, blockNumber, calls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rpclib.DataAndErr)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, []rpclib.EvmCall) error); ok {
		r1 = rf(ctx, blockNumber, calls)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEvmBatchCaller creates a new instance of EvmBatchCaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvmBatchCaller(t interface {
	mock.TestingT
	Cleanup(func())
}) *EvmBatchCaller {
	mock := &EvmBatchCaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}