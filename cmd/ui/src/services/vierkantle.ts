/* eslint-disable */
import Long from "long";
import type { CallContext, CallOptions } from "nice-grpc-common";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "nl.vierkantle";

export interface GetBoardRequest {
  /** Timezone offset in client's local time, as in Javascript Date.getTimezoneOffset() */
  timezoneOffsetMinutes: number;
}

export interface GetBoardResponse {
  board: Uint8Array;
  name: string;
}

export interface SubmitScoreRequest {
  anonymousId: number;
  teamSize: number;
  words: number;
  seconds: number;
  boardName: string;
}

export interface SubmitScoreResponse {
}

export interface GetScoresRequest {
  index: number;
  amount: number;
  myAnonymousId: number;
  boardName: string;
}

export interface GetScoresResponse {
  scores: { [key: number]: GetScoresResponse_Score };
  totalScores: number;
  yourScore?: number | undefined;
}

export interface GetScoresResponse_Score {
  name: string;
  teamSize: number;
  words: number;
  seconds: number;
}

export interface GetScoresResponse_ScoresEntry {
  key: number;
  value: GetScoresResponse_Score | undefined;
}

export interface WhoamiRequest {
}

export interface WhoamiResponse {
  username: string;
  admin: boolean;
}

export interface StartLoginRequest {
  /** Either username or email must be filled in. */
  username: string;
  email: string;
  urlPrefix: string;
}

export interface StartLoginResponse {
  found: boolean;
}

export interface FinishLoginRequest {
  token: string;
}

export interface FinishLoginResponse {
  username: string;
}

export interface RegisterRequest {
  username: string;
  /** Optional. */
  email: string;
}

export interface RegisterResponse {
}

export interface LogoutRequest {
}

export interface LogoutResponse {
}

export interface WordsForBoardRequest {
  board: Uint8Array;
}

export interface WordsForBoardResponse {
  board: Uint8Array;
  score: number;
}

export interface SeedBoardRequest {
  seedWord: string;
  width: number;
  height: number;
  attempts: number;
}

export interface SeedBoardResponse {
  board: Uint8Array;
  attempts: number;
  progress: number;
}

export interface FillInBoardRequest {
  board: Uint8Array;
  attempts: number;
}

export interface FillInBoardResponse {
  board: Uint8Array;
  attempts: number;
  progress: number;
}

export interface MarkWordTypeRequest {
  word: string;
  bonus: boolean;
}

export interface MarkWordTypeResponse {
}

export interface GetNewestBoardRequest {
}

export interface GetNewestBoardResponse {
  lastBoardName: string;
  boardsInQueue: number;
}

export interface AddBoardToQueueRequest {
  boardName: string;
  board: Uint8Array;
}

export interface AddBoardToQueueResponse {
  id: number;
}

export interface ListBoardQueueRequest {
}

export interface ListBoardQueueResponse {
  boards: ListBoardQueueResponse_Board[];
}

export interface ListBoardQueueResponse_Board {
  id: number;
  userId: number;
  username: string;
  boardName: string;
  addedAt: Date | undefined;
}

export interface GetBoardFromQueueRequest {
  id: number;
}

export interface GetBoardFromQueueResponse {
  boardName: string;
  board: Uint8Array;
}

export interface UpdateBoardInQueueRequest {
  id: number;
  boardName: string;
  board: Uint8Array;
}

export interface UpdateBoardInQueueResponse {
}

export interface RemoveBoardsFromQueueRequest {
  ids: number[];
}

export interface RemoveBoardsFromQueueResponse {
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
  return { board: new Uint8Array(), name: "" };
}

