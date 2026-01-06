package keeper

import (
	"context"
	"strconv"

	"perpdex/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) LiquidateLoan(ctx context.Context, msg *types.MsgLiquidateLoan) (*types.MsgLiquidateLoanResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	loan, found := k.GetLoan(sdkCtx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.Lender != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot liquidate: not the lender")
	}
	if loan.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	deadline, err := strconv.ParseInt(loan.Deadline, 10, 64)
	if err != nil {
		panic(err)
	}
	if sdkCtx.BlockHeight() < deadline {
		return nil, errorsmod.Wrap(types.ErrDeadline, "Cannot liquidate before deadline")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "liquidated"
	k.SetLoan(sdkCtx, loan)
	return &types.MsgLiquidateLoanResponse{}, nil
}
