package common

import (
	"context"
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/lavanet/lava/testutil/keeper"
	"github.com/lavanet/lava/utils/sigs"
	conflicttypes "github.com/lavanet/lava/x/conflict/types"
	conflictconstruct "github.com/lavanet/lava/x/conflict/types/construct"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"github.com/lavanet/lava/x/pairing/types"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	subscriptiontypes "github.com/lavanet/lava/x/subscription/types"
	"github.com/stretchr/testify/require"
)

func CreateNewAccount(ctx context.Context, keepers testkeeper.Keepers, balance int64) (acc sigs.Account) {
	acc = sigs.GenerateDeterministicFloatingKey(testkeeper.Randomizer)
	testkeeper.Randomizer.Inc()
	coins := sdk.NewCoins(sdk.NewCoin(keepers.StakingKeeper.BondDenom(sdk.UnwrapSDKContext(ctx)), sdk.NewInt(balance)))
	keepers.BankKeeper.SetBalance(sdk.UnwrapSDKContext(ctx), acc.Addr, coins)
	return
}

func StakeAccount(t *testing.T, ctx context.Context, keepers testkeeper.Keepers, servers testkeeper.Servers, acc sigs.Account, spec spectypes.Spec, stake int64, validator sigs.Account) {
	endpoints := []epochstoragetypes.Endpoint{}
	for _, collection := range spec.ApiCollections {
		endpoints = append(endpoints, epochstoragetypes.Endpoint{IPPORT: "123", ApiInterfaces: []string{collection.CollectionData.ApiInterface}, Geolocation: 1})
	}
	_, err := servers.PairingServer.StakeProvider(ctx, &types.MsgStakeProvider{Creator: acc.Addr.String(), ChainID: spec.Index, Amount: sdk.NewCoin(keepers.StakingKeeper.BondDenom(sdk.UnwrapSDKContext(ctx)), sdk.NewInt(stake)), Geolocation: 1, Endpoints: endpoints, Moniker: "prov", DelegateLimit: sdk.NewCoin(keepers.StakingKeeper.BondDenom(sdk.UnwrapSDKContext(ctx)), sdk.ZeroInt()), DelegateCommission: 100, Validator: sdk.ValAddress(validator.Addr).String()})
	require.NoError(t, err)
}

func BuySubscription(ctx context.Context, keepers testkeeper.Keepers, servers testkeeper.Servers, acc sigs.Account, plan string) {
	servers.SubscriptionServer.Buy(ctx, &subscriptiontypes.MsgBuy{Creator: acc.Addr.String(), Consumer: acc.Addr.String(), Index: plan, Duration: 1})
}

func BuildRelayRequest(ctx context.Context, provider string, contentHash []byte, cuSum uint64, spec string, qos *types.QualityOfServiceReport) *types.RelaySession {
	return BuildRelayRequestWithBadge(ctx, provider, contentHash, uint64(1), cuSum, spec, qos, nil)
}

func BuildRelayRequestWithSession(ctx context.Context, provider string, contentHash []byte, sessionId uint64, cuSum uint64, spec string, qos *types.QualityOfServiceReport) *types.RelaySession {
	return BuildRelayRequestWithBadge(ctx, provider, contentHash, sessionId, cuSum, spec, qos, nil)
}

func BuildRelayRequestWithBadge(ctx context.Context, provider string, contentHash []byte, sessionId uint64, cuSum uint64, spec string, qos *types.QualityOfServiceReport, badge *types.Badge) *types.RelaySession {
	relaySession := &types.RelaySession{
		Provider:    provider,
		ContentHash: contentHash,
		SessionId:   sessionId,
		SpecId:      spec,
		CuSum:       cuSum,
		Epoch:       sdk.UnwrapSDKContext(ctx).BlockHeight(),
		RelayNum:    0,
		QosReport:   qos,
		LavaChainId: sdk.UnwrapSDKContext(ctx).BlockHeader().ChainID,
		Badge:       badge,
	}
	if qos != nil {
		qos.ComputeQoS()
	}
	return relaySession
}

