package keeper

import (
	"context"

	"perpdex/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelLoan(ctx context.Context, msg *types.MsgCancelLoan) (*types.MsgCancelLoanResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	loan, found := k.GetLoan(sdk.UnwrapSDKContext(ctx), msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.Borrower != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not the borrower")
	}
	if loan.State != "requested" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)

	}
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "cancelled"
	k.SetLoan(sdk.UnwrapSDKContext(ctx), loan)
	return &types.MsgCancelLoanResponse{}, nil
}