export const GetBoardResponse = {
  encode(message: GetBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
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

  fromPartial(object: DeepPartial<GetBoardResponse>): GetBoardResponse {
    const message = createBaseGetBoardResponse();
    message.board = object.board ?? new Uint8Array();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseSubmitScoreRequest(): SubmitScoreRequest {
  return { anonymousId: 0, teamSize: 0, words: 0, seconds: 0, boardName: "" };
}

export const SubmitScoreRequest = {
  encode(message: SubmitScoreRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.anonymousId !== 0) {
      writer.uint32(8).uint32(message.anonymousId);
    }
    if (message.teamSize !== 0) {
      writer.uint32(16).int32(message.teamSize);
    }
    if (message.words !== 0) {
      writer.uint32(24).int32(message.words);
    }
    if (message.seconds !== 0) {
      writer.uint32(32).int32(message.seconds);
    }
    if (message.boardName !== "") {
      writer.uint32(42).string(message.boardName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SubmitScoreRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubmitScoreRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.anonymousId = reader.uint32();
          break;
        case 2:
          message.teamSize = reader.int32();
          break;
        case 3:
          message.words = reader.int32();
          break;
        case 4:
          message.seconds = reader.int32();
          break;
        case 5:
          message.boardName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<SubmitScoreRequest>): SubmitScoreRequest {
    const message = createBaseSubmitScoreRequest();
    message.anonymousId = object.anonymousId ?? 0;
    message.teamSize = object.teamSize ?? 0;
    message.words = object.words ?? 0;
    message.seconds = object.seconds ?? 0;
    message.boardName = object.boardName ?? "";
    return message;
  },
};

function createBaseSubmitScoreResponse(): SubmitScoreResponse {
  return {};
}

export const SubmitScoreResponse = {
  encode(_: SubmitScoreResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SubmitScoreResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubmitScoreResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<SubmitScoreResponse>): SubmitScoreResponse {
    const message = createBaseSubmitScoreResponse();
    return message;
  },
};

function createBaseGetScoresRequest(): GetScoresRequest {
  return { index: 0, amount: 0, myAnonymousId: 0, boardName: "" };
}

export const GetScoresRequest = {
  encode(message: GetScoresRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== 0) {
      writer.uint32(8).uint32(message.index);
    }
    if (message.amount !== 0) {
      writer.uint32(16).uint32(message.amount);
    }
    if (message.myAnonymousId !== 0) {
      writer.uint32(24).uint32(message.myAnonymousId);
    }
    if (message.boardName !== "") {
      writer.uint32(34).string(message.boardName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetScoresRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetScoresRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.uint32();
          break;
        case 2:
          message.amount = reader.uint32();
          break;
        case 3:
          message.myAnonymousId = reader.uint32();
          break;
        case 4:
          message.boardName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetScoresRequest>): GetScoresRequest {
    const message = createBaseGetScoresRequest();
    message.index = object.index ?? 0;
    message.amount = object.amount ?? 0;
    message.myAnonymousId = object.myAnonymousId ?? 0;
    message.boardName = object.boardName ?? "";
    return message;
  },
};

function createBaseGetScoresResponse(): GetScoresResponse {
  return { scores: {}, totalScores: 0, yourScore: undefined };
}

export const GetScoresResponse = {
  encode(message: GetScoresResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.scores).forEach(([key, value]) => {
      GetScoresResponse_ScoresEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim();
    });
    if (message.totalScores !== 0) {
      writer.uint32(16).int32(message.totalScores);
    }
    if (message.yourScore !== undefined) {
      writer.uint32(32).int32(message.yourScore);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetScoresResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetScoresResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = GetScoresResponse_ScoresEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.scores[entry1.key] = entry1.value;
          }
          break;
        case 2:
          message.totalScores = reader.int32();
          break;
        case 4:
          message.yourScore = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetScoresResponse>): GetScoresResponse {
    const message = createBaseGetScoresResponse();
    message.scores = Object.entries(object.scores ?? {}).reduce<{ [key: number]: GetScoresResponse_Score }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[Number(key)] = GetScoresResponse_Score.fromPartial(value);
        }
        return acc;
      },
      {},
    );
    message.totalScores = object.totalScores ?? 0;
    message.yourScore = object.yourScore ?? undefined;
    return message;
  },
};

function createBaseGetScoresResponse_Score(): GetScoresResponse_Score {
  return { name: "", teamSize: 0, words: 0, seconds: 0 };
}

export const GetScoresResponse_Score = {
  encode(message: GetScoresResponse_Score, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.teamSize !== 0) {
      writer.uint32(16).int32(message.teamSize);
    }
    if (message.words !== 0) {
      writer.uint32(24).int32(message.words);
    }
    if (message.seconds !== 0) {
      writer.uint32(32).int32(message.seconds);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetScoresResponse_Score {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetScoresResponse_Score();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.teamSize = reader.int32();
          break;
        case 3:
          message.words = reader.int32();
          break;
        case 4:
          message.seconds = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetScoresResponse_Score>): GetScoresResponse_Score {
    const message = createBaseGetScoresResponse_Score();
    message.name = object.name ?? "";
    message.teamSize = object.teamSize ?? 0;
    message.words = object.words ?? 0;
    message.seconds = object.seconds ?? 0;
    return message;
  },
};

function createBaseGetScoresResponse_ScoresEntry(): GetScoresResponse_ScoresEntry {
  return { key: 0, value: undefined };
}

export const GetScoresResponse_ScoresEntry = {
  encode(message: GetScoresResponse_ScoresEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== 0) {
      writer.uint32(8).int32(message.key);
    }
    if (message.value !== undefined) {
      GetScoresResponse_Score.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetScoresResponse_ScoresEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetScoresResponse_ScoresEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.int32();
          break;
        case 2:
          message.value = GetScoresResponse_Score.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetScoresResponse_ScoresEntry>): GetScoresResponse_ScoresEntry {
    const message = createBaseGetScoresResponse_ScoresEntry();
    message.key = object.key ?? 0;
    message.value = (object.value !== undefined && object.value !== null)
      ? GetScoresResponse_Score.fromPartial(object.value)
      : undefined;
    return message;
  },
};

function createBaseWhoamiRequest(): WhoamiRequest {
  return {};
}

export const WhoamiRequest = {
  encode(_: WhoamiRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WhoamiRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWhoamiRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<WhoamiRequest>): WhoamiRequest {
    const message = createBaseWhoamiRequest();
    return message;
  },
};

function createBaseWhoamiResponse(): WhoamiResponse {
  return { username: "", admin: false };
}

export const WhoamiResponse = {
  encode(message: WhoamiResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.admin === true) {
      writer.uint32(16).bool(message.admin);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WhoamiResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWhoamiResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        case 2:
          message.admin = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<WhoamiResponse>): WhoamiResponse {
    const message = createBaseWhoamiResponse();
    message.username = object.username ?? "";
    message.admin = object.admin ?? false;
    return message;
  },
};

function createBaseStartLoginRequest(): StartLoginRequest {
  return { username: "", email: "", urlPrefix: "" };
}

export const StartLoginRequest = {
  encode(message: StartLoginRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    if (message.urlPrefix !== "") {
      writer.uint32(26).string(message.urlPrefix);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StartLoginRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStartLoginRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        case 2:
          message.email = reader.string();
          break;
        case 3:
          message.urlPrefix = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<StartLoginRequest>): StartLoginRequest {
    const message = createBaseStartLoginRequest();
    message.username = object.username ?? "";
    message.email = object.email ?? "";
    message.urlPrefix = object.urlPrefix ?? "";
    return message;
  },
};

function createBaseStartLoginResponse(): StartLoginResponse {
  return { found: false };
}

export const StartLoginResponse = {
  encode(message: StartLoginResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.found === true) {
      writer.uint32(8).bool(message.found);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StartLoginResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStartLoginResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.found = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<StartLoginResponse>): StartLoginResponse {
    const message = createBaseStartLoginResponse();
    message.found = object.found ?? false;
    return message;
  },
};

function createBaseFinishLoginRequest(): FinishLoginRequest {
  return { token: "" };
}

export const FinishLoginRequest = {
  encode(message: FinishLoginRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FinishLoginRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFinishLoginRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<FinishLoginRequest>): FinishLoginRequest {
    const message = createBaseFinishLoginRequest();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseFinishLoginResponse(): FinishLoginResponse {
  return { username: "" };
}

export const FinishLoginResponse = {
  encode(message: FinishLoginResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FinishLoginResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFinishLoginResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<FinishLoginResponse>): FinishLoginResponse {
    const message = createBaseFinishLoginResponse();
    message.username = object.username ?? "";
    return message;
  },
};

function createBaseRegisterRequest(): RegisterRequest {
  return { username: "", email: "" };
}

export const RegisterRequest = {
  encode(message: RegisterRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.email !== "") {
      writer.uint32(18).string(message.email);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.username = reader.string();
          break;
        case 2:
          message.email = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<RegisterRequest>): RegisterRequest {
    const message = createBaseRegisterRequest();
    message.username = object.username ?? "";
    message.email = object.email ?? "";
    return message;
  },
};

function createBaseRegisterResponse(): RegisterResponse {
  return {};
}

export const RegisterResponse = {
  encode(_: RegisterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<RegisterResponse>): RegisterResponse {
    const message = createBaseRegisterResponse();
    return message;
  },
};

function createBaseLogoutRequest(): LogoutRequest {
  return {};
}

export const LogoutRequest = {
  encode(_: LogoutRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LogoutRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLogoutRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<LogoutRequest>): LogoutRequest {
    const message = createBaseLogoutRequest();
    return message;
  },
};

function createBaseLogoutResponse(): LogoutResponse {
  return {};
}

export const LogoutResponse = {
  encode(_: LogoutResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LogoutResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLogoutResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<LogoutResponse>): LogoutResponse {
    const message = createBaseLogoutResponse();
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
  return { board: new Uint8Array(), score: 0 };
}

export const WordsForBoardResponse = {
  encode(message: WordsForBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    if (message.score !== 0) {
      writer.uint32(21).float(message.score);
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
        case 2:
          message.score = reader.float();
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
    message.score = object.score ?? 0;
    return message;
  },
};

function createBaseSeedBoardRequest(): SeedBoardRequest {
  return { seedWord: "", width: 0, height: 0, attempts: 0 };
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
    if (message.attempts !== 0) {
      writer.uint32(32).int32(message.attempts);
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
        case 4:
          message.attempts = reader.int32();
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
    message.attempts = object.attempts ?? 0;
    return message;
  },
};

function createBaseSeedBoardResponse(): SeedBoardResponse {
  return { board: new Uint8Array(), attempts: 0, progress: 0 };
}

export const SeedBoardResponse = {
  encode(message: SeedBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    if (message.attempts !== 0) {
      writer.uint32(16).int32(message.attempts);
    }
    if (message.progress !== 0) {
      writer.uint32(24).int32(message.progress);
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
        case 2:
          message.attempts = reader.int32();
          break;
        case 3:
          message.progress = reader.int32();
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
    message.attempts = object.attempts ?? 0;
    message.progress = object.progress ?? 0;
    return message;
  },
};

function createBaseFillInBoardRequest(): FillInBoardRequest {
  return { board: new Uint8Array(), attempts: 0 };
}

export const FillInBoardRequest = {
  encode(message: FillInBoardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    if (message.attempts !== 0) {
      writer.uint32(16).int32(message.attempts);
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
        case 2:
          message.attempts = reader.int32();
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
    message.attempts = object.attempts ?? 0;
    return message;
  },
};

function createBaseFillInBoardResponse(): FillInBoardResponse {
  return { board: new Uint8Array(), attempts: 0, progress: 0 };
}

export const FillInBoardResponse = {
  encode(message: FillInBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.board.length !== 0) {
      writer.uint32(10).bytes(message.board);
    }
    if (message.attempts !== 0) {
      writer.uint32(16).int32(message.attempts);
    }
    if (message.progress !== 0) {
      writer.uint32(24).int32(message.progress);
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
        case 2:
          message.attempts = reader.int32();
          break;
        case 3:
          message.progress = reader.int32();
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
    message.attempts = object.attempts ?? 0;
    message.progress = object.progress ?? 0;
    return message;
  },
};

function createBaseMarkWordTypeRequest(): MarkWordTypeRequest {
  return { word: "", bonus: false };
}

export const MarkWordTypeRequest = {
  encode(message: MarkWordTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.word !== "") {
      writer.uint32(10).string(message.word);
    }
    if (message.bonus === true) {
      writer.uint32(16).bool(message.bonus);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MarkWordTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMarkWordTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.word = reader.string();
          break;
        case 2:
          message.bonus = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<MarkWordTypeRequest>): MarkWordTypeRequest {
    const message = createBaseMarkWordTypeRequest();
    message.word = object.word ?? "";
    message.bonus = object.bonus ?? false;
    return message;
  },
};

function createBaseMarkWordTypeResponse(): MarkWordTypeResponse {
  return {};
}

export const MarkWordTypeResponse = {
  encode(_: MarkWordTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MarkWordTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMarkWordTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<MarkWordTypeResponse>): MarkWordTypeResponse {
    const message = createBaseMarkWordTypeResponse();
    return message;
  },
};

function createBaseGetNewestBoardRequest(): GetNewestBoardRequest {
  return {};
}

export const GetNewestBoardRequest = {
  encode(_: GetNewestBoardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetNewestBoardRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetNewestBoardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<GetNewestBoardRequest>): GetNewestBoardRequest {
    const message = createBaseGetNewestBoardRequest();
    return message;
  },
};

function createBaseGetNewestBoardResponse(): GetNewestBoardResponse {
  return { lastBoardName: "", boardsInQueue: 0 };
}

export const GetNewestBoardResponse = {
  encode(message: GetNewestBoardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.lastBoardName !== "") {
      writer.uint32(10).string(message.lastBoardName);
    }
    if (message.boardsInQueue !== 0) {
      writer.uint32(16).int64(message.boardsInQueue);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetNewestBoardResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetNewestBoardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lastBoardName = reader.string();
          break;
        case 2:
          message.boardsInQueue = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetNewestBoardResponse>): GetNewestBoardResponse {
    const message = createBaseGetNewestBoardResponse();
    message.lastBoardName = object.lastBoardName ?? "";
    message.boardsInQueue = object.boardsInQueue ?? 0;
    return message;
  },
};

function createBaseAddBoardToQueueRequest(): AddBoardToQueueRequest {
  return { boardName: "", board: new Uint8Array() };
}

export const AddBoardToQueueRequest = {
  encode(message: AddBoardToQueueRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boardName !== "") {
      writer.uint32(10).string(message.boardName);
    }
    if (message.board.length !== 0) {
      writer.uint32(18).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddBoardToQueueRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddBoardToQueueRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.boardName = reader.string();
          break;
        case 2:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<AddBoardToQueueRequest>): AddBoardToQueueRequest {
    const message = createBaseAddBoardToQueueRequest();
    message.boardName = object.boardName ?? "";
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseAddBoardToQueueResponse(): AddBoardToQueueResponse {
  return { id: 0 };
}

export const AddBoardToQueueResponse = {
  encode(message: AddBoardToQueueResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddBoardToQueueResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddBoardToQueueResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<AddBoardToQueueResponse>): AddBoardToQueueResponse {
    const message = createBaseAddBoardToQueueResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseListBoardQueueRequest(): ListBoardQueueRequest {
  return {};
}

export const ListBoardQueueRequest = {
  encode(_: ListBoardQueueRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListBoardQueueRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListBoardQueueRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<ListBoardQueueRequest>): ListBoardQueueRequest {
    const message = createBaseListBoardQueueRequest();
    return message;
  },
};

function createBaseListBoardQueueResponse(): ListBoardQueueResponse {
  return { boards: [] };
}

export const ListBoardQueueResponse = {
  encode(message: ListBoardQueueResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.boards) {
      ListBoardQueueResponse_Board.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListBoardQueueResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListBoardQueueResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.boards.push(ListBoardQueueResponse_Board.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<ListBoardQueueResponse>): ListBoardQueueResponse {
    const message = createBaseListBoardQueueResponse();
    message.boards = object.boards?.map((e) => ListBoardQueueResponse_Board.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListBoardQueueResponse_Board(): ListBoardQueueResponse_Board {
  return { id: 0, userId: 0, username: "", boardName: "", addedAt: undefined };
}

export const ListBoardQueueResponse_Board = {
  encode(message: ListBoardQueueResponse_Board, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(16).int64(message.userId);
    }
    if (message.username !== "") {
      writer.uint32(26).string(message.username);
    }
    if (message.boardName !== "") {
      writer.uint32(34).string(message.boardName);
    }
    if (message.addedAt !== undefined) {
      Timestamp.encode(toTimestamp(message.addedAt), writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListBoardQueueResponse_Board {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListBoardQueueResponse_Board();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.userId = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.username = reader.string();
          break;
        case 4:
          message.boardName = reader.string();
          break;
        case 5:
          message.addedAt = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<ListBoardQueueResponse_Board>): ListBoardQueueResponse_Board {
    const message = createBaseListBoardQueueResponse_Board();
    message.id = object.id ?? 0;
    message.userId = object.userId ?? 0;
    message.username = object.username ?? "";
    message.boardName = object.boardName ?? "";
    message.addedAt = object.addedAt ?? undefined;
    return message;
  },
};

function createBaseGetBoardFromQueueRequest(): GetBoardFromQueueRequest {
  return { id: 0 };
}

export const GetBoardFromQueueRequest = {
  encode(message: GetBoardFromQueueRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBoardFromQueueRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBoardFromQueueRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetBoardFromQueueRequest>): GetBoardFromQueueRequest {
    const message = createBaseGetBoardFromQueueRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseGetBoardFromQueueResponse(): GetBoardFromQueueResponse {
  return { boardName: "", board: new Uint8Array() };
}

export const GetBoardFromQueueResponse = {
  encode(message: GetBoardFromQueueResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boardName !== "") {
      writer.uint32(10).string(message.boardName);
    }
    if (message.board.length !== 0) {
      writer.uint32(18).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetBoardFromQueueResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetBoardFromQueueResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.boardName = reader.string();
          break;
        case 2:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<GetBoardFromQueueResponse>): GetBoardFromQueueResponse {
    const message = createBaseGetBoardFromQueueResponse();
    message.boardName = object.boardName ?? "";
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseUpdateBoardInQueueRequest(): UpdateBoardInQueueRequest {
  return { id: 0, boardName: "", board: new Uint8Array() };
}

export const UpdateBoardInQueueRequest = {
  encode(message: UpdateBoardInQueueRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).int64(message.id);
    }
    if (message.boardName !== "") {
      writer.uint32(18).string(message.boardName);
    }
    if (message.board.length !== 0) {
      writer.uint32(26).bytes(message.board);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateBoardInQueueRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBoardInQueueRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.boardName = reader.string();
          break;
        case 3:
          message.board = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<UpdateBoardInQueueRequest>): UpdateBoardInQueueRequest {
    const message = createBaseUpdateBoardInQueueRequest();
    message.id = object.id ?? 0;
    message.boardName = object.boardName ?? "";
    message.board = object.board ?? new Uint8Array();
    return message;
  },
};

function createBaseUpdateBoardInQueueResponse(): UpdateBoardInQueueResponse {
  return {};
}

export const UpdateBoardInQueueResponse = {
  encode(_: UpdateBoardInQueueResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateBoardInQueueResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateBoardInQueueResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<UpdateBoardInQueueResponse>): UpdateBoardInQueueResponse {
    const message = createBaseUpdateBoardInQueueResponse();
    return message;
  },
};

function createBaseRemoveBoardsFromQueueRequest(): RemoveBoardsFromQueueRequest {
  return { ids: [] };
}

export const RemoveBoardsFromQueueRequest = {
  encode(message: RemoveBoardsFromQueueRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();
    for (const v of message.ids) {
      writer.int64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RemoveBoardsFromQueueRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRemoveBoardsFromQueueRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.ids.push(longToNumber(reader.int64() as Long));
            }
          } else {
            message.ids.push(longToNumber(reader.int64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<RemoveBoardsFromQueueRequest>): RemoveBoardsFromQueueRequest {
    const message = createBaseRemoveBoardsFromQueueRequest();
    message.ids = object.ids?.map((e) => e) || [];
    return message;
  },
};

function createBaseRemoveBoardsFromQueueResponse(): RemoveBoardsFromQueueResponse {
  return {};
}

export const RemoveBoardsFromQueueResponse = {
  encode(_: RemoveBoardsFromQueueResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RemoveBoardsFromQueueResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRemoveBoardsFromQueueResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(_: DeepPartial<RemoveBoardsFromQueueResponse>): RemoveBoardsFromQueueResponse {
    const message = createBaseRemoveBoardsFromQueueResponse();
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
    submitScore: {
      name: "SubmitScore",
      requestType: SubmitScoreRequest,
      requestStream: false,
      responseType: SubmitScoreResponse,
      responseStream: false,
      options: {},
    },
    getScores: {
      name: "GetScores",
      requestType: GetScoresRequest,
      requestStream: false,
      responseType: GetScoresResponse,
      responseStream: false,
      options: {},
    },
    whoami: {
      name: "Whoami",
      requestType: WhoamiRequest,
      requestStream: false,
      responseType: WhoamiResponse,
      responseStream: false,
      options: {},
    },
    startLogin: {
      name: "StartLogin",
      requestType: StartLoginRequest,
      requestStream: false,
      responseType: StartLoginResponse,
      responseStream: false,
      options: {},
    },
    finishLogin: {
      name: "FinishLogin",
      requestType: FinishLoginRequest,
      requestStream: false,
      responseType: FinishLoginResponse,
      responseStream: false,
      options: {},
    },
    register: {
      name: "Register",
      requestType: RegisterRequest,
      requestStream: false,
      responseType: RegisterResponse,
      responseStream: false,
      options: {},
    },
    logout: {
      name: "Logout",
      requestType: LogoutRequest,
      requestStream: false,
      responseType: LogoutResponse,
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
      responseStream: true,
      options: {},
    },
    fillInBoard: {
      name: "FillInBoard",
      requestType: FillInBoardRequest,
      requestStream: false,
      responseType: FillInBoardResponse,
      responseStream: true,
      options: {},
    },
    markWordType: {
      name: "MarkWordType",
      requestType: MarkWordTypeRequest,
      requestStream: false,
      responseType: MarkWordTypeResponse,
      responseStream: false,
      options: {},
    },
    getNewestBoard: {
      name: "GetNewestBoard",
      requestType: GetNewestBoardRequest,
      requestStream: false,
      responseType: GetNewestBoardResponse,
      responseStream: false,
      options: {},
    },
    addBoardToQueue: {
      name: "AddBoardToQueue",
      requestType: AddBoardToQueueRequest,
      requestStream: false,
      responseType: AddBoardToQueueResponse,
      responseStream: false,
      options: {},
    },
    listBoardQueue: {
      name: "ListBoardQueue",
      requestType: ListBoardQueueRequest,
      requestStream: false,
      responseType: ListBoardQueueResponse,
      responseStream: false,
      options: {},
    },
    getBoardFromQueue: {
      name: "GetBoardFromQueue",
      requestType: GetBoardFromQueueRequest,
      requestStream: false,
      responseType: GetBoardFromQueueResponse,
      responseStream: false,
      options: {},
    },
    updateBoardInQueue: {
      name: "UpdateBoardInQueue",
      requestType: UpdateBoardInQueueRequest,
      requestStream: false,
      responseType: UpdateBoardInQueueResponse,
      responseStream: false,
      options: {},
    },
    removeBoardsFromQueue: {
      name: "RemoveBoardsFromQueue",
      requestType: RemoveBoardsFromQueueRequest,
      requestStream: false,
      responseType: RemoveBoardsFromQueueResponse,
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
  submitScore(
    request: SubmitScoreRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<SubmitScoreResponse>>;
  getScores(request: GetScoresRequest, context: CallContext & CallContextExt): Promise<DeepPartial<GetScoresResponse>>;
  whoami(request: WhoamiRequest, context: CallContext & CallContextExt): Promise<DeepPartial<WhoamiResponse>>;
  startLogin(
    request: StartLoginRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<StartLoginResponse>>;
  finishLogin(
    request: FinishLoginRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<FinishLoginResponse>>;
  register(request: RegisterRequest, context: CallContext & CallContextExt): Promise<DeepPartial<RegisterResponse>>;
  logout(request: LogoutRequest, context: CallContext & CallContextExt): Promise<DeepPartial<LogoutResponse>>;
  /** Generator */
  wordsForBoard(
    request: WordsForBoardRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<WordsForBoardResponse>>;
  seedBoard(
    request: SeedBoardRequest,
    context: CallContext & CallContextExt,
  ): ServerStreamingMethodResult<DeepPartial<SeedBoardResponse>>;
  fillInBoard(
    request: FillInBoardRequest,
    context: CallContext & CallContextExt,
  ): ServerStreamingMethodResult<DeepPartial<FillInBoardResponse>>;
  markWordType(
    request: MarkWordTypeRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<MarkWordTypeResponse>>;
  getNewestBoard(
    request: GetNewestBoardRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<GetNewestBoardResponse>>;
  addBoardToQueue(
    request: AddBoardToQueueRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<AddBoardToQueueResponse>>;
  listBoardQueue(
    request: ListBoardQueueRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<ListBoardQueueResponse>>;
  getBoardFromQueue(
    request: GetBoardFromQueueRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<GetBoardFromQueueResponse>>;
  updateBoardInQueue(
    request: UpdateBoardInQueueRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<UpdateBoardInQueueResponse>>;
  removeBoardsFromQueue(
    request: RemoveBoardsFromQueueRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<RemoveBoardsFromQueueResponse>>;
  teamStream(
    request: AsyncIterable<TeamStreamClientMessage>,
    context: CallContext & CallContextExt,
  ): ServerStreamingMethodResult<DeepPartial<TeamStreamServerMessage>>;
}

export interface VierkantleServiceClient<CallOptionsExt = {}> {
  getBoard(request: DeepPartial<GetBoardRequest>, options?: CallOptions & CallOptionsExt): Promise<GetBoardResponse>;
  submitScore(
    request: DeepPartial<SubmitScoreRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<SubmitScoreResponse>;
  getScores(request: DeepPartial<GetScoresRequest>, options?: CallOptions & CallOptionsExt): Promise<GetScoresResponse>;
  whoami(request: DeepPartial<WhoamiRequest>, options?: CallOptions & CallOptionsExt): Promise<WhoamiResponse>;
  startLogin(
    request: DeepPartial<StartLoginRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<StartLoginResponse>;
  finishLogin(
    request: DeepPartial<FinishLoginRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<FinishLoginResponse>;
  register(request: DeepPartial<RegisterRequest>, options?: CallOptions & CallOptionsExt): Promise<RegisterResponse>;
  logout(request: DeepPartial<LogoutRequest>, options?: CallOptions & CallOptionsExt): Promise<LogoutResponse>;
  /** Generator */
  wordsForBoard(
    request: DeepPartial<WordsForBoardRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<WordsForBoardResponse>;
  seedBoard(
    request: DeepPartial<SeedBoardRequest>,
    options?: CallOptions & CallOptionsExt,
  ): AsyncIterable<SeedBoardResponse>;
  fillInBoard(
    request: DeepPartial<FillInBoardRequest>,
    options?: CallOptions & CallOptionsExt,
  ): AsyncIterable<FillInBoardResponse>;
  markWordType(
    request: DeepPartial<MarkWordTypeRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<MarkWordTypeResponse>;
  getNewestBoard(
    request: DeepPartial<GetNewestBoardRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<GetNewestBoardResponse>;
  addBoardToQueue(
    request: DeepPartial<AddBoardToQueueRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<AddBoardToQueueResponse>;
  listBoardQueue(
    request: DeepPartial<ListBoardQueueRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<ListBoardQueueResponse>;
  getBoardFromQueue(
    request: DeepPartial<GetBoardFromQueueRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<GetBoardFromQueueResponse>;
  updateBoardInQueue(
    request: DeepPartial<UpdateBoardInQueueRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<UpdateBoardInQueueResponse>;
  removeBoardsFromQueue(
    request: DeepPartial<RemoveBoardsFromQueueRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<RemoveBoardsFromQueueResponse>;
  teamStream(
    request: AsyncIterable<DeepPartial<TeamStreamClientMessage>>,
    options?: CallOptions & CallOptionsExt,
  ): AsyncIterable<TeamStreamServerMessage>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

export type ServerStreamingMethodResult<Response> = { [Symbol.asyncIterator](): AsyncIterator<Response, void> };
