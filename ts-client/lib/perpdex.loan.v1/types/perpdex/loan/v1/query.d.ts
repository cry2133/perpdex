import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { Loan } from "./loan";
import { Params } from "./params";
export declare const protobufPackage = "perpdex.loan.v1";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
/** QueryGetLoanRequest defines the QueryGetLoanRequest message. */
export interface QueryGetLoanRequest {
    id: number;
}
/** QueryGetLoanResponse defines the QueryGetLoanResponse message. */
export interface QueryGetLoanResponse {
    loan: Loan | undefined;
}
/** QueryAllLoanRequest defines the QueryAllLoanRequest message. */
export interface QueryAllLoanRequest {
    pagination: PageRequest | undefined;
}
/** QueryAllLoanResponse defines the QueryAllLoanResponse message. */
export interface QueryAllLoanResponse {
    loan: Loan[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: MessageFns<QueryParamsRequest>;
export declare const QueryParamsResponse: MessageFns<QueryParamsResponse>;
export declare const QueryGetLoanRequest: MessageFns<QueryGetLoanRequest>;
export declare const QueryGetLoanResponse: MessageFns<QueryGetLoanResponse>;
export declare const QueryAllLoanRequest: MessageFns<QueryAllLoanRequest>;
export declare const QueryAllLoanResponse: MessageFns<QueryAllLoanResponse>;
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** ListLoan Queries a list of Loan items. */
    GetLoan(request: QueryGetLoanRequest): Promise<QueryGetLoanResponse>;
    /** ListLoan defines the ListLoan RPC. */
    ListLoan(request: QueryAllLoanRequest): Promise<QueryAllLoanResponse>;
}
export declare const QueryServiceName = "perpdex.loan.v1.Query";
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    private readonly service;
    constructor(rpc: Rpc, opts?: {
        service?: string;
    });
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    GetLoan(request: QueryGetLoanRequest): Promise<QueryGetLoanResponse>;
    ListLoan(request: QueryAllLoanRequest): Promise<QueryAllLoanResponse>;
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
