package types

import (
	"bytes"
	"fmt"
	"sort"
	"testing"

	"github.com/line/ostracon/crypto/bls"
	"github.com/line/ostracon/crypto/composite"
	"github.com/line/ostracon/crypto/ed25519"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/line/ostracon/crypto"
	tmrand "github.com/line/ostracon/libs/rand"
	tmproto "github.com/line/ostracon/proto/ostracon/types"
	tmtime "github.com/line/ostracon/types/time"
)

// NOTE: privValidators are in order
func randVoteSet(
	height int64,
	round int32,
	signedMsgType tmproto.SignedMsgType,
	numValidators int,
	votingPower int64,
) (*VoteSet, *ValidatorSet, *VoterSet, []PrivValidator) {
	valSet, voterSet, privValidators := RandVoterSet(numValidators, votingPower)
	return NewVoteSet("test_chain_id", height, round, signedMsgType, voterSet), valSet, voterSet, privValidators
}

func randVoteSetForPrivKeys(
	height int64,
	round int32,
	signedMsgType tmproto.SignedMsgType,
	privKeys []crypto.PrivKey,
	votingPower int64,
) (*VoteSet, *ValidatorSet, *VoterSet, []PrivValidator) {
	valz := make([]*Validator, len(privKeys))
	privValidators := make([]PrivValidator, len(privKeys))
	for i := 0; i < len(privKeys); i++ {
		privVal := MockPV{privKeys[i], false, false}
		pubKey, err := privVal.GetPubKey()
		if err != nil {
			panic(fmt.Errorf("could not retrieve pubkey %w", err))
		}
		val := NewValidator(pubKey, votingPower)
		valz[i] = val
		privValidators[i] = privVal
	}
	vals := NewValidatorSet(valz)
	sort.Sort(PrivValidatorsByAddress(privValidators))
	voterSet := SelectVoter(vals, []byte{}, DefaultVoterParams())
	return NewVoteSet("test_chain_id", height, round, signedMsgType, voterSet), vals, voterSet, privValidators
}

func addVoteByAllVoterSet(t *testing.T, voteSet *VoteSet, privVals []PrivValidator, height int64, round int32) []Vote {
	votes := make([]Vote, voteSet.Size())
	for i, voter := range voteSet.voterSet.Voters {
		vote := &Vote{
			Type:             tmproto.PrevoteType,
			Height:           height,
			Round:            round,
			BlockID:          randBlockID(),
			Timestamp:        tmtime.Now(),
			ValidatorAddress: voter.Address,
			ValidatorIndex:   int32(i),
		}
		v := vote.ToProto()
		err := privVals[i].SignVote(voteSet.chainID, v)
		require.NoError(t, err)
		vote.Signature = v.Signature
		added, err := voteSet.AddVote(vote)
		require.NoError(t, err)
		require.True(t, added)
		votes[i] = *vote
	}
	return votes
}

// Convenience: Return new vote with different validator address/index
func withValidator(vote *Vote, addr []byte, idx int32) *Vote {
	vote = vote.Copy()
	vote.ValidatorAddress = addr
	vote.ValidatorIndex = idx
	return vote
}

// Convenience: Return new vote with different height
func withHeight(vote *Vote, height int64) *Vote {
	vote = vote.Copy()
	vote.Height = height
	return vote
}

// Convenience: Return new vote with different round
func withRound(vote *Vote, round int32) *Vote {
	vote = vote.Copy()
	vote.Round = round
	return vote
}

// Convenience: Return new vote with different type
func withType(vote *Vote, signedMsgType byte) *Vote {
	vote = vote.Copy()
	vote.Type = tmproto.SignedMsgType(signedMsgType)
	return vote
}

// Convenience: Return new vote with different blockHash
func withBlockHash(vote *Vote, blockHash []byte) *Vote {
	vote = vote.Copy()
	vote.BlockID.Hash = blockHash
	return vote
}

