// This program performs tests and benchmarks that SignerClient can connect to KMS and make API calls.
// To test, address the KMS connection to port 45666 on the machine running this program and run the following:
//
// $ cd test/kms
// $ go test -tags libsodium -bench . -benchmem
//
package main

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/line/ostracon/config"
	"github.com/line/ostracon/crypto"
	"github.com/line/ostracon/crypto/ed25519"
	"github.com/line/ostracon/crypto/vrf"
	"github.com/line/ostracon/libs/log"
	tmnet "github.com/line/ostracon/libs/net"
	"github.com/line/ostracon/node"
	"github.com/line/ostracon/privval"
	privvalproto "github.com/line/ostracon/proto/ostracon/privval"
	types2 "github.com/line/ostracon/proto/ostracon/types"
	"github.com/line/ostracon/types"
	"github.com/stretchr/testify/require"
)

var logger = log.NewOCLogger(log.NewSyncWriter(os.Stdout))

const chainID = "test-chain"
const listenAddr = "tcp://0.0.0.0:45666"

func BenchmarkKMS(b *testing.B) {
	chainID := "test-chain"
	protocol, address := tmnet.ProtocolAndAddress(listenAddr)
	ln, err := net.Listen(protocol, address)
	require.NoError(b, err)
	listener := privval.NewTCPListener(ln, ed25519.GenPrivKeyFromSecret([]byte("🏺")))
	endpoint := privval.NewSignerListenerEndpoint(logger, listener)
	client, err := privval.NewSignerClient(endpoint, chainID)
	require.NoError(b, err)

	// ensure connection and warm up
	b.Run("Ping", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ping(endpoint)
		}
		b.StopTimer()
	})

	benchmarkPrivValidator(b, client)
}

func BenchmarkFilePV(b *testing.B) {
	cfg := config.ResetTestRoot("BenchmarkFilePV")
	defer func() {
		var _ = os.RemoveAll(cfg.RootDir)
	}()

	n, err := node.NewOstraconNode(cfg, logger)
	require.NoError(b, err)

	benchmarkPrivValidator(b, n.PrivValidator())
}

func benchmarkPrivValidator(b *testing.B, pv types.PrivValidator) {
	pubKey := benchmarkGetPubKey(b, pv)
	benchmarkSignVote(b, pv, pubKey)
	benchmarkSignProposal(b, pv, pubKey)
	benchmarkVRFProof(b, pv, pubKey)
}

func benchmarkGetPubKey(b *testing.B, pv types.PrivValidator) crypto.PubKey {
	var pubKey crypto.PubKey
	var err error

	// performance measurement
	b.Run("GetPubKey", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pubKey, err = pv.GetPubKey()
		}
	})

	// evaluate execution results
	require.NoError(b, err)
	require.Equalf(b, ed25519.PubKeySize, len(pubKey.Bytes()), "PubKey: public key size = %d != %d",
		ed25519.PubKeySize, len(pubKey.Bytes()))
	return pubKey
}

func benchmarkSignVote(b *testing.B, pv types.PrivValidator, pubKey crypto.PubKey) {
	blockID := types.BlockID{
		Hash: make([]byte, 32),
		PartSetHeader: types.PartSetHeader{
			Total: 10,
			Hash:  make([]byte, 32),
		},
	}
	vote := types.Vote{
		Type:             types2.PrevoteType,
		Height:           1,
		Round:            0,
		BlockID:          blockID,
		Timestamp:        time.Now(),
		ValidatorAddress: pubKey.Address(),
		ValidatorIndex:   0,
		Signature:        nil,
	}
	pb := vote.ToProto()
	var err error

	// performance measurement
	b.Run("SignVote", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err = pv.SignVote(chainID, pb)
		}
	})

	// evaluate execution results
	require.NoError(b, err)
	require.Equalf(b, ed25519.SignatureSize, len(pb.Signature), "SignVote: signature size = %d != %d",
		ed25519.SignatureSize, len(pb.Signature))
	bytes := types.VoteSignBytes(chainID, pb)
	require.Truef(b, pubKey.VerifySignature(bytes, pb.Signature), "SignVote: signature verification")
}

func benchmarkSignProposal(b *testing.B, pv types.PrivValidator, pubKey crypto.PubKey) {
	blockID := types.BlockID{
		Hash: make([]byte, 32),
		PartSetHeader: types.PartSetHeader{
			Total: 10,
			Hash:  make([]byte, 32),
		},
	}
	proposal := types.Proposal{
		Type:      types2.ProposalType,
		Height:    2,
		Round:     0,
		POLRound:  -1,
		BlockID:   blockID,
		Timestamp: time.Now(),
		Signature: nil,
	}
	pb := proposal.ToProto()
	var err error

	// performance measurement
	b.Run("SignProposal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err = pv.SignProposal(chainID, pb)
		}
	})

	// evaluate execution results
	require.NoError(b, err)
	require.Equalf(b, ed25519.SignatureSize, len(pb.Signature), "SignProposal: signature size = %d != %d",
		ed25519.SignatureSize, len(pb.Signature))
	bytes := types.ProposalSignBytes(chainID, pb)
	require.Truef(b, pubKey.VerifySignature(bytes, pb.Signature), "SignProposal: signature verification")
}

func benchmarkVRFProof(b *testing.B, pv types.PrivValidator, pubKey crypto.PubKey) {
	message := []byte("hello, world")
	var proof crypto.Proof
	var err error

	// performance measurement
	b.Run("VRFProof", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			proof, err = pv.GenerateVRFProof(message)
		}
	})

	// evaluate execution results
	require.NoError(b, err)
	require.Equalf(b, vrf.ProofSize, len(proof), "VRFProof: proof size = %d != %d", len(proof), vrf.ProofSize)
	output, err := pubKey.VRFVerify(proof, message)
	require.NoError(b, err)
	require.Equalf(b, vrf.OutputSize, len(output), "VRFProof: output size = %d != %d", len(output), vrf.OutputSize)
}

func ping(sl *privval.SignerListenerEndpoint) {
	msg := privvalproto.Message{
		Sum: &privvalproto.Message_PingRequest{
			PingRequest: &privvalproto.PingRequest{},
		},
	}
	_, err := sl.SendRequest(msg)
	if err != nil {
		sl.Logger.Error("Benchmark::ping", "err", err)
	}
}
