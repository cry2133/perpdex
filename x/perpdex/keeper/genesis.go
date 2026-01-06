package keeper

import (
	"context"

	"perpdex/x/perpdex/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.LoanList {
		if err := k.Loan.Set(ctx, elem.Id, elem); err != nil {
			return err
		}
	}

	if err := k.LoanSeq.Set(ctx, genState.LoanCount); err != nil {
		return err
	}
	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	err = k.Loan.Walk(ctx, nil, func(key uint64, elem types.Loan) (bool, error) {
		genesis.LoanList = append(genesis.LoanList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.LoanCount, err = k.LoanSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}