// Convenience: Return new vote with different blockParts
func withBlockPartSetHeader(vote *Vote, blockPartsHeader PartSetHeader) *Vote {
	vote = vote.Copy()
	vote.BlockID.PartSetHeader = blockPartsHeader
	return vote
}

func TestVoteSet_AddVote_Good(t *testing.T) {
	height, round := int64(1), int32(0)
	voteSet, _, _, privValidators := randVoteSet(height, round, tmproto.PrevoteType, 10, 1)
	val0 := privValidators[0]

	val0p, err := val0.GetPubKey()
	require.NoError(t, err)
	val0Addr := val0p.Address()

	assert.Nil(t, voteSet.GetByAddress(val0Addr))
	assert.False(t, voteSet.BitArray().GetIndex(0))
	blockID, ok := voteSet.TwoThirdsMajority()
	assert.False(t, ok || !blockID.IsZero(), "there should be no 2/3 majority")

	vote := &Vote{
		ValidatorAddress: val0Addr,
		ValidatorIndex:   0, // since privValidators are in order
		Height:           height,
		Round:            round,
		Type:             tmproto.PrevoteType,
		Timestamp:        tmtime.Now(),
		BlockID:          BlockID{nil, PartSetHeader{}},
	}
	_, err = signAddVote(val0, vote, voteSet)
	require.NoError(t, err)

	assert.NotNil(t, voteSet.GetByAddress(val0Addr))
	assert.True(t, voteSet.BitArray().GetIndex(0))
	blockID, ok = voteSet.TwoThirdsMajority()
	assert.False(t, ok || !blockID.IsZero(), "there should be no 2/3 majority")
}

func TestVoteSet_AddVote_Bad(t *testing.T) {
	height, round := int64(1), int32(0)
	voteSet, _, _, privValidators := randVoteSet(height, round, tmproto.PrevoteType, 10, 1)

	voteProto := &Vote{
		ValidatorAddress: nil,
		ValidatorIndex:   -1,
		Height:           height,
		Round:            round,
		Timestamp:        tmtime.Now(),
		Type:             tmproto.PrevoteType,
		BlockID:          BlockID{nil, PartSetHeader{}},
		Signature:        []byte{},
	}

	// val0 votes for nil.
	{
		pubKey, err := privValidators[0].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 0)
		added, err := signAddVote(privValidators[0], vote, voteSet)
		if !added || err != nil {
			t.Errorf("expected VoteSet.Add to succeed")
		}
	}

	// val0 votes again for some block.
	{
		pubKey, err := privValidators[0].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 0)
		added, err := signAddVote(privValidators[0], withBlockHash(vote, tmrand.Bytes(32)), voteSet)
		if added || err == nil {
			t.Errorf("expected VoteSet.Add to fail, conflicting vote.")
		}
	}

	// val1 votes on another height
	{
		pubKey, err := privValidators[1].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 1)
		added, err := signAddVote(privValidators[1], withHeight(vote, height+1), voteSet)
		if added || err == nil {
			t.Errorf("expected VoteSet.Add to fail, wrong height")
		}
	}

	// val2 votes on another round
	{
		pubKey, err := privValidators[2].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 2)
		added, err := signAddVote(privValidators[2], withRound(vote, round+1), voteSet)
		if added || err == nil {
			t.Errorf("expected VoteSet.Add to fail, wrong round")
		}
	}

	// val3 votes of another type.
	{
		pubKey, err := privValidators[3].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 3)
		added, err := signAddVote(privValidators[3], withType(vote, byte(tmproto.PrecommitType)), voteSet)
		if added || err == nil {
			t.Errorf("expected VoteSet.Add to fail, wrong type")
		}
	}
}