func CreateResponseConflictMsgDetectionForTest(ctx context.Context, consumer, provider0, provider1 sigs.Account, spec spectypes.Spec) (detectionMsg *conflicttypes.MsgDetection, reply1, reply2 *types.RelayReply, errRet error) {
	detectionMsg = &conflicttypes.MsgDetection{
		Creator: consumer.Addr.String(),
		ResponseConflict: &conflicttypes.ResponseConflict{
			ConflictRelayData0: initConflictRelayData(),
			ConflictRelayData1: initConflictRelayData(),
		},
	}

	// Prepare request and session for provider0.
	prepareRelayData(ctx, detectionMsg.ResponseConflict.ConflictRelayData0, provider0, spec)
	// Sign the session data with the consumer's private key.
	if err := signSessionData(consumer, detectionMsg.ResponseConflict.ConflictRelayData0.Request.RelaySession); err != nil {
		return detectionMsg, nil, nil, err
	}

	// Duplicate the request for provider1 and update provider-specific fields.
	duplicateRequestForProvider(detectionMsg.ResponseConflict, provider1, consumer)
	// Sign the session data with the consumer's private key.
	if err := signSessionData(consumer, detectionMsg.ResponseConflict.ConflictRelayData1.Request.RelaySession); err != nil {
		return detectionMsg, nil, nil, err
	}

	// Create and sign replies for both providers.
	reply1, err := createAndSignReply(provider0, detectionMsg.ResponseConflict.ConflictRelayData0.Request, spec, false)
	if err != nil {
		return detectionMsg, nil, nil, err
	}

	reply2, err = createAndSignReply(provider1, detectionMsg.ResponseConflict.ConflictRelayData1.Request, spec, true)
	if err != nil {
		return detectionMsg, nil, nil, err
	}

	// Construct final conflict relay data with the replies.
	conflictRelayData0, err := finalizeConflictRelayData(consumer, provider0, detectionMsg.ResponseConflict.ConflictRelayData0, reply1)
	if err != nil {
		return detectionMsg, nil, nil, err
	}
	conflictRelayData1, err := finalizeConflictRelayData(consumer, provider1, detectionMsg.ResponseConflict.ConflictRelayData1, reply2)
	if err != nil {
		return detectionMsg, nil, nil, err
	}

	detectionMsg.ResponseConflict.ConflictRelayData0 = conflictRelayData0
	detectionMsg.ResponseConflict.ConflictRelayData1 = conflictRelayData1

	return detectionMsg, reply1, reply2, nil
}

// initConflictRelayData initializes the structure for holding relay conflict data.
func initConflictRelayData() *conflicttypes.ConflictRelayData {
	return &conflicttypes.ConflictRelayData{
		Request: &types.RelayRequest{},
		Reply:   &conflicttypes.ReplyMetadata{},
	}
}

// prepareRelayData prepares relay data for a given provider.
func prepareRelayData(ctx context.Context, conflictData *conflicttypes.ConflictRelayData, provider sigs.Account, spec spectypes.Spec) {
	relayData := &types.RelayPrivateData{
		ConnectionType: "",
		ApiUrl:         "",
		Data:           []byte("DUMMYREQUEST"),
		RequestBlock:   100,
		ApiInterface:   "",
		Salt:           []byte{1},
	}

	conflictData.Request.RelayData = relayData
	conflictData.Request.RelaySession = &types.RelaySession{
		Provider:    provider.Addr.String(),
		ContentHash: sigs.HashMsg(relayData.GetContentHashData()),
		SessionId:   uint64(1),
		SpecId:      spec.Index,
		CuSum:       0,
		Epoch:       sdk.UnwrapSDKContext(ctx).BlockHeight(),
		RelayNum:    0,
		QosReport:   &types.QualityOfServiceReport{Latency: sdk.OneDec(), Availability: sdk.OneDec(), Sync: sdk.OneDec()},
	}

}

// signSessionData signs the session data with the consumer's private key.
func signSessionData(consumer sigs.Account, relaySession *pairingtypes.RelaySession) error {
	sig, err := sigs.Sign(consumer.SK, *relaySession)
	if err != nil {
		return err
	}

	relaySession.Sig = sig
	return nil
}

