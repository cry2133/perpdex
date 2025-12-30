package perp

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"github.com/cry2133/perpdex/x/perp/types"
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
					RpcMethod: "ListPosition",
					Use:       "list-position",
					Short:     "List all position",
				},
				{
					RpcMethod:      "GetPosition",
					Use:            "get-position [id]",
					Short:          "Gets a position",
					Alias:          []string{"show-position"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
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
					RpcMethod:      "OpenPosition",
					Use:            "open-position [pair] [side] [leverage] [margin]",
					Short:          "Send a open-position tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pair"}, {ProtoField: "side"}, {ProtoField: "leverage"}, {ProtoField: "margin"}},
				},
				{
					RpcMethod:      "ClosePosition",
					Use:            "close-position [position-id]",
					Short:          "Send a close-position tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "position_id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
