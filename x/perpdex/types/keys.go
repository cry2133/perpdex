package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "perpdex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// GovModuleName duplicates the gov module's name to avoid a dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/v0.52.0-beta.2/x/gov/types/keys.go#L9
	GovModuleName = "gov"

	PriceKey = "Price/value/"

	PriceCountKey = "Price/count/"

	// Add new key for symbol index
	PriceSymbolKey = "Price/symbol/"
)

// ParamsKey is the prefix to retrieve all Params
var ParamsKey = collections.NewPrefix("p_perpdex")
