// Code generated by mockery v1.0.0. DO NOT EDIT.

package syncasyncmocks

import (
	context "context"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"
	core "github.com/hyperledger/firefly/pkg/core"

	mock "github.com/stretchr/testify/mock"

	syncasync "github.com/hyperledger/firefly/internal/syncasync"

	system "github.com/hyperledger/firefly/internal/events/system"
)

// Bridge is an autogenerated mock type for the Bridge type
type Bridge struct {
	mock.Mock
}

// Init provides a mock function with given fields: sysevents
func (_m *Bridge) Init(sysevents system.EventInterface) {
	_m.Called(sysevents)
}

// WaitForIdentity provides a mock function with given fields: ctx, id, send
func (_m *Bridge) WaitForIdentity(ctx context.Context, id *fftypes.UUID, send syncasync.RequestSender) (*core.Identity, error) {
	ret := _m.Called(ctx, id, send)

	var r0 *core.Identity
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) *core.Identity); ok {
		r0 = rf(ctx, id, send)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) error); ok {
		r1 = rf(ctx, id, send)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForInvokeOperation provides a mock function with given fields: ctx, id, send
func (_m *Bridge) WaitForInvokeOperation(ctx context.Context, id *fftypes.UUID, send syncasync.RequestSender) (*core.Operation, error) {
	ret := _m.Called(ctx, id, send)

	var r0 *core.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) *core.Operation); ok {
		r0 = rf(ctx, id, send)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) error); ok {
		r1 = rf(ctx, id, send)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForMessage provides a mock function with given fields: ctx, id, send
func (_m *Bridge) WaitForMessage(ctx context.Context, id *fftypes.UUID, send syncasync.RequestSender) (*core.Message, error) {
	ret := _m.Called(ctx, id, send)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) *core.Message); ok {
		r0 = rf(ctx, id, send)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) error); ok {
		r1 = rf(ctx, id, send)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForReply provides a mock function with given fields: ctx, id, send
func (_m *Bridge) WaitForReply(ctx context.Context, id *fftypes.UUID, send syncasync.RequestSender) (*core.MessageInOut, error) {
	ret := _m.Called(ctx, id, send)

	var r0 *core.MessageInOut
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) *core.MessageInOut); ok {
		r0 = rf(ctx, id, send)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.MessageInOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) error); ok {
		r1 = rf(ctx, id, send)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForTokenApproval provides a mock function with given fields: ctx, id, send
func (_m *Bridge) WaitForTokenApproval(ctx context.Context, id *fftypes.UUID, send syncasync.RequestSender) (*core.TokenApproval, error) {
	ret := _m.Called(ctx, id, send)

	var r0 *core.TokenApproval
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) *core.TokenApproval); ok {
		r0 = rf(ctx, id, send)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TokenApproval)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) error); ok {
		r1 = rf(ctx, id, send)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForTokenPool provides a mock function with given fields: ctx, id, send
func (_m *Bridge) WaitForTokenPool(ctx context.Context, id *fftypes.UUID, send syncasync.RequestSender) (*core.TokenPool, error) {
	ret := _m.Called(ctx, id, send)

	var r0 *core.TokenPool
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) *core.TokenPool); ok {
		r0 = rf(ctx, id, send)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TokenPool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) error); ok {
		r1 = rf(ctx, id, send)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForTokenTransfer provides a mock function with given fields: ctx, id, send
func (_m *Bridge) WaitForTokenTransfer(ctx context.Context, id *fftypes.UUID, send syncasync.RequestSender) (*core.TokenTransfer, error) {
	ret := _m.Called(ctx, id, send)

	var r0 *core.TokenTransfer
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) *core.TokenTransfer); ok {
		r0 = rf(ctx, id, send)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TokenTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, syncasync.RequestSender) error); ok {
		r1 = rf(ctx, id, send)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
