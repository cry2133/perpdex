import { DeliverTxResponse, StdFee } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { IgniteClient } from "../client";
import { Api } from "./rest";
import { MsgUpdateParams } from "./types/perpdex/perpdex/v1/tx";
import { MsgCreatePrice } from "./types/perpdex/perpdex/v1/tx";
import { MsgUpdatePrice } from "./types/perpdex/perpdex/v1/tx";
import { MsgDeletePrice } from "./types/perpdex/perpdex/v1/tx";
export { MsgUpdateParams, MsgCreatePrice, MsgUpdatePrice, MsgDeletePrice };
type sendMsgUpdateParamsParams = {
    value: MsgUpdateParams;
    fee?: StdFee;
    memo?: string;
};
type sendMsgCreatePriceParams = {
    value: MsgCreatePrice;
    fee?: StdFee;
    memo?: string;
};
type sendMsgUpdatePriceParams = {
    value: MsgUpdatePrice;
    fee?: StdFee;
    memo?: string;
};
type sendMsgDeletePriceParams = {
    value: MsgDeletePrice;
    fee?: StdFee;
    memo?: string;
};
type msgUpdateParamsParams = {
    value: MsgUpdateParams;
};
type msgCreatePriceParams = {
    value: MsgCreatePrice;
};
type msgUpdatePriceParams = {
    value: MsgUpdatePrice;
};
type msgDeletePriceParams = {
    value: MsgDeletePrice;
};
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
    prefix: string;
    signer?: OfflineSigner;
}
export declare const txClient: ({ signer, prefix, addr }?: TxClientOptions) => {
    sendMsgUpdateParams({ value, fee, memo }: sendMsgUpdateParamsParams): Promise<DeliverTxResponse>;
    sendMsgCreatePrice({ value, fee, memo }: sendMsgCreatePriceParams): Promise<DeliverTxResponse>;
    sendMsgUpdatePrice({ value, fee, memo }: sendMsgUpdatePriceParams): Promise<DeliverTxResponse>;
    sendMsgDeletePrice({ value, fee, memo }: sendMsgDeletePriceParams): Promise<DeliverTxResponse>;
    msgUpdateParams({ value }: msgUpdateParamsParams): EncodeObject;
    msgCreatePrice({ value }: msgCreatePriceParams): EncodeObject;
    msgUpdatePrice({ value }: msgUpdatePriceParams): EncodeObject;
    msgDeletePrice({ value }: msgDeletePriceParams): EncodeObject;
};
interface QueryClientOptions {
    addr: string;
}
export declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Api<unknown>;
declare class SDKModule {
    query: ReturnType<typeof queryClient>;
    tx: ReturnType<typeof txClient>;
    structure: Record<string, unknown>;
    registry: Array<[string, GeneratedType]>;
    constructor(client: IgniteClient);
    updateTX(client: IgniteClient): void;
}
declare const IgntModule: (test: IgniteClient) => {
    module: {
        PerpdexPerpdexV_1: SDKModule;
    };
    registry: [string, GeneratedType][];
};
export default IgntModule;
