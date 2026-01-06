import { DeliverTxResponse, StdFee } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { IgniteClient } from "../client";
import { Api } from "./rest";
import { MsgUpdateParams } from "./types/perpdex/loan/v1/tx";
import { MsgRequestLoan } from "./types/perpdex/loan/v1/tx";
import { MsgCancelLoan } from "./types/perpdex/loan/v1/tx";
import { MsgApproveLoan } from "./types/perpdex/loan/v1/tx";
import { MsgRepayLoan } from "./types/perpdex/loan/v1/tx";
import { MsgLiquidateLoan } from "./types/perpdex/loan/v1/tx";
export { MsgUpdateParams, MsgRequestLoan, MsgCancelLoan, MsgApproveLoan, MsgRepayLoan, MsgLiquidateLoan };
type sendMsgUpdateParamsParams = {
    value: MsgUpdateParams;
    fee?: StdFee;
    memo?: string;
};
type sendMsgRequestLoanParams = {
    value: MsgRequestLoan;
    fee?: StdFee;
    memo?: string;
};
type sendMsgCancelLoanParams = {
    value: MsgCancelLoan;
    fee?: StdFee;
    memo?: string;
};
type sendMsgApproveLoanParams = {
    value: MsgApproveLoan;
    fee?: StdFee;
    memo?: string;
};
type sendMsgRepayLoanParams = {
    value: MsgRepayLoan;
    fee?: StdFee;
    memo?: string;
};
type sendMsgLiquidateLoanParams = {
    value: MsgLiquidateLoan;
    fee?: StdFee;
    memo?: string;
};
type msgUpdateParamsParams = {
    value: MsgUpdateParams;
};
type msgRequestLoanParams = {
    value: MsgRequestLoan;
};
type msgCancelLoanParams = {
    value: MsgCancelLoan;
};
type msgApproveLoanParams = {
    value: MsgApproveLoan;
};
type msgRepayLoanParams = {
    value: MsgRepayLoan;
};
type msgLiquidateLoanParams = {
    value: MsgLiquidateLoan;
};
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
    prefix: string;
    signer?: OfflineSigner;
}
export declare const txClient: ({ signer, prefix, addr }?: TxClientOptions) => {
    sendMsgUpdateParams({ value, fee, memo }: sendMsgUpdateParamsParams): Promise<DeliverTxResponse>;
    sendMsgRequestLoan({ value, fee, memo }: sendMsgRequestLoanParams): Promise<DeliverTxResponse>;
    sendMsgCancelLoan({ value, fee, memo }: sendMsgCancelLoanParams): Promise<DeliverTxResponse>;
    sendMsgApproveLoan({ value, fee, memo }: sendMsgApproveLoanParams): Promise<DeliverTxResponse>;
    sendMsgRepayLoan({ value, fee, memo }: sendMsgRepayLoanParams): Promise<DeliverTxResponse>;
    sendMsgLiquidateLoan({ value, fee, memo }: sendMsgLiquidateLoanParams): Promise<DeliverTxResponse>;
    msgUpdateParams({ value }: msgUpdateParamsParams): EncodeObject;
    msgRequestLoan({ value }: msgRequestLoanParams): EncodeObject;
    msgCancelLoan({ value }: msgCancelLoanParams): EncodeObject;
    msgApproveLoan({ value }: msgApproveLoanParams): EncodeObject;
    msgRepayLoan({ value }: msgRepayLoanParams): EncodeObject;
    msgLiquidateLoan({ value }: msgLiquidateLoanParams): EncodeObject;
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
        PerpdexLoanV_1: SDKModule;
    };
    registry: [string, GeneratedType][];
};
export default IgntModule;
