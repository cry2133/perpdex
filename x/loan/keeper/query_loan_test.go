package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"perpdex/x/loan/keeper"
	"perpdex/x/loan/types"
)

func createNLoan(keeper keeper.Keeper, ctx context.Context, n int) []types.Loan {
	items := make([]types.Loan, n)
	for i := range items {
		iu := uint64(i)
		items[i].Id = iu
		items[i].Amount = strconv.Itoa(i)
		items[i].Fee = strconv.Itoa(i)
		items[i].Collateral = strconv.Itoa(i)
		items[i].Deadline = strconv.Itoa(i)
		items[i].State = strconv.Itoa(i)
		items[i].Borrower = strconv.Itoa(i)
		items[i].Lender = strconv.Itoa(i)
		_ = keeper.Loan.Set(ctx, iu, items[i])
		_ = keeper.LoanSeq.Set(ctx, iu)
	}
	return items
}

func TestLoanQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNLoan(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetLoanRequest
		response *types.QueryGetLoanResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLoanRequest{Id: msgs[0].Id},
			response: &types.QueryGetLoanResponse{Loan: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetLoanRequest{Id: msgs[1].Id},
			response: &types.QueryGetLoanResponse{Loan: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetLoanRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetLoan(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestLoanQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNLoan(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLoanRequest {
		return &types.QueryAllLoanRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListLoan(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Loan), step)
			require.Subset(t, msgs, resp.Loan)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListLoan(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Loan), step)
			require.Subset(t, msgs, resp.Loan)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListLoan(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Loan)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListLoan(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
