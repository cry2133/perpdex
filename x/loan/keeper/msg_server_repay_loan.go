package keeper

import (
	"context"

	"perpdex/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RepayLoan(ctx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	loan, found := k.GetLoan(sdk.UnwrapSDKContext(ctx), msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	if msg.Creator != loan.Borrower {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot repay: not the borrower")
	}
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)
	fee, _ := sdk.ParseCoinsNormalized(loan.Fee)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	err := k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, fee)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "repayed"
	k.SetLoan(sdk.UnwrapSDKContext(ctx), loan)
	return &types.MsgRepayLoanResponse{}, nil
}
