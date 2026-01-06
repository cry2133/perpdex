import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Params } from "./params";
export declare const protobufPackage = "perpdex.loan.v1";
/** MsgUpdateParams is the Msg/UpdateParams request type. */
export interface MsgUpdateParams {
    /** authority is the address that controls the module (defaults to x/gov unless overwritten). */
    authority: string;
    /** NOTE: All parameters must be supplied. */
    params: Params | undefined;
}
/**
 * MsgUpdateParamsResponse defines the response structure for executing a
 * MsgUpdateParams message.
 */
export interface MsgUpdateParamsResponse {
}
/** MsgRequestLoan defines the MsgRequestLoan message. */
export interface MsgRequestLoan {
    creator: string;
    amount: string;
    fee: string;
    collateral: string;
    deadline: string;
}
/** MsgRequestLoanResponse defines the MsgRequestLoanResponse message. */
export interface MsgRequestLoanResponse {
}
/** MsgCancelLoan defines the MsgCancelLoan message. */
export interface MsgCancelLoan {
    creator: string;
    id: number;
}
/** MsgCancelLoanResponse defines the MsgCancelLoanResponse message. */
export interface MsgCancelLoanResponse {
}
/** MsgApproveLoan defines the MsgApproveLoan message. */
export interface MsgApproveLoan {
    creator: string;
    id: number;
}
/** MsgApproveLoanResponse defines the MsgApproveLoanResponse message. */
export interface MsgApproveLoanResponse {
}
/** MsgRepayLoan defines the MsgRepayLoan message. */
export interface MsgRepayLoan {
    creator: string;
    id: number;
}
/** MsgRepayLoanResponse defines the MsgRepayLoanResponse message. */
export interface MsgRepayLoanResponse {
}
/** MsgLiquidateLoan defines the MsgLiquidateLoan message. */
export interface MsgLiquidateLoan {
    creator: string;
    id: number;
}
/** MsgLiquidateLoanResponse defines the MsgLiquidateLoanResponse message. */
export interface MsgLiquidateLoanResponse {
}
export declare const MsgUpdateParams: MessageFns<MsgUpdateParams>;
export declare const MsgUpdateParamsResponse: MessageFns<MsgUpdateParamsResponse>;
export declare const MsgRequestLoan: MessageFns<MsgRequestLoan>;
export declare const MsgRequestLoanResponse: MessageFns<MsgRequestLoanResponse>;
export declare const MsgCancelLoan: MessageFns<MsgCancelLoan>;
export declare const MsgCancelLoanResponse: MessageFns<MsgCancelLoanResponse>;
export declare const MsgApproveLoan: MessageFns<MsgApproveLoan>;
export declare const MsgApproveLoanResponse: MessageFns<MsgApproveLoanResponse>;
export declare const MsgRepayLoan: MessageFns<MsgRepayLoan>;
export declare const MsgRepayLoanResponse: MessageFns<MsgRepayLoanResponse>;
export declare const MsgLiquidateLoan: MessageFns<MsgLiquidateLoan>;
export declare const MsgLiquidateLoanResponse: MessageFns<MsgLiquidateLoanResponse>;
/** Msg defines the Msg service. */
export interface Msg {
    /**
     * UpdateParams defines a (governance) operation for updating the module
     * parameters. The authority defaults to the x/gov module account.
     */
    UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
    /** RequestLoan defines the RequestLoan RPC. */
    RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse>;
    /** CancelLoan defines the CancelLoan RPC. */
    CancelLoan(request: MsgCancelLoan): Promise<MsgCancelLoanResponse>;
    /** ApproveLoan defines the ApproveLoan RPC. */
    ApproveLoan(request: MsgApproveLoan): Promise<MsgApproveLoanResponse>;
    /** RepayLoan defines the RepayLoan RPC. */
    RepayLoan(request: MsgRepayLoan): Promise<MsgRepayLoanResponse>;
    /** LiquidateLoan defines the LiquidateLoan RPC. */
    LiquidateLoan(request: MsgLiquidateLoan): Promise<MsgLiquidateLoanResponse>;
}
export declare const MsgServiceName = "perpdex.loan.v1.Msg";
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    private readonly service;
    constructor(rpc: Rpc, opts?: {
        service?: string;
    });
    UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
    RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse>;
    CancelLoan(request: MsgCancelLoan): Promise<MsgCancelLoanResponse>;
    ApproveLoan(request: MsgApproveLoan): Promise<MsgApproveLoanResponse>;
    RepayLoan(request: MsgRepayLoan): Promise<MsgRepayLoanResponse>;
    LiquidateLoan(request: MsgLiquidateLoan): Promise<MsgLiquidateLoanResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;
export type DeepPartial<T> = T extends Builtin ? T : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P : P & {
    [K in keyof P]: Exact<P[K], I[K]>;
} & {
    [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
};
export interface MessageFns<T> {
    encode(message: T, writer?: BinaryWriter): BinaryWriter;
    decode(input: BinaryReader | Uint8Array, length?: number): T;
    fromJSON(object: any): T;
    toJSON(message: T): unknown;
    create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
    fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
export {};
