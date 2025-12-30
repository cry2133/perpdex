package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cry2133/perpdex/x/perp/keeper"
	"github.com/cry2133/perpdex/x/perp/types"
)

func createNPosition(keeper keeper.Keeper, ctx context.Context, n int) []types.Position {
	items := make([]types.Position, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Trader = strconv.Itoa(i)
		items[i].Pair = strconv.Itoa(i)
		items[i].Side = strconv.Itoa(i)
		items[i].Size = strconv.Itoa(i)
		items[i].EntryPrice = strconv.Itoa(i)
		items[i].Margin = strconv.Itoa(i)
		items[i].Leverage = strconv.Itoa(i)
		_ = keeper.Position.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestPositionQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNPosition(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPositionRequest
		response *types.QueryGetPositionResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPositionRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetPositionResponse{Position: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPositionRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetPositionResponse{Position: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPositionRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetPosition(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestPositionQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNPosition(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPositionRequest {
		return &types.QueryAllPositionRequest{
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
			resp, err := qs.ListPosition(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Position), step)
			require.Subset(t, msgs, resp.Position)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListPosition(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Position), step)
			require.Subset(t, msgs, resp.Position)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListPosition(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Position)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListPosition(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
