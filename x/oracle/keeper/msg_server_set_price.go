package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/cry2133/perpdex/x/oracle/types"
)

func (k msgServer) SetPrice(ctx context.Context, msg *types.MsgSetPrice) (*types.MsgSetPriceResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgSetPriceResponse{}, nil
}
