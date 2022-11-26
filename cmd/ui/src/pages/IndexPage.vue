<template>
  <q-page>
    <div class="row items-center justify-evenly" style="width: 400px; margin: 0 auto;">
      <div class="col">
        <template v-if="!multiplayer">
          <p>
            Je speelt alleen. Dat is prima, geen oordeel. Maar als je wel vrienden
            hebt, kies dan je naam en maak of join een team:
          </p>
          <p>
            Je naam: <input type="text" v-model="playerName" @change="changePlayerName" />
          </p>
          <p>
            <button @click="createTeam">Maak een team</button>
          </p>
          <p>
            Of: <input type="text" v-model="token" /> <button @click="joinTeam">Join een team</button>
          </p>
          <p>
            <span v-if="multiplayerError" style="color: red;">{{ multiplayerError }}</span>
          </p>
        </template>
        <template v-else>
          <p>
            Je speelt in een team met {{ otherPlayers }}. Nodig meer mensen uit met
            de volgende token: <code>{{ multiplayer.token }}</code>
          </p>
        </template>
      </div>
    </div>
    <div class="row items-center justify-evenly">{{ wordsRemaining }} words remaining</div>
    <div class="row items-center justify-evenly">{{ wordMessage }}</div>
    <div class="row items-center justify-evenly">
      <div v-if="error">{{ error }}</div>
      <div v-else-if="!board">Loading board...</div>
      <div class="game" v-else>
        <VierkantleBoard
          :board="board"
          @word="word(null, $event)"
        />
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import VierkantleBoard from 'components/VierkantleBoard.vue';
import { Board } from 'src/components/models';
import { computed, onMounted, ref } from 'vue';
import { StorageSerializers, useStorage } from '@vueuse/core';
import { createChannel, createClient } from 'nice-grpc-web';
import { VierkantleServiceDefinition, VierkantleServiceClient, TeamStreamClientMessage } from '../services/vierkantle';
import { grpc } from "@improbable-eng/grpc-web";

const board_ = useStorage<Board | undefined>("board", undefined, undefined, { serializer: StorageSerializers.object });

const board = computed(() => {
  return board_.value;
});

const error = ref("");

const host = window.location.origin + "/api";
const channel = createChannel(host, grpc.WebsocketTransport());
const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);

onMounted(async () => {
  const today = new Date().toISOString().slice(0, 10);

  if (board_.value && board_.value.loadedAt === today) {
    // Today's board is already present, keep it
    return
  }

  // Download a new board
  board_.value = null;
  try {
    const f = await fetch("/board.json");
    board_.value = await f.json() as Board;
    board_.value.loadedAt = today;
  } catch(e) {
    error.value = e as string;
  }
});

const wordMessage = ref(".");

const wordsRemaining = computed(() => {
  if (!board.value) {
    return 0;
  }
  let r = 0;
  Object.values(board.value.words).forEach((word) => {
    if (!word.bonus && !word.guessed) {
      r += 1;
    }
  });
  return r;
});

function word(who: string | null, word: string) {
  if (who) {
    wordMessage.value = `${who}: `
  } else {
    wordMessage.value = "";
  }

  if (word.length < 4) {
    wordMessage.value += "te kort: " + word
    return
  }
  const wordInBoard = board.value!.words[word];
  if (!wordInBoard) {
    wordMessage.value += "nope: " + word
    return
  }
  if (wordInBoard.bonus) {
    wordMessage.value += "bonus: " + word
  } else {
    wordMessage.value += word
  }
  if (wordInBoard.guessed) {
    wordMessage.value += " (had je al)"
  }
  wordInBoard.guessed = true
  if (multiplayer.value && !who) {
    multiplayer.value.queue.push({
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

interface MultiplayerStatus {
  token: string
  currentName: string
  players: string[]
  queue: AsyncMessageQueue<TeamStreamClientMessage>;
}

const playerName = ref<string>("Gast");
const token = ref<string>();
const multiplayer = ref<MultiplayerStatus>()
const multiplayerError = ref<string>();
const otherPlayers = computed(() => {
  if (!multiplayer.value) {
    return "niemand anders";
  }
  const list = multiplayer.value.players.filter((f) => f != multiplayer.value!.currentName);
  if (list.length == 0) {
    return "niemand anders";
  } else if (list.length == 1) {
    return list[0];
  } else {
    const last = list.pop();
    return list.join(", ") + " en " + last;
  }
});

function changePlayerName() {
  if (multiplayer.value && playerName.value != multiplayer.value.currentName) {
    // send message to change player name; when
    // we receive new teaminfo, change playerName.value
  }
}

async function openStream(token: string | null, name: string) {
  const queue = new AsyncMessageQueue<TeamStreamClientMessage>();
  multiplayerError.value = undefined;
  multiplayer.value = undefined;
  try {
    const stream = client.teamStream(queue);
    if (!token) {
      queue.push({
        create: { name },
      })
    } else {
      queue.push({
        join: { token, name },
      })
    }
    for await (const response of stream) {
      console.log("have message", response);
      if (response.error) {
        throw response.error.error;
      } else if (response.team) {
        multiplayer.value = {
          token: response.team.token,
          currentName: response.team.yourName,
          players: response.team.players,
          queue,
        };
      } else if (response.word) {
        word(response.word.player, response.word.word);
      }
    }
  } catch (e) {
    console.log("failed with: ", e);
    if (typeof e === "string") {
      multiplayerError.value = e;
    } else {
      multiplayerError.value = JSON.stringify(e, null, 0);
    }
    queue.close();
  }
}

function createTeam() {
  if (playerName.value) {
    openStream(null, playerName.value);
  }
}

function joinTeam() {
  if (token.value && playerName.value) {
    openStream(token.value, playerName.value);
  }
}
</script>

<style lang="scss">
.game {
  min-width: 300px;
  width: 100%;
  max-width: 600px;
  aspect-ratio: 1 / 1;

  margin: 20px 40px;
  font-size: 40px;
}

@media screen and (max-width: 400px) {
  .game {
    font-size: 30px;
  }
}
</style>
