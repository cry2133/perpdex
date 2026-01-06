package keeper

import (
	"context"

	"perpdex/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveLoan(ctx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	loan, found := k.GetLoan(sdk.UnwrapSDKContext(ctx), msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.State != "requested" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrWrongLoanState, "Cannot parse coins in loan amount")
	}
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}
	loan.Lender = msg.Creator
	loan.State = "approved"
	k.SetLoan(sdk.UnwrapSDKContext(ctx), loan)
	return &types.MsgApproveLoanResponse{}, nil
}
