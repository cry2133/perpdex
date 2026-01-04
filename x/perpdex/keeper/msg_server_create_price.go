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

	var post = types.Price{
		Creator: msg.Creator,
		Symbol:  msg.Symbol,
		Price:   msg.Price,
	}
	id := k.AppendPrice(ctx, post)

	return &types.MsgCreatePriceResponse{
		Id: id,
	}, nil
}
