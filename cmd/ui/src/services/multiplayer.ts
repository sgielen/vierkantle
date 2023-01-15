import { createChannel, createClient } from 'nice-grpc-web';
import { VierkantleServiceDefinition, VierkantleServiceClient, TeamStreamClientMessage, TeamStreamServerMessage } from '../services/vierkantle';
import { grpc } from "@improbable-eng/grpc-web";

export class Multiplayer {
  public name: string;
  public onMessage: (msg: TeamStreamServerMessage) => void;
  public token: string;
  public players: string[];
  private queue: AsyncMessageQueue<TeamStreamClientMessage> | undefined;

  constructor(name: string, onMessage: (msg: TeamStreamServerMessage) => void, token?: string) {
    this.name = name;
    this.onMessage = onMessage;
    this.token = token ?? "";
    this.players = [];
  }

  public disconnect() {
    this.players = [];
    if (this.queue) {
      try {
        this.queue.close();
      } catch(e) {
        // ignore
      }
    }
    this.queue = undefined;
  }

  public async connect(): Promise<void> {
    this.disconnect();

    const host = window.location.origin + "/api";
    const channel = createChannel(host, grpc.WebsocketTransport());
    const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);
    this.queue = new AsyncMessageQueue<TeamStreamClientMessage>();

    // Initially, connect within this async function so that we can throw immediately if
    // connecting fails
    const stream = client.teamStream(this.queue);
    if (!this.token) {
      this.queue.push({
        create: { name: this.name },
      });
    } else {
      this.queue.push({
        join: { token: this.token, name: this.name },
      });
    }

    let haveFirstMessage!: (msg: TeamStreamServerMessage) => void;
    const firstMessage = new Promise<TeamStreamServerMessage>((resolve) => haveFirstMessage = resolve);

    // After initial handshake, read further messages in the background;
    // if this fails, reconnect. DON'T await here.
    this.runBackground(stream, haveFirstMessage);
    const response = await firstMessage;
    if (response.error) {
      this.disconnect();
      throw response.error.error;
    } else if (!response.team) {
      this.disconnect();
      throw "Invalid first message: " + JSON.stringify(response, null, 0);
    }
  }

  private async runBackground(stream: AsyncIterable<TeamStreamServerMessage>, firstMessage?: (msg: TeamStreamServerMessage) => void): Promise<void> {
    try {
      for await (const response of stream) {
        if (firstMessage) {
          firstMessage(response);
          firstMessage = undefined;
        }
        this.onMessage(response);
        if (response.team) {
          this.token = response.team.token;
          this.name = response.team.yourName;
          this.players = response.team.players;
        }
      }
    } catch (e) {
      console.log("Stream error: ", e);
    }

    if (this.queue) {
      // stream closed, but we should be connected, wait a bit then try to reconnect
      await new Promise((resolve) => setTimeout(resolve, 100));
      try {
        await this.connect();
      } catch(e) {
        console.log(e);
      }
    }
  }

  public async changePlayerName(name: string): Promise<void> {
    if (this.name == name) {
      return;
    }
    // TODO: send message to change player name; when
    // we receive new teaminfo, change playerName.value
  }

  public async sendWord(word: string): Promise<void> {
    this.queue?.push({
      word: {word}
    })
  }
}

class AsyncMessageQueue<T> {
  private closed = false;
  private queue: T[] = [];
  private waiter: (() => void) | null = null;

  public async *[Symbol.asyncIterator]() {
    while (!this.closed) {
      while (this.queue.length > 0) {
        yield this.queue.shift()!;
      }
      await new Promise<void>((resolve) => this.waiter = resolve);
    }
  }
  private wake() {
    if (this.waiter) {
      this.waiter();
      this.waiter = null;
    }
  }
  public push(message: T) {
    this.queue.push(message);
    this.wake();
  }
  public close() {
    this.closed = true;
    this.wake();
  }
}
