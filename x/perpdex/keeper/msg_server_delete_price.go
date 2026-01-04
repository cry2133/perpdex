package keeper

import (
	"context"

	"perpdex/x/perpdex/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) DeletePrice(ctx context.Context, msg *types.MsgDeletePrice) (*types.MsgDeletePriceResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgDeletePriceResponse{}, nil
}
