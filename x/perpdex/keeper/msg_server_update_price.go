package keeper

import (
	"context"
	"fmt"

	"perpdex/x/perpdex/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdatePrice(ctx context.Context, msg *types.MsgUpdatePrice) (*types.MsgUpdatePriceResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	var post = types.Price{
		Creator: msg.Creator,
		Id:      msg.Id,
		Symbol:  msg.Symbol,
		Price:   msg.Price,
	}
	val, found := k.GetPrice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetPrice(ctx, post)
	return &types.MsgUpdatePriceResponse{}, nil
}
