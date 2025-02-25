// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	abcicli "github.com/line/ostracon/abci/client"
	types "github.com/line/ostracon/abci/types"
	mock "github.com/stretchr/testify/mock"
)

// AppConnMempool is an autogenerated mock type for the AppConnMempool type
type AppConnMempool struct {
	mock.Mock
}

// BeginRecheckTxSync provides a mock function with given fields: _a0
func (_m *AppConnMempool) BeginRecheckTxSync(_a0 types.RequestBeginRecheckTx) (*types.ResponseBeginRecheckTx, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseBeginRecheckTx
	if rf, ok := ret.Get(0).(func(types.RequestBeginRecheckTx) *types.ResponseBeginRecheckTx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseBeginRecheckTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestBeginRecheckTx) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTxAsync provides a mock function with given fields: _a0, _a1
func (_m *AppConnMempool) CheckTxAsync(_a0 types.RequestCheckTx, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(types.RequestCheckTx, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// CheckTxSync provides a mock function with given fields: _a0
func (_m *AppConnMempool) CheckTxSync(_a0 types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseCheckTx
	if rf, ok := ret.Get(0).(func(types.RequestCheckTx) *types.ResponseCheckTx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseCheckTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestCheckTx) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndRecheckTxSync provides a mock function with given fields: _a0
func (_m *AppConnMempool) EndRecheckTxSync(_a0 types.RequestEndRecheckTx) (*types.ResponseEndRecheckTx, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseEndRecheckTx
	if rf, ok := ret.Get(0).(func(types.RequestEndRecheckTx) *types.ResponseEndRecheckTx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseEndRecheckTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.RequestEndRecheckTx) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Error provides a mock function with given fields:
func (_m *AppConnMempool) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushAsync provides a mock function with given fields: _a0
func (_m *AppConnMempool) FlushAsync(_a0 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// FlushSync provides a mock function with given fields:
func (_m *AppConnMempool) FlushSync() (*types.ResponseFlush, error) {
	ret := _m.Called()

	var r0 *types.ResponseFlush
	if rf, ok := ret.Get(0).(func() *types.ResponseFlush); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseFlush)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetGlobalCallback provides a mock function with given fields: _a0
func (_m *AppConnMempool) SetGlobalCallback(_a0 abcicli.GlobalCallback) {
	_m.Called(_a0)
}
