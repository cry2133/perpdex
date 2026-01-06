package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"

	"perpdex/x/perpdex/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListPrice(ctx context.Context, req *types.QueryListPriceRequest) (*types.QueryListPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	storeAdapter := runtime.KVStoreAdapter(q.k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.PriceKey))

	var posts []types.Price
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var post types.Price
		if err := q.k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}

		posts = append(posts, post)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListPriceResponse{Post: posts, Pagination: pageRes}, nil
}
