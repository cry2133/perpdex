package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/loan module sentinel errors
var (
	ErrInvalidSigner  = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrDeadline       = errors.Register(ModuleName, 3, "deadline")
	ErrWrongLoanState = errors.Register(ModuleName, 2, "wrong loan state")
)
