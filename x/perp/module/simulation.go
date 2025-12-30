package perp

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	perpsimulation "github.com/cry2133/perpdex/x/perp/simulation"
	"github.com/cry2133/perpdex/x/perp/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	perpGenesis := types.GenesisState{
		Params: types.DefaultParams(),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&perpGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgOpenPosition          = "op_weight_msg_perp"
		defaultWeightMsgOpenPosition int = 100
	)

	var weightMsgOpenPosition int
	simState.AppParams.GetOrGenerate(opWeightMsgOpenPosition, &weightMsgOpenPosition, nil,
		func(_ *rand.Rand) {
			weightMsgOpenPosition = defaultWeightMsgOpenPosition
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOpenPosition,
		perpsimulation.SimulateMsgOpenPosition(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgClosePosition          = "op_weight_msg_perp"
		defaultWeightMsgClosePosition int = 100
	)

	var weightMsgClosePosition int
	simState.AppParams.GetOrGenerate(opWeightMsgClosePosition, &weightMsgClosePosition, nil,
		func(_ *rand.Rand) {
			weightMsgClosePosition = defaultWeightMsgClosePosition
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClosePosition,
		perpsimulation.SimulateMsgClosePosition(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
