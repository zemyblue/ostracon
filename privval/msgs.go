package privval

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	privvalproto "github.com/tendermint/tendermint/proto/tendermint/privval"

	oprivvalproto "github.com/line/ostracon/proto/ostracon/privval"
)

// TODO: Add ChainIDRequest

func mustWrapMsg(pb proto.Message) oprivvalproto.Message {
	msg := oprivvalproto.Message{}

	switch pb := pb.(type) {
	case *oprivvalproto.Message:
		msg = *pb
	case *privvalproto.PubKeyRequest:
		msg.Sum = &oprivvalproto.Message_PubKeyRequest{PubKeyRequest: pb}
	case *oprivvalproto.PubKeyResponse:
		msg.Sum = &oprivvalproto.Message_PubKeyResponse{PubKeyResponse: pb}
	case *privvalproto.SignVoteRequest:
		msg.Sum = &oprivvalproto.Message_SignVoteRequest{SignVoteRequest: pb}
	case *privvalproto.SignedVoteResponse:
		msg.Sum = &oprivvalproto.Message_SignedVoteResponse{SignedVoteResponse: pb}
	case *privvalproto.SignedProposalResponse:
		msg.Sum = &oprivvalproto.Message_SignedProposalResponse{SignedProposalResponse: pb}
	case *privvalproto.SignProposalRequest:
		msg.Sum = &oprivvalproto.Message_SignProposalRequest{SignProposalRequest: pb}
	case *oprivvalproto.VRFProofRequest:
		msg.Sum = &oprivvalproto.Message_VrfProofRequest{VrfProofRequest: pb}
	case *oprivvalproto.VRFProofResponse:
		msg.Sum = &oprivvalproto.Message_VrfProofResponse{VrfProofResponse: pb}
	case *privvalproto.PingRequest:
		msg.Sum = &oprivvalproto.Message_PingRequest{PingRequest: pb}
	case *privvalproto.PingResponse:
		msg.Sum = &oprivvalproto.Message_PingResponse{PingResponse: pb}
	default:
		panic(fmt.Errorf("unknown message type %T", pb))
	}

	return msg
}
