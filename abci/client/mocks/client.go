// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	abci "github.com/tendermint/tendermint/abci/types"

	abcicli "github.com/line/ostracon/abci/client"
	ocabci "github.com/line/ostracon/abci/types"
	log "github.com/line/ostracon/libs/log"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// ApplySnapshotChunkAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) ApplySnapshotChunkAsync(_a0 abci.RequestApplySnapshotChunk, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestApplySnapshotChunk, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// ApplySnapshotChunkSync provides a mock function with given fields: _a0
func (_m *Client) ApplySnapshotChunkSync(_a0 abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseApplySnapshotChunk
	if rf, ok := ret.Get(0).(func(abci.RequestApplySnapshotChunk) *abci.ResponseApplySnapshotChunk); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseApplySnapshotChunk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestApplySnapshotChunk) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginBlockAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) BeginBlockAsync(_a0 ocabci.RequestBeginBlock, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(ocabci.RequestBeginBlock, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// BeginBlockSync provides a mock function with given fields: _a0
func (_m *Client) BeginBlockSync(_a0 ocabci.RequestBeginBlock) (*abci.ResponseBeginBlock, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseBeginBlock
	if rf, ok := ret.Get(0).(func(ocabci.RequestBeginBlock) *abci.ResponseBeginBlock); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseBeginBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ocabci.RequestBeginBlock) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginRecheckTxAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) BeginRecheckTxAsync(_a0 ocabci.RequestBeginRecheckTx, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(ocabci.RequestBeginRecheckTx, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// BeginRecheckTxSync provides a mock function with given fields: _a0
func (_m *Client) BeginRecheckTxSync(_a0 ocabci.RequestBeginRecheckTx) (*ocabci.ResponseBeginRecheckTx, error) {
	ret := _m.Called(_a0)

	var r0 *ocabci.ResponseBeginRecheckTx
	if rf, ok := ret.Get(0).(func(ocabci.RequestBeginRecheckTx) *ocabci.ResponseBeginRecheckTx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ocabci.ResponseBeginRecheckTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ocabci.RequestBeginRecheckTx) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTxAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) CheckTxAsync(_a0 abci.RequestCheckTx, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestCheckTx, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// CheckTxSync provides a mock function with given fields: _a0
func (_m *Client) CheckTxSync(_a0 abci.RequestCheckTx) (*ocabci.ResponseCheckTx, error) {
	ret := _m.Called(_a0)

	var r0 *ocabci.ResponseCheckTx
	if rf, ok := ret.Get(0).(func(abci.RequestCheckTx) *ocabci.ResponseCheckTx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ocabci.ResponseCheckTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestCheckTx) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommitAsync provides a mock function with given fields: _a0
func (_m *Client) CommitAsync(_a0 abcicli.ResponseCallback) *abcicli.ReqRes {
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

// CommitSync provides a mock function with given fields:
func (_m *Client) CommitSync() (*abci.ResponseCommit, error) {
	ret := _m.Called()

	var r0 *abci.ResponseCommit
	if rf, ok := ret.Get(0).(func() *abci.ResponseCommit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseCommit)
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

// DeliverTxAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) DeliverTxAsync(_a0 abci.RequestDeliverTx, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestDeliverTx, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// DeliverTxSync provides a mock function with given fields: _a0
func (_m *Client) DeliverTxSync(_a0 abci.RequestDeliverTx) (*abci.ResponseDeliverTx, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseDeliverTx
	if rf, ok := ret.Get(0).(func(abci.RequestDeliverTx) *abci.ResponseDeliverTx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseDeliverTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestDeliverTx) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EchoAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) EchoAsync(_a0 string, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(string, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// EchoSync provides a mock function with given fields: _a0
func (_m *Client) EchoSync(_a0 string) (*abci.ResponseEcho, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseEcho
	if rf, ok := ret.Get(0).(func(string) *abci.ResponseEcho); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseEcho)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndBlockAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) EndBlockAsync(_a0 abci.RequestEndBlock, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestEndBlock, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// EndBlockSync provides a mock function with given fields: _a0
func (_m *Client) EndBlockSync(_a0 abci.RequestEndBlock) (*ocabci.ResponseEndBlock, error) {
	ret := _m.Called(_a0)

	var r0 *ocabci.ResponseEndBlock
	if rf, ok := ret.Get(0).(func(abci.RequestEndBlock) *ocabci.ResponseEndBlock); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ocabci.ResponseEndBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestEndBlock) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndRecheckTxAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) EndRecheckTxAsync(_a0 ocabci.RequestEndRecheckTx, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(ocabci.RequestEndRecheckTx, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// EndRecheckTxSync provides a mock function with given fields: _a0
func (_m *Client) EndRecheckTxSync(_a0 ocabci.RequestEndRecheckTx) (*ocabci.ResponseEndRecheckTx, error) {
	ret := _m.Called(_a0)

	var r0 *ocabci.ResponseEndRecheckTx
	if rf, ok := ret.Get(0).(func(ocabci.RequestEndRecheckTx) *ocabci.ResponseEndRecheckTx); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ocabci.ResponseEndRecheckTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ocabci.RequestEndRecheckTx) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Error provides a mock function with given fields:
func (_m *Client) Error() error {
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
func (_m *Client) FlushAsync(_a0 abcicli.ResponseCallback) *abcicli.ReqRes {
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
func (_m *Client) FlushSync() (*abci.ResponseFlush, error) {
	ret := _m.Called()

	var r0 *abci.ResponseFlush
	if rf, ok := ret.Get(0).(func() *abci.ResponseFlush); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseFlush)
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

// GetGlobalCallback provides a mock function with given fields:
func (_m *Client) GetGlobalCallback() abcicli.GlobalCallback {
	ret := _m.Called()

	var r0 abcicli.GlobalCallback
	if rf, ok := ret.Get(0).(func() abcicli.GlobalCallback); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(abcicli.GlobalCallback)
		}
	}

	return r0
}

// InfoAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) InfoAsync(_a0 abci.RequestInfo, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestInfo, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// InfoSync provides a mock function with given fields: _a0
func (_m *Client) InfoSync(_a0 abci.RequestInfo) (*abci.ResponseInfo, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseInfo
	if rf, ok := ret.Get(0).(func(abci.RequestInfo) *abci.ResponseInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestInfo) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitChainAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) InitChainAsync(_a0 ocabci.RequestInitChain, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(ocabci.RequestInitChain, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// InitChainSync provides a mock function with given fields: _a0
func (_m *Client) InitChainSync(_a0 ocabci.RequestInitChain) (*ocabci.ResponseInitChain, error) {
	ret := _m.Called(_a0)

	var r0 *ocabci.ResponseInitChain
	if rf, ok := ret.Get(0).(func(ocabci.RequestInitChain) *ocabci.ResponseInitChain); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ocabci.ResponseInitChain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ocabci.RequestInitChain) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRunning provides a mock function with given fields:
func (_m *Client) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListSnapshotsAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) ListSnapshotsAsync(_a0 abci.RequestListSnapshots, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestListSnapshots, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// ListSnapshotsSync provides a mock function with given fields: _a0
func (_m *Client) ListSnapshotsSync(_a0 abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseListSnapshots
	if rf, ok := ret.Get(0).(func(abci.RequestListSnapshots) *abci.ResponseListSnapshots); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseListSnapshots)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestListSnapshots) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadSnapshotChunkAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) LoadSnapshotChunkAsync(_a0 abci.RequestLoadSnapshotChunk, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestLoadSnapshotChunk, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// LoadSnapshotChunkSync provides a mock function with given fields: _a0
func (_m *Client) LoadSnapshotChunkSync(_a0 abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseLoadSnapshotChunk
	if rf, ok := ret.Get(0).(func(abci.RequestLoadSnapshotChunk) *abci.ResponseLoadSnapshotChunk); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseLoadSnapshotChunk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestLoadSnapshotChunk) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OfferSnapshotAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) OfferSnapshotAsync(_a0 abci.RequestOfferSnapshot, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestOfferSnapshot, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// OfferSnapshotSync provides a mock function with given fields: _a0
func (_m *Client) OfferSnapshotSync(_a0 abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseOfferSnapshot
	if rf, ok := ret.Get(0).(func(abci.RequestOfferSnapshot) *abci.ResponseOfferSnapshot); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseOfferSnapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestOfferSnapshot) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnReset provides a mock function with given fields:
func (_m *Client) OnReset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStart provides a mock function with given fields:
func (_m *Client) OnStart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStop provides a mock function with given fields:
func (_m *Client) OnStop() {
	_m.Called()
}

// QueryAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) QueryAsync(_a0 abci.RequestQuery, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestQuery, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// QuerySync provides a mock function with given fields: _a0
func (_m *Client) QuerySync(_a0 abci.RequestQuery) (*abci.ResponseQuery, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseQuery
	if rf, ok := ret.Get(0).(func(abci.RequestQuery) *abci.ResponseQuery); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseQuery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestQuery) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Quit provides a mock function with given fields:
func (_m *Client) Quit() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Reset provides a mock function with given fields:
func (_m *Client) Reset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGlobalCallback provides a mock function with given fields: _a0
func (_m *Client) SetGlobalCallback(_a0 abcicli.GlobalCallback) {
	_m.Called(_a0)
}

// SetLogger provides a mock function with given fields: _a0
func (_m *Client) SetLogger(_a0 log.Logger) {
	_m.Called(_a0)
}

// SetOptionAsync provides a mock function with given fields: _a0, _a1
func (_m *Client) SetOptionAsync(_a0 abci.RequestSetOption, _a1 abcicli.ResponseCallback) *abcicli.ReqRes {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	if rf, ok := ret.Get(0).(func(abci.RequestSetOption, abcicli.ResponseCallback) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	return r0
}

// SetOptionSync provides a mock function with given fields: _a0
func (_m *Client) SetOptionSync(_a0 abci.RequestSetOption) (*abci.ResponseSetOption, error) {
	ret := _m.Called(_a0)

	var r0 *abci.ResponseSetOption
	if rf, ok := ret.Get(0).(func(abci.RequestSetOption) *abci.ResponseSetOption); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abci.ResponseSetOption)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abci.RequestSetOption) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Client) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Client) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Client) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
