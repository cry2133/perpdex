package keeper

import (
	"context"

	"perpdex/x/loan/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) ApproveLoan(ctx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgApproveLoanResponse{}, nil
}
