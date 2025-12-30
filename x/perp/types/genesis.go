package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:      DefaultParams(),
		PositionMap: []Position{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	positionIndexMap := make(map[string]struct{})

	for _, elem := range gs.PositionMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := positionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for position")
		}
		positionIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
