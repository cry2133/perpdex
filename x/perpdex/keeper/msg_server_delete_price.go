package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"perpdex/x/perpdex/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) DeletePrice(ctx context.Context, msg *types.MsgDeletePrice) (*types.MsgDeletePriceResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	val, found := k.GetPrice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.RemovePrice(ctx, msg.Id)

	return &types.MsgDeletePriceResponse{}, nil
}
