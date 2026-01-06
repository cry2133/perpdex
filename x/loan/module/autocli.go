package loan

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"perpdex/x/loan/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListLoan",
					Use:       "list-loan",
					Short:     "List all loan",
				},
				{
					RpcMethod:      "GetLoan",
					Use:            "get-loan [id]",
					Short:          "Gets a loan by id",
					Alias:          []string{"show-loan"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "RequestLoan",
					Use:            "request-loan [amount] [fee] [collateral] [deadline]",
					Short:          "Send a request-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "fee"}, {ProtoField: "collateral"}, {ProtoField: "deadline"}},
				},
				{
					RpcMethod:      "CancelLoan",
					Use:            "cancel-loan [id]",
					Short:          "Send a cancel-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "ApproveLoan",
					Use:            "approve-loan [id]",
					Short:          "Send a approve-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "RepayLoan",
					Use:            "repay-loan [id]",
					Short:          "Send a repay-loan tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
