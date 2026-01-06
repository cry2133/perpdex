package perpdex

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	perpdexsimulation "perpdex/x/perpdex/simulation"
	"perpdex/x/perpdex/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	perpdexGenesis := types.GenesisState{
		Params: types.DefaultParams(),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&perpdexGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreatePrice          = "op_weight_msg_perpdex"
		defaultWeightMsgCreatePrice int = 100
	)

	var weightMsgCreatePrice int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePrice, &weightMsgCreatePrice, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePrice = defaultWeightMsgCreatePrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePrice,
		perpdexsimulation.SimulateMsgCreatePrice(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdatePrice          = "op_weight_msg_perpdex"
		defaultWeightMsgUpdatePrice int = 100
	)

	var weightMsgUpdatePrice int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePrice, &weightMsgUpdatePrice, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePrice = defaultWeightMsgUpdatePrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePrice,
		perpdexsimulation.SimulateMsgUpdatePrice(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeletePrice          = "op_weight_msg_perpdex"
		defaultWeightMsgDeletePrice int = 100
	)

	var weightMsgDeletePrice int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePrice, &weightMsgDeletePrice, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePrice = defaultWeightMsgDeletePrice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePrice,
		perpdexsimulation.SimulateMsgDeletePrice(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgRequestLoan          = "op_weight_msg_perpdex"
		defaultWeightMsgRequestLoan int = 100
	)

	var weightMsgRequestLoan int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestLoan, &weightMsgRequestLoan, nil,
		func(_ *rand.Rand) {
			weightMsgRequestLoan = defaultWeightMsgRequestLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestLoan,
		perpdexsimulation.SimulateMsgRequestLoan(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgApproveLoan          = "op_weight_msg_perpdex"
		defaultWeightMsgApproveLoan int = 100
	)

	var weightMsgApproveLoan int
	simState.AppParams.GetOrGenerate(opWeightMsgApproveLoan, &weightMsgApproveLoan, nil,
		func(_ *rand.Rand) {
			weightMsgApproveLoan = defaultWeightMsgApproveLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveLoan,
		perpdexsimulation.SimulateMsgApproveLoan(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCancelLoan          = "op_weight_msg_perpdex"
		defaultWeightMsgCancelLoan int = 100
	)

	var weightMsgCancelLoan int
	simState.AppParams.GetOrGenerate(opWeightMsgCancelLoan, &weightMsgCancelLoan, nil,
		func(_ *rand.Rand) {
			weightMsgCancelLoan = defaultWeightMsgCancelLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelLoan,
		perpdexsimulation.SimulateMsgCancelLoan(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgRepayLoan          = "op_weight_msg_perpdex"
		defaultWeightMsgRepayLoan int = 100
	)

	var weightMsgRepayLoan int
	simState.AppParams.GetOrGenerate(opWeightMsgRepayLoan, &weightMsgRepayLoan, nil,
		func(_ *rand.Rand) {
			weightMsgRepayLoan = defaultWeightMsgRepayLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayLoan,
		perpdexsimulation.SimulateMsgRepayLoan(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgLiquidateLoan          = "op_weight_msg_perpdex"
		defaultWeightMsgLiquidateLoan int = 100
	)

	var weightMsgLiquidateLoan int
	simState.AppParams.GetOrGenerate(opWeightMsgLiquidateLoan, &weightMsgLiquidateLoan, nil,
		func(_ *rand.Rand) {
			weightMsgLiquidateLoan = defaultWeightMsgLiquidateLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiquidateLoan,
		perpdexsimulation.SimulateMsgLiquidateLoan(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
