/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "nl.vierkantle";

export interface GetBoardRequest {
  timezoneOffsetMinutes: number;
}

export interface GetBoardResponse {
  board: Uint8Array;
}

export interface WordsForBoardRequest {
  board: Uint8Array;
}

export interface WordsForBoardResponse {
  board: Uint8Array;
}

export interface SeedBoardRequest {
  seedWord: string;
  width: number;
  height: number;
}

export interface SeedBoardResponse {
  board: Uint8Array;
}

export interface FillInBoardRequest {
  board: Uint8Array;
}

export interface FillInBoardResponse {
  board: Uint8Array;
}

export interface CreateTeamRequest {
  name: string;
}

export interface JoinTeamRequest {
  token: string;
  name: string;
}

export interface ChangeNameRequest {
  name: string;
}

export interface WordGuessedRequest {
  word: string;
}

export interface TeamStreamClientMessage {
  /**
   * The stream must start with one create or join message. After this, the
   * server responds with TeamInfoResponse, after which the create or join
   * must not be sent anymore. If the player name already exists in the team,
   * the server will respond with an ErrorResponse and the stream will be
   * closed.
   */
  create?: CreateTeamRequest | undefined;
  join?:
    | JoinTeamRequest
    | undefined;
  /**
   * When the player changes their name, send a ChangeNameRequest; the
   * server will respond with a new TeamInfoResponse to everyone. If the
   * name is already in use, the server will respond with an ErrorResponse
   * and the name will not be changed.
   */
  name?:
    | ChangeNameRequest
    | undefined;
  /**
   * When the player guesses a word, send this request; the server will
   * send a WordGuessedResponse to all other players.
   */
  word?: WordGuessedRequest | undefined;
}

export interface ErrorResponse {
  error: string;
}

export interface TeamInfoResponse {
  token: string;
  yourName: string;
  players: string[];
}

export interface WordGuessedResponse {
  player: string;
  word: string;
}

export interface TeamStreamServerMessage {
  error?: ErrorResponse | undefined;
  team?: TeamInfoResponse | undefined;
  word?: WordGuessedResponse | undefined;
}

function createBaseGetBoardRequest(): GetBoardRequest {
  return { timezoneOffsetMinutes: 0 };
}

export const GetBoardRequest = {
  encode(message: GetBoardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.timezoneOffsetMinutes !== 0) {
      writer.uint32(8).int32(message.timezoneOffsetMinutes);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBoardRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBoardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.timezoneOffsetMinutes = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetBoardRequest>): GetBoardRequest {
    const message = createBaseGetBoardRequest();
    message.timezoneOffsetMinutes = object.timezoneOffsetMinutes ?? 0;
    return message;
  },
};

function createBaseGetBoardResponse(): GetBoardResponse {
  return { board: new Uint8Array() };
}

