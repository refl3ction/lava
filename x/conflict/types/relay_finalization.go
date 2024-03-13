package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils/sigs"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
)

func NewRelayFinalization(relaySession *pairingtypes.RelaySession, relayReply *pairingtypes.RelayReply, consumerAddr sdk.AccAddress, blockDistanceToFinalization int64) RelayFinalization {
	return RelayFinalization{
		FinalizedBlocksHashes:       relayReply.FinalizedBlocksHashes,
		LatestBlock:                 relayReply.LatestBlock,
		Sig:                         relayReply.SigBlocks,
		ConsumerAddress:             string(consumerAddr.Bytes()),
		BlockDistanceToFinalization: blockDistanceToFinalization,
		SpecId:                      relaySession.SpecId,
		Epoch:                       relaySession.Epoch,
	}
}

func (rf RelayFinalization) GetSignature() []byte {
	return rf.Sig
}

func (rf RelayFinalization) DataToSign() []byte {
	latestBlockBytes := sigs.EncodeUint64(uint64(rf.LatestBlock))
	blockDistanceToFinalizationBytes := sigs.EncodeUint64(uint64(rf.BlockDistanceToFinalization))
	epochBytes := sigs.EncodeUint64(uint64(rf.Epoch))
	msgParts := [][]byte{
		latestBlockBytes,
		rf.FinalizedBlocksHashes,
		[]byte(rf.ConsumerAddress),
		blockDistanceToFinalizationBytes,
		[]byte(rf.SpecId),
		epochBytes,
	}
	return sigs.Join(msgParts)
}

func (rf RelayFinalization) HashRounds() int {
	return 1
}
