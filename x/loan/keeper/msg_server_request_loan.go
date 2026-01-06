package keeper

import (
	"context"
	"strconv"

	"perpdex/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RequestLoan(ctx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	amount, _ := sdk.ParseCoinsNormalized(msg.Amount)
	if !amount.IsValid() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
	}
	if amount.Empty() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
	}
	fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	if !fee.IsValid() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "fee is not a valid Coins object")
	}
	deadline, err := strconv.ParseInt(msg.Deadline, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deadline is not an integer")
	}
	if deadline <= 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deadline should be a positive integer")
	}
	collateral, _ := sdk.ParseCoinsNormalized(msg.Collateral)
	if !collateral.IsValid() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collateral is not a valid Coins object")
	}
	if collateral.Empty() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collateral is empty")
	}

	borrower, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	var loan = types.Loan{
		Amount:     msg.Amount,
		Fee:        msg.Fee,
		Collateral: msg.Collateral,
		Deadline:   msg.Deadline,
		State:      "requested",
		Borrower:   msg.Creator,
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}
	k.AppendLoan(sdk.UnwrapSDKContext(ctx), loan)
	return &types.MsgRequestLoanResponse{}, nil
}