func TestVoteSet_List(t *testing.T) {
	height, round := int64(1), int32(0)
	voteSet, _, _, privVals := randVoteSet(height, round, tmproto.PrevoteType, 10, 1)

	votes := addVoteByAllVoterSet(t, voteSet, privVals, height, round)
	assert.Equal(t, votes, voteSet.List())

	voteSet = nil
	assert.Nil(t, voteSet.List())
}

func TestVoteSet_HasAll(t *testing.T) {
	height, round := int64(1), int32(0)
	voteSet, _, _, privVals := randVoteSet(height, round, tmproto.PrevoteType, 10, 1)
	assert.False(t, voteSet.HasAll())

	addVoteByAllVoterSet(t, voteSet, privVals, height, round)
	assert.True(t, voteSet.HasAll())

	voteSet = nil
	assert.False(t, voteSet.HasAll())
}

func TestVoteSet_2_3Majority(t *testing.T) {
	height, round := int64(1), int32(0)
	voteSet, _, _, privValidators := randVoteSet(height, round, tmproto.PrevoteType, 10, 1)

	voteProto := &Vote{
		ValidatorAddress: nil, // NOTE: must fill in
		ValidatorIndex:   -1,  // NOTE: must fill in
		Height:           height,
		Round:            round,
		Type:             tmproto.PrevoteType,
		Timestamp:        tmtime.Now(),
		BlockID:          BlockID{nil, PartSetHeader{}},
	}
	// 6 out of 10 voted for nil.
	for i := int32(0); i < 6; i++ {
		pubKey, err := privValidators[i].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, i)
		_, err = signAddVote(privValidators[i], vote, voteSet)
		require.NoError(t, err)
	}
	blockID, ok := voteSet.TwoThirdsMajority()
	assert.False(t, ok || !blockID.IsZero(), "there should be no 2/3 majority")

	// 7th validator voted for some blockhash
	{
		pubKey, err := privValidators[6].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 6)
		_, err = signAddVote(privValidators[6], withBlockHash(vote, tmrand.Bytes(32)), voteSet)
		require.NoError(t, err)
		blockID, ok = voteSet.TwoThirdsMajority()
		assert.False(t, ok || !blockID.IsZero(), "there should be no 2/3 majority")
	}

	// 8th validator voted for nil.
	{
		pubKey, err := privValidators[7].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 7)
		_, err = signAddVote(privValidators[7], vote, voteSet)
		require.NoError(t, err)
		blockID, ok = voteSet.TwoThirdsMajority()
		assert.True(t, ok || blockID.IsZero(), "there should be 2/3 majority for nil")
	}
}

