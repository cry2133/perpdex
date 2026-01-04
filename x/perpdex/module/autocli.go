package perpdex

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"perpdex/x/perpdex/types"
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
					RpcMethod:      "ShowPrice",
					Use:            "show-price [id]",
					Short:          "Query show-price",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowSymbolPrice",
					Use:            "show-symbol-price [symbol]",
					Short:          "Query show-symbol-price",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "symbol"}},
				},

				{
					RpcMethod:      "ListPrice",
					Use:            "list-price ",
					Short:          "Query list-price",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod:      "CreatePrice",
					Use:            "create-price [symbol] [price]",
					Short:          "Send a create-price tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "symbol"}, {ProtoField: "price"}},
				},
				{
					RpcMethod:      "UpdatePrice",
					Use:            "update-price [symbol] [price] [id]",
					Short:          "Send a update-price tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "symbol"}, {ProtoField: "price"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeletePrice",
					Use:            "delete-price [id]",
					Short:          "Send a delete-price tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
