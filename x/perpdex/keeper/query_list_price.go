package keeper

import (
	"context"

	"perpdex/x/perpdex/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListPrice(ctx context.Context, req *types.QueryListPriceRequest) (*types.QueryListPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryListPriceResponse{}, nil
}
