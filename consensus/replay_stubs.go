package consensus

import (
	tmabci "github.com/tendermint/tendermint/abci/types"

	abci "github.com/line/ostracon/abci/types"
	"github.com/line/ostracon/libs/clist"
	mempl "github.com/line/ostracon/mempool"
	tmstate "github.com/line/ostracon/proto/ostracon/state"
	"github.com/line/ostracon/proxy"
	"github.com/line/ostracon/types"
)

//-----------------------------------------------------------------------------

type emptyMempool struct{}

var _ mempl.Mempool = emptyMempool{}

func (emptyMempool) Lock()     {}
func (emptyMempool) Unlock()   {}
func (emptyMempool) Size() int { return 0 }
func (emptyMempool) CheckTxSync(_ types.Tx, _ mempl.TxInfo) (*abci.Response, error) {
	return nil, nil
}
func (emptyMempool) CheckTxAsync(_ types.Tx, _ mempl.TxInfo, _ func(error), _ func(*abci.Response)) {
}
func (emptyMempool) ReapMaxBytesMaxGas(_, _ int64) types.Txs          { return types.Txs{} }
func (emptyMempool) ReapMaxBytesMaxGasMaxTxs(_, _, _ int64) types.Txs { return types.Txs{} }
func (emptyMempool) ReapMaxTxs(n int) types.Txs                       { return types.Txs{} }
func (emptyMempool) Update(
	_ *types.Block,
	_ []*tmabci.ResponseDeliverTx,
	_ mempl.PreCheckFunc,
	_ mempl.PostCheckFunc,
) error {
	return nil
}
func (emptyMempool) Flush()                        {}
func (emptyMempool) FlushAppConn() error           { return nil }
func (emptyMempool) TxsAvailable() <-chan struct{} { return make(chan struct{}) }
func (emptyMempool) EnableTxsAvailable()           {}
func (emptyMempool) TxsBytes() int64               { return 0 }

func (emptyMempool) TxsFront() *clist.CElement    { return nil }
func (emptyMempool) TxsWaitChan() <-chan struct{} { return nil }

func (emptyMempool) InitWAL() error { return nil }
func (emptyMempool) CloseWAL()      {}

//-----------------------------------------------------------------------------
// mockProxyApp uses ABCIResponses to give the right results.
//
// Useful because we don't want to call Commit() twice for the same block on
// the real app.

func newMockProxyApp(appHash []byte, abciResponses *tmstate.ABCIResponses) proxy.AppConnConsensus {
	clientCreator := proxy.NewLocalClientCreator(&mockProxyApp{
		appHash:       appHash,
		abciResponses: abciResponses,
	})
	cli, _ := clientCreator.NewABCIClient()
	err := cli.Start()
	if err != nil {
		panic(err)
	}
	return proxy.NewAppConnConsensus(cli)
}

type mockProxyApp struct {
	abci.BaseApplication

	appHash       []byte
	txCount       int
	abciResponses *tmstate.ABCIResponses
}

func (mock *mockProxyApp) DeliverTx(req tmabci.RequestDeliverTx) tmabci.ResponseDeliverTx {
	r := mock.abciResponses.DeliverTxs[mock.txCount]
	mock.txCount++
	if r == nil {
		return tmabci.ResponseDeliverTx{}
	}
	return *r
}

func (mock *mockProxyApp) EndBlock(req tmabci.RequestEndBlock) abci.ResponseEndBlock {
	mock.txCount = 0
	return *mock.abciResponses.EndBlock
}

func (mock *mockProxyApp) Commit() tmabci.ResponseCommit {
	return tmabci.ResponseCommit{Data: mock.appHash}
}