// duplicateRequestForProvider duplicates request data for another provider and signs it.
func duplicateRequestForProvider(conflict *conflicttypes.ResponseConflict, provider, consumer sigs.Account) {
	// Clone request data
	temp, _ := conflict.ConflictRelayData0.Request.Marshal()
	conflict.ConflictRelayData1.Request.Unmarshal(temp)

	conflict.ConflictRelayData1.Request.RelaySession.Provider = provider.Addr.String()
	conflict.ConflictRelayData1.Request.RelaySession.Sig = []byte{}
}

// createAndSignReply creates a reply for a provider and signs it.
func createAndSignReply(provider sigs.Account, request *types.RelayRequest, spec spectypes.Spec, addDiffData bool) (*types.RelayReply, error) {
	reply := &types.RelayReply{
		Data:                  []byte("DUMMYREPLY"),
		Sig:                   request.RelaySession.Sig,
		LatestBlock:           request.RelayData.RequestBlock + int64(spec.BlockDistanceForFinalizedData),
		FinalizedBlocksHashes: []byte{},
		SigBlocks:             request.RelaySession.Sig,
		Metadata:              []types.Metadata{},
	}

	if addDiffData {
		reply.Data = append(reply.Data, []byte("DIFF")...)
	}

	relayExchange := types.NewRelayExchange(*request, *reply)
	sig, err := sigs.Sign(provider.SK, relayExchange)
	if err != nil {
		return reply, err
	}

	reply.Sig = sig
	return reply, nil
}

// finalizeConflictRelayData updates the conflict relay data with the reply information.
func finalizeConflictRelayData(consumer, provider sigs.Account, conflictData *conflicttypes.ConflictRelayData, reply *types.RelayReply) (*conflicttypes.ConflictRelayData, error) {
	relayFinalization := conflicttypes.NewRelayFinalization(conflictData.Request.RelaySession, reply, consumer.Addr, 0)
	sigBlocks, err := sigs.Sign(provider.SK, relayFinalization)
	if err != nil {
		return nil, err
	}
	reply.SigBlocks = sigBlocks
	conflictRelayData := conflictconstruct.ConstructConflictRelayData(reply, conflictData.Request)
	return conflictRelayData, nil
}

func CreateRelayFinalizationForTest(ctx context.Context, consumer, provider sigs.Account, epoch, latestBlock int64, finalizationBlockHashes map[int64]string, spec spectypes.Spec) (*conflicttypes.RelayFinalization, error) {
	relayFinalization := initConflictRelayFinalization(epoch, latestBlock, spec, consumer)

	err := setRelayFinalizationFinalizedBlocksHashes(relayFinalization, finalizationBlockHashes)
	if err != nil {
		return relayFinalization, err
	}

	// Sign relay reply for provider
	sig0, err := sigs.Sign(provider.SK, relayFinalization)
	if err != nil {
		return relayFinalization, err
	}
	relayFinalization.Sig = sig0

	return relayFinalization, nil
}

// initConflictRelayData initializes the structure for holding relay conflict data.
func initConflictRelayFinalization(epoch, latestBlock int64, spec spectypes.Spec, consumer sigs.Account) *conflicttypes.RelayFinalization {
	return &conflicttypes.RelayFinalization{
		FinalizedBlocksHashes:       []byte{},
		LatestBlock:                 latestBlock,
		ConsumerAddress:             consumer.Addr.String(),
		Sig:                         []byte{},
		BlockDistanceToFinalization: int64(spec.BlockDistanceForFinalizedData),
		SpecId:                      spec.Index,
		Epoch:                       epoch,
	}
}

// setRelayFinalizationFinalizedBlocksHashes sets the finalized blocks hashes in the relay finalization
func setRelayFinalizationFinalizedBlocksHashes(relayFinalization *conflicttypes.RelayFinalization, finalizedBlocksHashes map[int64]string) error {
	jsonStr, err := json.Marshal(finalizedBlocksHashes)
	if err != nil {
		return err
	}

	relayFinalization.FinalizedBlocksHashes = []byte(jsonStr)
	return nil
}
