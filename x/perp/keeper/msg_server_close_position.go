package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/cry2133/perpdex/x/perp/types"
)

func (k msgServer) ClosePosition(ctx context.Context, msg *types.MsgClosePosition) (*types.MsgClosePositionResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgClosePositionResponse{}, nil
}
