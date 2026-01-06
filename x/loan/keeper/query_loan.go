package keeper

import (
	"context"
	"errors"

	"perpdex/x/loan/types"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListLoan(ctx context.Context, req *types.QueryAllLoanRequest) (*types.QueryAllLoanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	loans, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Loan,
		req.Pagination,
		func(_ uint64, value types.Loan) (types.Loan, error) {
			return value, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLoanResponse{Loan: loans, Pagination: pageRes}, nil
}

func (q queryServer) GetLoan(ctx context.Context, req *types.QueryGetLoanRequest) (*types.QueryGetLoanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	loan, err := q.k.Loan.Get(ctx, req.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetLoanResponse{Loan: loan}, nil
}
