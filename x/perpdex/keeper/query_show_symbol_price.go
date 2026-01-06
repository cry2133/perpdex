package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"perpdex/x/perpdex/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ShowSymbolPrice(ctx context.Context, req *types.QueryShowSymbolPriceRequest) (*types.QueryShowSymbolPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	price, found := q.k.GetPriceBySymbol(ctx, req.Symbol)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryShowSymbolPriceResponse{Post: price}, nil
}
