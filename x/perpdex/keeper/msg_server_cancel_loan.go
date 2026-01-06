package keeper

import (
	"context"

	"perpdex/x/perpdex/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) CancelLoan(ctx context.Context, msg *types.MsgCancelLoan) (*types.MsgCancelLoanResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgCancelLoanResponse{}, nil
}
