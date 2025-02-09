package statesync

import (
	"context"
	"fmt"
	"github.com/line/ostracon/config"
	"github.com/line/ostracon/libs/log"
	tmrand "github.com/line/ostracon/libs/rand"
	"github.com/line/ostracon/light"
	"github.com/line/ostracon/proto/ostracon/state"
	tmproto "github.com/line/ostracon/proto/ostracon/types"
	tmversion "github.com/line/ostracon/proto/ostracon/version"
	ctypes "github.com/line/ostracon/rpc/core/types"
	rpcserver "github.com/line/ostracon/rpc/jsonrpc/server"
	rpctypes "github.com/line/ostracon/rpc/jsonrpc/types"
	"github.com/line/ostracon/types"
	tmtime "github.com/line/ostracon/types/time"
	"github.com/line/ostracon/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestNewLightClientStateProvider(t *testing.T) {
	setupVars(t)
	cfg.SetRoot(os.TempDir())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	listeners, servers, closeListenersFunc := serveTestRPCServers(t, cfg, 2)
	defer closeListenersFunc(listeners)
	type args struct {
		ctx           context.Context
		chainID       string
		version       state.Version
		initialHeight int64
		servers       []string
		trustOptions  light.TrustOptions
		logger        log.Logger
	}
	successFunc := func(t assert.TestingT, err error, i ...interface{}) bool {
		return assert.NoError(t, err)
	}
	serversErrorFunc := func(t assert.TestingT, err error, i ...interface{}) bool {
		return assert.Error(t, err) &&
			assert.Contains(t, err.Error(), "at least 2 RPC servers are required, got ")
	}
	genesisErrorFunc := func(t assert.TestingT, err error, i ...interface{}) bool {
		return assert.Error(t, err) &&
			assert.Contains(t, err.Error(), "failed to retrieve genesis doc: ")
	}
	lightErrorFunc := func(t assert.TestingT, err error, i ...interface{}) bool {
		return assert.Error(t, err) &&
			assert.Contains(t, err.Error(), "invalid TrustOptions: negative or zero period")
	}
	tests := []struct {
		name    string
		args    args
		want    StateProvider
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				chainID: chainId,
				servers: servers,
				logger:  log.NewNopLogger(),
				trustOptions: light.TrustOptions{
					Period: cfg.StateSync.TrustPeriod,
					Height: 1,
					Hash:   header.Hash(),
				}},
			want:    &lightClientStateProvider{},
			wantErr: successFunc,
		},
		{
			name:    "empty servers",
			args:    args{},
			want:    nil,
			wantErr: serversErrorFunc,
		},
		{
			name:    "duplicated servers",
			args:    args{servers: []string{"a", "a"}},
			want:    nil,
			wantErr: serversErrorFunc,
		},
		{
			name:    "fail genesis",
			args:    args{servers: []string{"a", "b"}},
			want:    nil,
			wantErr: genesisErrorFunc,
		},
		{
			name:    "fail light client",
			args:    args{ctx: ctx, servers: servers},
			want:    nil,
			wantErr: lightErrorFunc,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLightClientStateProvider(
				tt.args.ctx,
				tt.args.chainID,
				tt.args.version,
				tt.args.initialHeight,
				tt.args.servers,
				tt.args.trustOptions,
				tt.args.logger)
			if !tt.wantErr(t, err) {
				return
			}
			assert.IsType(t, tt.want, got)
		})
	}
}

const (
	height = int64(1)
	round  = int32(0)
	size   = 1
	index  = int32(0)
)

var (
	chainId       string
	cfg           *config.Config
	genDoc        *types.GenesisDoc
	privVals      []*types.PrivValidator
	vals          []*types.Validator
	votersIndices []int32
	header        *types.Header
	commit        *types.Commit
)

