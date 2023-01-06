package coregrpc

import (
	"context"

	tmabci "github.com/tendermint/tendermint/abci/types"
	tmgrpc "github.com/tendermint/tendermint/rpc/grpc"

	abci "github.com/line/ostracon/abci/types"
	core "github.com/line/ostracon/rpc/core"
	rpctypes "github.com/line/ostracon/rpc/jsonrpc/types"
)

type broadcastAPI struct {
}

func (bapi *broadcastAPI) Ping(ctx context.Context, req *tmgrpc.RequestPing) (*tmgrpc.ResponsePing, error) {
	// kvstore so we can check if the server is up
	return &tmgrpc.ResponsePing{}, nil
}

func (bapi *broadcastAPI) BroadcastTx(ctx context.Context, req *tmgrpc.RequestBroadcastTx) (*ResponseBroadcastTx, error) {
	// NOTE: there's no way to get client's remote address
	// see https://stackoverflow.com/questions/33684570/session-and-remote-ip-address-in-grpc-go
	res, err := core.BroadcastTxCommit(&rpctypes.Context{}, req.Tx)
	if err != nil {
		return nil, err
	}

	return &ResponseBroadcastTx{
		CheckTx: &abci.ResponseCheckTx{
			Code: res.CheckTx.Code,
			Data: res.CheckTx.Data,
			Log:  res.CheckTx.Log,
		},
		DeliverTx: &tmabci.ResponseDeliverTx{
			Code: res.DeliverTx.Code,
			Data: res.DeliverTx.Data,
			Log:  res.DeliverTx.Log,
		},
	}, nil
}
