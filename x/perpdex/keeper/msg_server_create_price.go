package keeper

import (
	"context"

	"perpdex/x/perpdex/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) CreatePrice(ctx context.Context, msg *types.MsgCreatePrice) (*types.MsgCreatePriceResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgCreatePriceResponse{}, nil
}