func setupVars(t *testing.T) {
	// config
	chainId = fmt.Sprintf("test-chain-%v", tmrand.Str(6))
	cfg = config.TestConfig()
	// getDoc
	genDoc = &types.GenesisDoc{
		ChainID:         chainId,
		GenesisTime:     tmtime.Now(),
		ConsensusParams: types.DefaultConsensusParams(),
		VoterParams:     types.DefaultVoterParams(),
	}
	// validators
	privVals = make([]*types.PrivValidator, size)
	vals = make([]*types.Validator, size)
	votersIndices = make([]int32, size)
	for i := 0; i < size; i++ {
		val, privVal := types.RandValidator(true, 1)
		val.VotingWeight = val.VotingPower
		privVals[i] = &privVal
		vals[i] = val
		votersIndices[i] = int32(i)
	}
	// header
	valSet, err := types.ValidatorSetFromExistingValidators(vals)
	require.NoError(t, err)
	voterSet := types.ToVoterAll(vals)
	header = &types.Header{
		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		ChainID:         chainId,
		Height:          height,
		VotersHash:      voterSet.Hash(),
		ValidatorsHash:  valSet.Hash(),
		ProposerAddress: vals[index].Address,
		Round:           round,
	}
	// block id
	hash := tmrand.Bytes(32)
	blockId := types.BlockID{
		Hash: header.Hash(),
		PartSetHeader: types.PartSetHeader{
			Total: 1,
			Hash:  hash,
		},
	}
	// vote
	vote := &types.Vote{
		ValidatorAddress: vals[index].Address,
		ValidatorIndex:   index,
		Height:           height,
		Round:            round,
		Timestamp:        tmtime.Now(),
		Type:             tmproto.PrecommitType,
		BlockID:          blockId,
	}
	v := vote.ToProto()
	require.NoError(t, (*privVals[index]).SignVote(chainId, v))
	vote.Signature = v.Signature
	vote.Timestamp = v.Timestamp
	// commit
	commit = &types.Commit{
		Height:  height,
		Round:   round,
		BlockID: blockId,
		Signatures: []types.CommitSig{
			{
				BlockIDFlag:      types.BlockIDFlagCommit,
				ValidatorAddress: vote.ValidatorAddress,
				Timestamp:        vote.Timestamp,
				Signature:        vote.Signature,
			},
		},
	}
}

func serveTestRPCServers(t *testing.T, config *config.Config, num int,
) (listeners []*net.Listener, servers []string, closeListenersFunc func(listeners []*net.Listener)) {
	// Start the RPC server
	mux := http.NewServeMux()
	rpcserver.RegisterRPCFuncs(mux, routes, log.TestingLogger())
	wm := rpcserver.NewWebsocketManager(routes)
	mux.HandleFunc("/websocket", wm.WebsocketHandler)
	rpcConfig := rpcserver.DefaultConfig()
	listeners = make([]*net.Listener, num)
	servers = make([]string, num)
	for i := 0; i < num; i++ {
		listener, err := rpcserver.Listen("tcp://127.0.0.1:0", rpcConfig)
		require.NoError(t, err)
		listeners[i] = &listener
		servers[i] = listener.Addr().String()
		go func() {
			_ = rpcserver.Serve(listener, mux, log.NewNopLogger(), rpcConfig)
		}()
	}
	closeListenersFunc = func(listeners []*net.Listener) {
		for _, listener := range listeners {
			require.NoError(t, (*listener).Close())
		}
	}
	return listeners, servers, closeListenersFunc
}

var routes = map[string]*rpcserver.RPCFunc{
	"genesis":           rpcserver.NewRPCFunc(genesisFunc, ""),
	"commit":            rpcserver.NewRPCFunc(commitFunc, "height"),
	"validators_voters": rpcserver.NewRPCFunc(validatorsFunc, "height,page,per_page"),
}

func genesisFunc(ctx *rpctypes.Context) (*ctypes.ResultGenesis, error) {
	return &ctypes.ResultGenesis{Genesis: genDoc}, nil
}

func commitFunc(ctx *rpctypes.Context, heightPtr *int64) (*ctypes.ResultCommit, error) {
	return ctypes.NewResultCommit(header, commit, true), nil
}

func validatorsFunc(ctx *rpctypes.Context, heightPtr *int64, pagePtr, perPagePtr *int,
) (*ctypes.ResultValidatorsWithVoters, error) {
	return &ctypes.ResultValidatorsWithVoters{
		BlockHeight:  height,
		Validators:   vals,
		VoterIndices: votersIndices,
		Count:        size,
		Total:        size,
	}, nil
}
