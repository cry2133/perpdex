import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Params } from "./params";
export declare const protobufPackage = "perpdex.perpdex.v1";
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
/** MsgCreatePrice defines the MsgCreatePrice message. */
export interface MsgCreatePrice {
    creator: string;
    symbol: string;
    price: string;
}
/** MsgCreatePriceResponse defines the MsgCreatePriceResponse message. */
export interface MsgCreatePriceResponse {
    id: number;
}
/** MsgUpdatePrice defines the MsgUpdatePrice message. */
export interface MsgUpdatePrice {
    creator: string;
    symbol: string;
    price: string;
    id: number;
}
/** MsgUpdatePriceResponse defines the MsgUpdatePriceResponse message. */
export interface MsgUpdatePriceResponse {
}
/** MsgDeletePrice defines the MsgDeletePrice message. */
export interface MsgDeletePrice {
    creator: string;
    id: number;
}
/** MsgDeletePriceResponse defines the MsgDeletePriceResponse message. */
export interface MsgDeletePriceResponse {
}
export declare const MsgUpdateParams: MessageFns<MsgUpdateParams>;
export declare const MsgUpdateParamsResponse: MessageFns<MsgUpdateParamsResponse>;
export declare const MsgCreatePrice: MessageFns<MsgCreatePrice>;
export declare const MsgCreatePriceResponse: MessageFns<MsgCreatePriceResponse>;
export declare const MsgUpdatePrice: MessageFns<MsgUpdatePrice>;
export declare const MsgUpdatePriceResponse: MessageFns<MsgUpdatePriceResponse>;
export declare const MsgDeletePrice: MessageFns<MsgDeletePrice>;
export declare const MsgDeletePriceResponse: MessageFns<MsgDeletePriceResponse>;
/** Msg defines the Msg service. */
export interface Msg {
    /**
     * UpdateParams defines a (governance) operation for updating the module
     * parameters. The authority defaults to the x/gov module account.
     */
    UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
    /** CreatePrice defines the CreatePrice RPC. */
    CreatePrice(request: MsgCreatePrice): Promise<MsgCreatePriceResponse>;
    /** UpdatePrice defines the UpdatePrice RPC. */
    UpdatePrice(request: MsgUpdatePrice): Promise<MsgUpdatePriceResponse>;
    /** DeletePrice defines the DeletePrice RPC. */
    DeletePrice(request: MsgDeletePrice): Promise<MsgDeletePriceResponse>;
}
export declare const MsgServiceName = "perpdex.perpdex.v1.Msg";
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    private readonly service;
    constructor(rpc: Rpc, opts?: {
        service?: string;
    });
    UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
    CreatePrice(request: MsgCreatePrice): Promise<MsgCreatePriceResponse>;
    UpdatePrice(request: MsgUpdatePrice): Promise<MsgUpdatePriceResponse>;
    DeletePrice(request: MsgDeletePrice): Promise<MsgDeletePriceResponse>;
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