export const GetBoardResponse = {
  encode(message: GetBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBoardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBoardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetBoardResponse>): GetBoardResponse {
    const message = createBaseGetBoardResponse();
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseWordsForBoardRequest(): WordsForBoardRequest {
  return { board: new Uint8Array() };
}

export const WordsForBoardRequest = {
  encode(message: WordsForBoardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WordsForBoardRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWordsForBoardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<WordsForBoardRequest>): WordsForBoardRequest {
    const message = createBaseWordsForBoardRequest();
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseWordsForBoardResponse(): WordsForBoardResponse {
  return { board: new Uint8Array() };
}

export const WordsForBoardResponse = {
  encode(message: WordsForBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WordsForBoardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWordsForBoardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<WordsForBoardResponse>): WordsForBoardResponse {
    const message = createBaseWordsForBoardResponse();
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseSeedBoardRequest(): SeedBoardRequest {
  return { seedWord: "", width: 0, height: 0 };
}

export const SeedBoardRequest = {
  encode(message: SeedBoardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.seedWord !== "") {
      writer.uint32(10).string(message.seedWord);
    }
    if (message.width !== 0) {
      writer.uint32(16).int32(message.width);
    }
    if (message.height !== 0) {
      writer.uint32(24).int32(message.height);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SeedBoardRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSeedBoardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.seedWord = reader.string();
          break;
        case 2:
          message.width = reader.int32();
          break;
        case 3:
          message.height = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<SeedBoardRequest>): SeedBoardRequest {
    const message = createBaseSeedBoardRequest();
    message.seedWord = object.seedWord ?? "";
    message.width = object.width ?? 0;
    message.height = object.height ?? 0;
    return message;
  },
};

function createBaseSeedBoardResponse(): SeedBoardResponse {
  return { board: new Uint8Array() };
}

export const SeedBoardResponse = {
  encode(message: SeedBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SeedBoardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSeedBoardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<SeedBoardResponse>): SeedBoardResponse {
    const message = createBaseSeedBoardResponse();
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseFillInBoardRequest(): FillInBoardRequest {
  return { board: new Uint8Array() };
}

export const FillInBoardRequest = {
  encode(message: FillInBoardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FillInBoardRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFillInBoardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<FillInBoardRequest>): FillInBoardRequest {
    const message = createBaseFillInBoardRequest();
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseFillInBoardResponse(): FillInBoardResponse {
  return { board: new Uint8Array() };
}

export const FillInBoardResponse = {
  encode(message: FillInBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FillInBoardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFillInBoardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<FillInBoardResponse>): FillInBoardResponse {
    const message = createBaseFillInBoardResponse();
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseCreateTeamRequest(): CreateTeamRequest {
  return { name: "" };
}

export const CreateTeamRequest = {
  encode(message: CreateTeamRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTeamRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTeamRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<CreateTeamRequest>): CreateTeamRequest {
    const message = createBaseCreateTeamRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseJoinTeamRequest(): JoinTeamRequest {
  return { token: "", name: "" };
}

export const JoinTeamRequest = {
  encode(message: JoinTeamRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JoinTeamRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJoinTeamRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<JoinTeamRequest>): JoinTeamRequest {
    const message = createBaseJoinTeamRequest();
    message.token = object.token ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseChangeNameRequest(): ChangeNameRequest {
  return { name: "" };
}

export const ChangeNameRequest = {
  encode(message: ChangeNameRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChangeNameRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChangeNameRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<ChangeNameRequest>): ChangeNameRequest {
    const message = createBaseChangeNameRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseWordGuessedRequest(): WordGuessedRequest {
  return { word: "" };
}

export const WordGuessedRequest = {
  encode(message: WordGuessedRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.word !== "") {
      writer.uint32(10).string(message.word);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WordGuessedRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWordGuessedRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.word = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<WordGuessedRequest>): WordGuessedRequest {
    const message = createBaseWordGuessedRequest();
    message.word = object.word ?? "";
    return message;
  },
};

function createBaseTeamStreamClientMessage(): TeamStreamClientMessage {
  return { create: undefined, join: undefined, name: undefined, word: undefined };
}

export const TeamStreamClientMessage = {
  encode(message: TeamStreamClientMessage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.create !== undefined) {
      CreateTeamRequest.encode(message.create, writer.uint32(10).fork()).ldelim();
    }
    if (message.join !== undefined) {
      JoinTeamRequest.encode(message.join, writer.uint32(18).fork()).ldelim();
    }
    if (message.name !== undefined) {
      ChangeNameRequest.encode(message.name, writer.uint32(26).fork()).ldelim();
    }
    if (message.word !== undefined) {
      WordGuessedRequest.encode(message.word, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TeamStreamClientMessage {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTeamStreamClientMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.create = CreateTeamRequest.decode(reader, reader.uint32());
          break;
        case 2:
          message.join = JoinTeamRequest.decode(reader, reader.uint32());
          break;
        case 3:
          message.name = ChangeNameRequest.decode(reader, reader.uint32());
          break;
        case 4:
          message.word = WordGuessedRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<TeamStreamClientMessage>): TeamStreamClientMessage {
    const message = createBaseTeamStreamClientMessage();
    message.create = (object.create !== undefined && object.create !== null)
      ? CreateTeamRequest.fromPartial(object.create)
      : undefined;
    message.join = (object.join !== undefined && object.join !== null)
      ? JoinTeamRequest.fromPartial(object.join)
      : undefined;
    message.name = (object.name !== undefined && object.name !== null)
      ? ChangeNameRequest.fromPartial(object.name)
      : undefined;
    message.word = (object.word !== undefined && object.word !== null)
      ? WordGuessedRequest.fromPartial(object.word)
      : undefined;
    return message;
  },
};

function createBaseErrorResponse(): ErrorResponse {
  return { error: "" };
}

export const ErrorResponse = {
  encode(message: ErrorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.error !== "") {
      writer.uint32(10).string(message.error);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ErrorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseErrorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.error = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<ErrorResponse>): ErrorResponse {
    const message = createBaseErrorResponse();
    message.error = object.error ?? "";
    return message;
  },
};

function createBaseTeamInfoResponse(): TeamInfoResponse {
  return { token: "", yourName: "", players: [] };
}

export const TeamInfoResponse = {
  encode(message: TeamInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.yourName !== "") {
      writer.uint32(18).string(message.yourName);
    }
    for (const v of message.players) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TeamInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTeamInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.yourName = reader.string();
          break;
        case 3:
          message.players.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<TeamInfoResponse>): TeamInfoResponse {
    const message = createBaseTeamInfoResponse();
    message.token = object.token ?? "";
    message.yourName = object.yourName ?? "";
    message.players = object.players?.map((e) => e) || [];
    return message;
  },
};

function createBaseWordGuessedResponse(): WordGuessedResponse {
  return { player: "", word: "" };
}

export const WordGuessedResponse = {
  encode(message: WordGuessedResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.player !== "") {
      writer.uint32(10).string(message.player);
    }
    if (message.word !== "") {
      writer.uint32(18).string(message.word);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WordGuessedResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWordGuessedResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player = reader.string();
          break;
        case 2:
          message.word = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<WordGuessedResponse>): WordGuessedResponse {
    const message = createBaseWordGuessedResponse();
    message.player = object.player ?? "";
    message.word = object.word ?? "";
    return message;
  },
};

function createBaseTeamStreamServerMessage(): TeamStreamServerMessage {
  return { error: undefined, team: undefined, word: undefined };
}

export const TeamStreamServerMessage = {
  encode(message: TeamStreamServerMessage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.error !== undefined) {
      ErrorResponse.encode(message.error, writer.uint32(10).fork()).ldelim();
    }
    if (message.team !== undefined) {
      TeamInfoResponse.encode(message.team, writer.uint32(18).fork()).ldelim();
    }
    if (message.word !== undefined) {
      WordGuessedResponse.encode(message.word, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TeamStreamServerMessage {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTeamStreamServerMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.error = ErrorResponse.decode(reader, reader.uint32());
          break;
        case 2:
          message.team = TeamInfoResponse.decode(reader, reader.uint32());
          break;
        case 3:
          message.word = WordGuessedResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<TeamStreamServerMessage>): TeamStreamServerMessage {
    const message = createBaseTeamStreamServerMessage();
    message.error = (object.error !== undefined && object.error !== null)
      ? ErrorResponse.fromPartial(object.error)
      : undefined;
    message.team = (object.team !== undefined && object.team !== null)
      ? TeamInfoResponse.fromPartial(object.team)
      : undefined;
    message.word = (object.word !== undefined && object.word !== null)
      ? WordGuessedResponse.fromPartial(object.word)
      : undefined;
    return message;
  },
};

export type VierkantleServiceDefinition = typeof VierkantleServiceDefinition;
export const VierkantleServiceDefinition = {
  name: "VierkantleService",
  fullName: "nl.vierkantle.VierkantleService",
  methods: {
    getBoard: {
      name: "GetBoard",
      requestType: GetBoardRequest,
      requestStream: false,
      responseType: GetBoardResponse,
      responseStream: false,
      options: {},
    },
    /** Generator */
    wordsForBoard: {
      name: "WordsForBoard",
      requestType: WordsForBoardRequest,
      requestStream: false,
      responseType: WordsForBoardResponse,
      responseStream: false,
      options: {},
    },
    seedBoard: {
      name: "SeedBoard",
      requestType: SeedBoardRequest,
      requestStream: false,
      responseType: SeedBoardResponse,
      responseStream: false,
      options: {},
    },
    fillInBoard: {
      name: "FillInBoard",
      requestType: FillInBoardRequest,
      requestStream: false,
      responseType: FillInBoardResponse,
      responseStream: false,
      options: {},
    },
    teamStream: {
      name: "TeamStream",
      requestType: TeamStreamClientMessage,
      requestStream: true,
      responseType: TeamStreamServerMessage,
      responseStream: true,
      options: {},
    },
  },
} as const;

export interface VierkantleServiceImplementation<CallContextExt = {}> {
  getBoard(request: GetBoardRequest, context: CallContext & CallContextExt): Promise<DeepPartial<GetBoardResponse>>;
  /** Generator */
  wordsForBoard(
    request: WordsForBoardRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<WordsForBoardResponse>>;
  seedBoard(request: SeedBoardRequest, context: CallContext & CallContextExt): Promise<DeepPartial<SeedBoardResponse>>;
  fillInBoard(
    request: FillInBoardRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<FillInBoardResponse>>;
  teamStream(
    request: AsyncIterable<TeamStreamClientMessage>,
    context: CallContext & CallContextExt,
  ): ServerStreamingMethodResult<DeepPartial<TeamStreamServerMessage>>;
}

export interface VierkantleServiceClient<CallOptionsExt = {}> {
  getBoard(request: DeepPartial<GetBoardRequest>, options?: CallOptions & CallOptionsExt): Promise<GetBoardResponse>;
  /** Generator */
  wordsForBoard(
    request: DeepPartial<WordsForBoardRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<WordsForBoardResponse>;
  seedBoard(request: DeepPartial<SeedBoardRequest>, options?: CallOptions & CallOptionsExt): Promise<SeedBoardResponse>;
  fillInBoard(
    request: DeepPartial<FillInBoardRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<FillInBoardResponse>;
  teamStream(
    request: AsyncIterable<DeepPartial<TeamStreamClientMessage>>,
    options?: CallOptions & CallOptionsExt,
  ): AsyncIterable<TeamStreamServerMessage>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

export type ServerStreamingMethodResult<Response> = { [Symbol.asyncIterator](): AsyncIterator<Response, void> };
