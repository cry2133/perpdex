import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateParams } from "./types/perpdex/perpdex/v1/tx";
import { MsgCreatePrice } from "./types/perpdex/perpdex/v1/tx";
import { MsgUpdatePrice } from "./types/perpdex/perpdex/v1/tx";
import { MsgDeletePrice } from "./types/perpdex/perpdex/v1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/perpdex.perpdex.v1.MsgUpdateParams", MsgUpdateParams],
    ["/perpdex.perpdex.v1.MsgCreatePrice", MsgCreatePrice],
    ["/perpdex.perpdex.v1.MsgUpdatePrice", MsgUpdatePrice],
    ["/perpdex.perpdex.v1.MsgDeletePrice", MsgDeletePrice],
    
];

export { msgTypes }