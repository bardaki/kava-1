package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kava-labs/kava/x/liquidstaking/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the liquidstaking MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) MintDerivative(goCtx context.Context, msg *types.MsgMintDerivative) (*types.MsgMintDerivativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validator, err := sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}

	err = k.keeper.MintDerivative(ctx, validator, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Validator),
		),
	)
	return &types.MsgMintDerivativeResponse{}, nil
}

func (k msgServer) BurnDerivative(goCtx context.Context, msg *types.MsgBurnDerivative) (*types.MsgBurnDerivativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	validator, err := sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}

	err = k.keeper.BurnDerivative(ctx, validator, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Validator),
		),
	)
	return &types.MsgBurnDerivativeResponse{}, nil
}