func TestVoteSet_2_3MajorityRedux(t *testing.T) {
	height, round := int64(1), int32(0)
	voteSet, _, _, privValidators := randVoteSet(height, round, tmproto.PrevoteType, 100, 1)

	blockHash := crypto.CRandBytes(32)
	blockPartsTotal := uint32(123)
	blockPartSetHeader := PartSetHeader{blockPartsTotal, crypto.CRandBytes(32)}

	voteProto := &Vote{
		ValidatorAddress: nil, // NOTE: must fill in
		ValidatorIndex:   -1,  // NOTE: must fill in
		Height:           height,
		Round:            round,
		Timestamp:        tmtime.Now(),
		Type:             tmproto.PrevoteType,
		BlockID:          BlockID{blockHash, blockPartSetHeader},
	}

	// 66 out of 100 voted for nil.
	for i := int32(0); i < 66; i++ {
		pubKey, err := privValidators[i].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, i)
		_, err = signAddVote(privValidators[i], vote, voteSet)
		require.NoError(t, err)
	}
	blockID, ok := voteSet.TwoThirdsMajority()
	assert.False(t, ok || !blockID.IsZero(),
		"there should be no 2/3 majority")

	// 67th validator voted for nil
	{
		pubKey, err := privValidators[66].GetPubKey()
		require.NoError(t, err)
		adrr := pubKey.Address()
		vote := withValidator(voteProto, adrr, 66)
		_, err = signAddVote(privValidators[66], withBlockHash(vote, nil), voteSet)
		require.NoError(t, err)
		blockID, ok = voteSet.TwoThirdsMajority()
		assert.False(t, ok || !blockID.IsZero(),
			"there should be no 2/3 majority: last vote added was nil")
	}

	// 68th validator voted for a different BlockParts PartSetHeader
	{
		pubKey, err := privValidators[67].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 67)
		blockPartsHeader := PartSetHeader{blockPartsTotal, crypto.CRandBytes(32)}
		_, err = signAddVote(privValidators[67], withBlockPartSetHeader(vote, blockPartsHeader), voteSet)
		require.NoError(t, err)
		blockID, ok = voteSet.TwoThirdsMajority()
		assert.False(t, ok || !blockID.IsZero(),
			"there should be no 2/3 majority: last vote added had different PartSetHeader Hash")
	}

	// 69th validator voted for different BlockParts Total
	{
		pubKey, err := privValidators[68].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 68)
		blockPartsHeader := PartSetHeader{blockPartsTotal + 1, blockPartSetHeader.Hash}
		_, err = signAddVote(privValidators[68], withBlockPartSetHeader(vote, blockPartsHeader), voteSet)
		require.NoError(t, err)
		blockID, ok = voteSet.TwoThirdsMajority()
		assert.False(t, ok || !blockID.IsZero(),
			"there should be no 2/3 majority: last vote added had different PartSetHeader Total")
	}

	// 70th validator voted for different BlockHash
	{
		pubKey, err := privValidators[69].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 69)
		_, err = signAddVote(privValidators[69], withBlockHash(vote, tmrand.Bytes(32)), voteSet)
		require.NoError(t, err)
		blockID, ok = voteSet.TwoThirdsMajority()
		assert.False(t, ok || !blockID.IsZero(),
			"there should be no 2/3 majority: last vote added had different BlockHash")
	}

	// 71st validator voted for the right BlockHash & BlockPartSetHeader
	{
		pubKey, err := privValidators[70].GetPubKey()
		require.NoError(t, err)
		addr := pubKey.Address()
		vote := withValidator(voteProto, addr, 70)
		_, err = signAddVote(privValidators[70], vote, voteSet)
		require.NoError(t, err)
		blockID, ok = voteSet.TwoThirdsMajority()
		assert.True(t, ok && blockID.Equals(BlockID{blockHash, blockPartSetHeader}),
			"there should be 2/3 majority")
	}
}

