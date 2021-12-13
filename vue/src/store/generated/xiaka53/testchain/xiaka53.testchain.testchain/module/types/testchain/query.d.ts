import { Reader, Writer } from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Post } from "../testchain/post";
export declare const protobufPackage = "xiaka53.testchain.testchain";
export interface QueryPostsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryPostsResponse {
    Post: Post[];
    pagination: PageResponse | undefined;
}
export declare const QueryPostsRequest: {
    encode(message: QueryPostsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryPostsRequest;
    fromJSON(object: any): QueryPostsRequest;
    toJSON(message: QueryPostsRequest): unknown;
    fromPartial(object: DeepPartial<QueryPostsRequest>): QueryPostsRequest;
};
export declare const QueryPostsResponse: {
    encode(message: QueryPostsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryPostsResponse;
    fromJSON(object: any): QueryPostsResponse;
    toJSON(message: QueryPostsResponse): unknown;
    fromPartial(object: DeepPartial<QueryPostsResponse>): QueryPostsResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a list of posts items. */
    Posts(request: QueryPostsRequest): Promise<QueryPostsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Posts(request: QueryPostsRequest): Promise<QueryPostsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
