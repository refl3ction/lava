package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/pairing/types"
	subsciption "github.com/lavanet/lava/x/subscription/keeper"
	subsciptiontypes "github.com/lavanet/lava/x/subscription/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProviderMonthlyPayout(goCtx context.Context, req *types.QueryProviderMonthlyPayoutRequest) (*types.QueryProviderMonthlyPayoutResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	providerAddr, err := sdk.AccAddressFromBech32(req.Provider)
	if err != nil {
		return nil, utils.LavaFormatError("invalid provider address", err,
			utils.Attribute{Key: "provider", Value: req.Provider},
		)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var total uint64

	subs := k.subscriptionKeeper.GetAllSubscriptionsIndices(ctx)

	var details []*types.SubscriptionPayout
	for _, sub := range subs {
		trackedCuInds := k.subscriptionKeeper.GetAllSubTrackedCuIndices(ctx, subsciptiontypes.CuTrackerKey(sub, req.Provider, ""))

		for _, trackCU := range trackedCuInds {
			_, _, chainID := subsciptiontypes.DecodeCuTrackerKey(trackCU)

			subObj, found := k.subscriptionKeeper.GetSubscription(ctx, sub)
			if !found {
				return nil, utils.LavaFormatError("cannot get tracked CU", fmt.Errorf("subscription not found"),
					utils.Attribute{Key: "sub", Value: sub},
					utils.Attribute{Key: "provider", Value: req.Provider},
					utils.Attribute{Key: "chain_id", Value: chainID},
				)
			}

			plan, err := k.subscriptionKeeper.GetPlanFromSubscription(ctx, sub, subObj.Block)
			if err != nil {
				return nil, err
			}

			providerCu, found, _ := k.subscriptionKeeper.GetTrackedCu(ctx, sub, req.Provider, chainID, subObj.Block)
			if !found {
				continue
			}
			totalTokenAmount := plan.Price.Amount
			totalCuTracked := subObj.MonthCuTotal - subObj.MonthCuLeft
			// Sanity check - totalCuTracked > 0
			if totalCuTracked <= 0 {
				return nil, utils.LavaFormatWarning("totalCuTracked is zero or negative", fmt.Errorf("critical: Attempt to divide by zero or negative number"),
					utils.LogAttr("subObj.MonthCuTotal", subObj.MonthCuTotal),
					utils.LogAttr("subObj.MonthCuLeft", subObj.MonthCuLeft),
					utils.LogAttr("totalCuTracked", totalCuTracked),
					utils.LogAttr("sub_block", subObj.Block),
				)
			}

			if plan.Price.Amount.Quo(sdk.NewIntFromUint64(totalCuTracked)).GT(sdk.NewIntFromUint64(subsciption.LIMIT_TOKEN_PER_CU)) {
				totalTokenAmount = sdk.NewIntFromUint64(subsciption.LIMIT_TOKEN_PER_CU * totalCuTracked)
			}

			totalMonthlyReward := k.subscriptionKeeper.CalcTotalMonthlyReward(ctx, totalTokenAmount, providerCu, totalCuTracked)

			// calculate only the provider reward
			providerReward, _, err := k.dualstakingKeeper.RewardProvidersAndDelegators(ctx, providerAddr, chainID, sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), totalMonthlyReward)), subsciptiontypes.ModuleName, true, true, true)
			if err != nil {
				return nil, err
			}

			details = append(details, &types.SubscriptionPayout{
				Subscription: sub,
				ChainId:      chainID,
				Amount:       providerReward.AmountOf(k.stakingKeeper.BondDenom(ctx)).Uint64(),
			})
			total += providerReward.AmountOf(k.stakingKeeper.BondDenom(ctx)).Uint64()
		}
	}

	return &types.QueryProviderMonthlyPayoutResponse{Total: total, Details: details}, nil
}