func TestVoteSet_Conflicts(t *testing.T) {
	height, round := int64(1), int32(0)
	voteSet, _, _, privValidators := randVoteSet(height, round, tmproto.PrevoteType, 4, 1)
	blockHash1 := tmrand.Bytes(32)
	blockHash2 := tmrand.Bytes(32)

	voteProto := &Vote{
		ValidatorAddress: nil,
		ValidatorIndex:   -1,
		Height:           height,
		Round:            round,
		Timestamp:        tmtime.Now(),
		Type:             tmproto.PrevoteType,
		BlockID:          BlockID{nil, PartSetHeader{}},
	}

	val0, err := privValidators[0].GetPubKey()
	require.NoError(t, err)
	val0Addr := val0.Address()

	// val0 votes for nil.
	{
		vote := withValidator(voteProto, val0Addr, 0)
		added, err := signAddVote(privValidators[0], vote, voteSet)
		if !added || err != nil {
			t.Errorf("expected VoteSet.Add to succeed")
		}
	}

	// val0 votes again for blockHash1.
	{
		vote := withValidator(voteProto, val0Addr, 0)
		added, err := signAddVote(privValidators[0], withBlockHash(vote, blockHash1), voteSet)
		assert.False(t, added, "conflicting vote")
		assert.Error(t, err, "conflicting vote")
	}

	// start tracking blockHash1
	err = voteSet.SetPeerMaj23("peerA", BlockID{blockHash1, PartSetHeader{}})
	require.NoError(t, err)

	// val0 votes again for blockHash1.
	{
		vote := withValidator(voteProto, val0Addr, 0)
		added, err := signAddVote(privValidators[0], withBlockHash(vote, blockHash1), voteSet)
		assert.True(t, added, "called SetPeerMaj23()")
		assert.Error(t, err, "conflicting vote")
	}

	// attempt tracking blockHash2, should fail because already set for peerA.
	err = voteSet.SetPeerMaj23("peerA", BlockID{blockHash2, PartSetHeader{}})
	require.Error(t, err)

	// val0 votes again for blockHash1.
	{
		vote := withValidator(voteProto, val0Addr, 0)
		added, err := signAddVote(privValidators[0], withBlockHash(vote, blockHash2), voteSet)
		assert.False(t, added, "duplicate SetPeerMaj23() from peerA")
		assert.Error(t, err, "conflicting vote")
	}

	// val1 votes for blockHash1.
	{
		pv, err := privValidators[1].GetPubKey()
		assert.NoError(t, err)
		addr := pv.Address()
		vote := withValidator(voteProto, addr, 1)
		added, err := signAddVote(privValidators[1], withBlockHash(vote, blockHash1), voteSet)
		if !added || err != nil {
			t.Errorf("expected VoteSet.Add to succeed")
		}
	}

	// check
	if voteSet.HasTwoThirdsMajority() {
		t.Errorf("we shouldn't have 2/3 majority yet")
	}
	if voteSet.HasTwoThirdsAny() {
		t.Errorf("we shouldn't have 2/3 if any votes yet")
	}

	// val2 votes for blockHash2.
	{
		pv, err := privValidators[2].GetPubKey()
		assert.NoError(t, err)
		addr := pv.Address()
		vote := withValidator(voteProto, addr, 2)
		added, err := signAddVote(privValidators[2], withBlockHash(vote, blockHash2), voteSet)
		if !added || err != nil {
			t.Errorf("expected VoteSet.Add to succeed")
		}
	}

	// check
	if voteSet.HasTwoThirdsMajority() {
		t.Errorf("we shouldn't have 2/3 majority yet")
	}
	if !voteSet.HasTwoThirdsAny() {
		t.Errorf("we should have 2/3 if any votes")
	}

	// now attempt tracking blockHash1
	err = voteSet.SetPeerMaj23("peerB", BlockID{blockHash1, PartSetHeader{}})
	require.NoError(t, err)

	// val2 votes for blockHash1.
	{
		pv, err := privValidators[2].GetPubKey()
		assert.NoError(t, err)
		addr := pv.Address()
		vote := withValidator(voteProto, addr, 2)
		added, err := signAddVote(privValidators[2], withBlockHash(vote, blockHash1), voteSet)
		assert.True(t, added)
		assert.Error(t, err, "conflicting vote")
	}

	// check
	if !voteSet.HasTwoThirdsMajority() {
		t.Errorf("we should have 2/3 majority for blockHash1")
	}
	blockIDMaj23, _ := voteSet.TwoThirdsMajority()
	if !bytes.Equal(blockIDMaj23.Hash, blockHash1) {
		t.Errorf("got the wrong 2/3 majority blockhash")
	}
	if !voteSet.HasTwoThirdsAny() {
		t.Errorf("we should have 2/3 if any votes")
	}
}

