import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import { Price } from "./price";
export declare const protobufPackage = "perpdex.perpdex.v1";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
/** QueryShowPriceRequest defines the QueryShowPriceRequest message. */
export interface QueryShowPriceRequest {
    id: number;
}
/** QueryShowPriceResponse defines the QueryShowPriceResponse message. */
export interface QueryShowPriceResponse {
    post: Price | undefined;
}
/** QueryShowSymbolPriceRequest defines the QueryShowSymbolPriceRequest message. */
export interface QueryShowSymbolPriceRequest {
    symbol: string;
}
/** QueryShowSymbolPriceResponse defines the QueryShowSymbolPriceResponse message. */
export interface QueryShowSymbolPriceResponse {
    post: Price | undefined;
}
/** QueryListPriceRequest defines the QueryListPriceRequest message. */
export interface QueryListPriceRequest {
    pagination: PageRequest | undefined;
}
/** QueryListPriceResponse defines the QueryListPriceResponse message. */
export interface QueryListPriceResponse {
    post: Price[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: MessageFns<QueryParamsRequest>;
export declare const QueryParamsResponse: MessageFns<QueryParamsResponse>;
export declare const QueryShowPriceRequest: MessageFns<QueryShowPriceRequest>;
export declare const QueryShowPriceResponse: MessageFns<QueryShowPriceResponse>;
export declare const QueryShowSymbolPriceRequest: MessageFns<QueryShowSymbolPriceRequest>;
export declare const QueryShowSymbolPriceResponse: MessageFns<QueryShowSymbolPriceResponse>;
export declare const QueryListPriceRequest: MessageFns<QueryListPriceRequest>;
export declare const QueryListPriceResponse: MessageFns<QueryListPriceResponse>;
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** ShowPrice Queries a list of ShowPrice items. */
    ShowPrice(request: QueryShowPriceRequest): Promise<QueryShowPriceResponse>;
    /** ShowSymbolPrice Queries a list of ShowSymbolPrice items. */
    ShowSymbolPrice(request: QueryShowSymbolPriceRequest): Promise<QueryShowSymbolPriceResponse>;
    /** ListPrice Queries a list of ListPrice items. */
    ListPrice(request: QueryListPriceRequest): Promise<QueryListPriceResponse>;
}
export declare const QueryServiceName = "perpdex.perpdex.v1.Query";
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    private readonly service;
    constructor(rpc: Rpc, opts?: {
        service?: string;
    });
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    ShowPrice(request: QueryShowPriceRequest): Promise<QueryShowPriceResponse>;
    ShowSymbolPrice(request: QueryShowSymbolPriceRequest): Promise<QueryShowSymbolPriceResponse>;
    ListPrice(request: QueryListPriceRequest): Promise<QueryListPriceResponse>;
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
