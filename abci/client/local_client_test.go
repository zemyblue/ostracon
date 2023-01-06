package abcicli

import (
	"testing"
	"time"

	tmabci "github.com/tendermint/tendermint/abci/types"

	abci "github.com/line/ostracon/abci/types"
	"github.com/stretchr/testify/require"
)

type sampleApp struct {
	abci.BaseApplication
}

func newDoneChan(t *testing.T) chan struct{} {
	result := make(chan struct{})
	go func() {
		select {
		case <-time.After(time.Second):
			require.Fail(t, "callback is not called for a second")
		case <-result:
			return
		}
	}()
	return result
}

func getResponseCallback(t *testing.T) ResponseCallback {
	doneChan := newDoneChan(t)
	return func(res *abci.Response) {
		require.NotNil(t, res)
		doneChan <- struct{}{}
	}
}

func TestLocalClientCalls(t *testing.T) {
	app := sampleApp{}
	c := NewLocalClient(nil, app)

	c.SetGlobalCallback(func(*abci.Request, *abci.Response) {
	})

	c.EchoAsync("msg", getResponseCallback(t))
	c.FlushAsync(getResponseCallback(t))
	c.InfoAsync(tmabci.RequestInfo{}, getResponseCallback(t))
	c.SetOptionAsync(tmabci.RequestSetOption{}, getResponseCallback(t))
	c.DeliverTxAsync(tmabci.RequestDeliverTx{}, getResponseCallback(t))
	c.CheckTxAsync(tmabci.RequestCheckTx{}, getResponseCallback(t))
	c.QueryAsync(tmabci.RequestQuery{}, getResponseCallback(t))
	c.CommitAsync(getResponseCallback(t))
	c.InitChainAsync(abci.RequestInitChain{}, getResponseCallback(t))
	c.BeginBlockAsync(abci.RequestBeginBlock{}, getResponseCallback(t))
	c.EndBlockAsync(tmabci.RequestEndBlock{}, getResponseCallback(t))
	c.BeginRecheckTxAsync(abci.RequestBeginRecheckTx{}, getResponseCallback(t))
	c.EndRecheckTxAsync(abci.RequestEndRecheckTx{}, getResponseCallback(t))
	c.ListSnapshotsAsync(tmabci.RequestListSnapshots{}, getResponseCallback(t))
	c.OfferSnapshotAsync(tmabci.RequestOfferSnapshot{}, getResponseCallback(t))
	c.LoadSnapshotChunkAsync(tmabci.RequestLoadSnapshotChunk{}, getResponseCallback(t))
	c.ApplySnapshotChunkAsync(tmabci.RequestApplySnapshotChunk{}, getResponseCallback(t))

	_, err := c.EchoSync("msg")
	require.NoError(t, err)

	_, err = c.FlushSync()
	require.NoError(t, err)

	_, err = c.InfoSync(tmabci.RequestInfo{})
	require.NoError(t, err)

	_, err = c.SetOptionSync(tmabci.RequestSetOption{})
	require.NoError(t, err)

	_, err = c.DeliverTxSync(tmabci.RequestDeliverTx{})
	require.NoError(t, err)

	_, err = c.CheckTxSync(tmabci.RequestCheckTx{})
	require.NoError(t, err)

	_, err = c.QuerySync(tmabci.RequestQuery{})
	require.NoError(t, err)

	_, err = c.CommitSync()
	require.NoError(t, err)

	_, err = c.InitChainSync(abci.RequestInitChain{})
	require.NoError(t, err)

	_, err = c.BeginBlockSync(abci.RequestBeginBlock{})
	require.NoError(t, err)

	_, err = c.EndBlockSync(tmabci.RequestEndBlock{})
	require.NoError(t, err)

	_, err = c.BeginRecheckTxSync(abci.RequestBeginRecheckTx{})
	require.NoError(t, err)

	_, err = c.EndRecheckTxSync(abci.RequestEndRecheckTx{})
	require.NoError(t, err)

	_, err = c.ListSnapshotsSync(tmabci.RequestListSnapshots{})
	require.NoError(t, err)

	_, err = c.OfferSnapshotSync(tmabci.RequestOfferSnapshot{})
	require.NoError(t, err)

	_, err = c.LoadSnapshotChunkSync(tmabci.RequestLoadSnapshotChunk{})
	require.NoError(t, err)

	_, err = c.ApplySnapshotChunkSync(tmabci.RequestApplySnapshotChunk{})
	require.NoError(t, err)
}