func TestVoteSet_MakeCommit(t *testing.T) {
	height, round := int64(1), int32(0)
	voteSet, _, _, privValidators := randVoteSet(height, round, tmproto.PrecommitType, 10, 1)
	blockHash, blockPartSetHeader := crypto.CRandBytes(32), PartSetHeader{123, crypto.CRandBytes(32)}

	voteProto := &Vote{
		ValidatorAddress: nil,
		ValidatorIndex:   -1,
		Height:           height,
		Round:            round,
		Timestamp:        tmtime.Now(),
		Type:             tmproto.PrecommitType,
		BlockID:          BlockID{blockHash, blockPartSetHeader},
	}

	// 6 out of 10 voted for some block.
	for i := int32(0); i < 6; i++ {
		pv, err := privValidators[i].GetPubKey()
		assert.NoError(t, err)
		addr := pv.Address()
		vote := withValidator(voteProto, addr, i)
		_, err = signAddVote(privValidators[i], vote, voteSet)
		if err != nil {
			t.Error(err)
		}
	}

	// MakeCommit should fail.
	assert.Panics(t, func() { voteSet.MakeCommit() }, "Doesn't have +2/3 majority")

	// 7th voted for some other block.
	{
		pv, err := privValidators[6].GetPubKey()
		assert.NoError(t, err)
		addr := pv.Address()
		vote := withValidator(voteProto, addr, 6)
		vote = withBlockHash(vote, tmrand.Bytes(32))
		vote = withBlockPartSetHeader(vote, PartSetHeader{123, tmrand.Bytes(32)})

		_, err = signAddVote(privValidators[6], vote, voteSet)
		require.NoError(t, err)
	}

	// The 8th voted like everyone else.
	{
		pv, err := privValidators[7].GetPubKey()
		assert.NoError(t, err)
		addr := pv.Address()
		vote := withValidator(voteProto, addr, 7)
		_, err = signAddVote(privValidators[7], vote, voteSet)
		require.NoError(t, err)
	}

	// The 9th voted for nil.
	{
		pv, err := privValidators[8].GetPubKey()
		assert.NoError(t, err)
		addr := pv.Address()
		vote := withValidator(voteProto, addr, 8)
		vote.BlockID = BlockID{}

		_, err = signAddVote(privValidators[8], vote, voteSet)
		require.NoError(t, err)
	}

	commit := voteSet.MakeCommit()

	// Commit should have 10 elements
	assert.Equal(t, 10, len(commit.Signatures))

	// Ensure that Commit is good.
	if err := commit.ValidateBasic(); err != nil {
		t.Errorf("error in Commit.ValidateBasic(): %v", err)
	}

	// Signature aggregation
	t.Run("SignatureAggregation", func(t *testing.T) {
		privKeys := [...]crypto.PrivKey{
			bls.GenPrivKey(),
			composite.GenPrivKey(),
			ed25519.GenPrivKey(),
			bls.GenPrivKey(),
		}
		voteSet, _, _, privValidators := randVoteSetForPrivKeys(height, round, tmproto.PrecommitType, privKeys[:], 1)
		for i := range privKeys {
			pubKey, err := privValidators[i].GetPubKey()
			if err != nil {
				t.Fatal(err)
			}
			addr := pubKey.Address()
			vote := withValidator(voteProto, addr, int32(i))
			fmt.Printf("*** %v - %v\n", vote.ValidatorAddress, addr)
			if _, err := signAddVote(privValidators[i], vote, voteSet); err != nil {
				t.Error(err)
			}
		}

		commit := voteSet.MakeCommit()

		assert.Equal(t, height, commit.Height)
		assert.Equal(t, round, commit.Round)
		assert.NotNil(t, commit.AggregatedSignature)
		// The order of commit.Signatures is sorted by address.
		for i := range commit.Signatures {
			idx := 0
			for {
				if bytes.Equal(privKeys[idx].PubKey().Address(), commit.Signatures[i].ValidatorAddress) {
					break
				}
				idx++
			}
			if _, ok := privKeys[idx].(ed25519.PrivKey); ok {
				assert.NotNil(t, commit.Signatures[i].Signature)
			} else if _, ok := privKeys[idx].(bls.PrivKey); ok {
				assert.Nil(t, commit.Signatures[i].Signature)
			} else if _, ok := privKeys[idx].(composite.PrivKey); ok {
				assert.Nil(t, commit.Signatures[i].Signature)
			}
		}
	})
}
