/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "nl.vierkantle";

export interface HelloRequest {
  m: string;
}

export interface HelloResponse {
  m: string;
}

export interface HelloStreamRequest {
  m: string;
}

export interface HelloStreamResponse {
  m: string;
}

function createBaseHelloRequest(): HelloRequest {
  return { m: "" };
}

export const HelloRequest = {
  encode(message: HelloRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.m !== "") {
      writer.uint32(10).string(message.m);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HelloRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHelloRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.m = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<HelloRequest>): HelloRequest {
    const message = createBaseHelloRequest();
    message.m = object.m ?? "";
    return message;
  },
};

function createBaseHelloResponse(): HelloResponse {
  return { m: "" };
}

export const HelloResponse = {
  encode(message: HelloResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.m !== "") {
      writer.uint32(10).string(message.m);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HelloResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHelloResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.m = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<HelloResponse>): HelloResponse {
    const message = createBaseHelloResponse();
    message.m = object.m ?? "";
    return message;
  },
};

function createBaseHelloStreamRequest(): HelloStreamRequest {
  return { m: "" };
}

export const HelloStreamRequest = {
  encode(message: HelloStreamRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.m !== "") {
      writer.uint32(10).string(message.m);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HelloStreamRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHelloStreamRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.m = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<HelloStreamRequest>): HelloStreamRequest {
    const message = createBaseHelloStreamRequest();
    message.m = object.m ?? "";
    return message;
  },
};

function createBaseHelloStreamResponse(): HelloStreamResponse {
  return { m: "" };
}

export const HelloStreamResponse = {
  encode(message: HelloStreamResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.m !== "") {
      writer.uint32(10).string(message.m);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HelloStreamResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHelloStreamResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.m = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<HelloStreamResponse>): HelloStreamResponse {
    const message = createBaseHelloStreamResponse();
    message.m = object.m ?? "";
    return message;
  },
};

export type VierkantleServiceDefinition = typeof VierkantleServiceDefinition;
export const VierkantleServiceDefinition = {
  name: "VierkantleService",
  fullName: "nl.vierkantle.VierkantleService",
  methods: {
    hello: {
      name: "Hello",
      requestType: HelloRequest,
      requestStream: false,
      responseType: HelloResponse,
      responseStream: false,
      options: {},
    },
    helloStream: {
      name: "HelloStream",
      requestType: HelloStreamRequest,
      requestStream: true,
      responseType: HelloStreamResponse,
      responseStream: true,
      options: {},
    },
  },
} as const;

export interface VierkantleServiceImplementation<CallContextExt = {}> {
  hello(request: HelloRequest, context: CallContext & CallContextExt): Promise<DeepPartial<HelloResponse>>;
  helloStream(
    request: AsyncIterable<HelloStreamRequest>,
    context: CallContext & CallContextExt,
  ): ServerStreamingMethodResult<DeepPartial<HelloStreamResponse>>;
}

export interface VierkantleServiceClient<CallOptionsExt = {}> {
  hello(request: DeepPartial<HelloRequest>, options?: CallOptions & CallOptionsExt): Promise<HelloResponse>;
  helloStream(
    request: AsyncIterable<DeepPartial<HelloStreamRequest>>,
    options?: CallOptions & CallOptionsExt,
  ): AsyncIterable<HelloStreamResponse>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

export type ServerStreamingMethodResult<Response> = { [Symbol.asyncIterator](): AsyncIterator<Response, void> };
