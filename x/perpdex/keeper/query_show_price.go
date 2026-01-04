package keeper

import (
	"context"

	"perpdex/x/perpdex/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ShowPrice(ctx context.Context, req *types.QueryShowPriceRequest) (*types.QueryShowPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryShowPriceResponse{}, nil
}
