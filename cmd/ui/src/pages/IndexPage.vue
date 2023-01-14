<template>
  <q-header elevated>
    <q-toolbar>
      <q-btn
        class="only-if-small"
        flat
        dense
        round
        icon="menu"
        aria-label="Menu"
        @click="toggleWordList"
      />

      <q-toolbar-title>
        Vierkantle
      </q-toolbar-title>
    </q-toolbar>
  </q-header>

  <q-page-container>
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
              Of: <input type="text" v-model="token" placeholder="Code" /> <button @click="joinTeam">Join een team</button>
            </p>
            <p>
              <span v-if="multiplayerError" style="color: red;">{{ multiplayerError }}</span>
            </p>
          </template>
          <template v-else>
            <p>
              Je speelt in een team met {{ otherPlayers }}. Nodig meer mensen uit met
              de volgende code: <code>{{ multiplayer.token }}</code>
            </p>
          </template>
        </div>
      </div>

      <div class="row items-center q-ma-xl">
        <div class="wordlist" v-show="wordListOpen">
          <template v-for="(words, length) in wordsByLength" :key="length">
            <template v-if="words && words.filter(f => !f[1].bonus).length">
              <h6>{{ length }} letters</h6>
              <template v-for="([word, wstate]) in words" :key="word">
                <template v-if="!wstate.bonus">
                  <template v-if="wstate.guessed">
                    <tt>{{ word }}</tt><br/>
                  </template>
                  <template v-else>
                    <tt>{{ wordWithStars(word) }}</tt><br />
                  </template>
                </template>
              </template>
            </template>
          </template>
        </div>
        <div class="game-wrap items-center justify-evenly">
          <div class="row items-center justify-evenly">{{ wordsRemaining }} words remaining</div>
          <div class="row items-center justify-evenly">{{ wordMessage }}</div>
          <div class="row items-center justify-evenly">
            <div v-if="error">{{ error }}</div>
            <div v-else-if="!board">Loading board...</div>
            <div class="game" v-else>
              <VierkantleBoard
                :board="board"
                @word="word(null, $event)"
                @partialWord="partialWord($event)"
              />
            </div>
          </div>
        </div>
      </div>
    </q-page>
  </q-page-container>
</template>

<script setup lang="ts">
import VierkantleBoard from 'components/VierkantleBoard.vue';
import { Board, WordInBoard } from 'src/components/models';
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
  // TODO: Sometimes, I want to replace the board during the
  // day. So, always download the board, and replace the stored
  // one if it's different. Should find a way to do expiry, so
  // this code doesn't have to always download it
  /*
  const today = new Date().toISOString().slice(0, 10);

  if (board_.value && board_.value.loadedAt === today) {
    // Today's board is already present, keep it
    return
  }
  board_.value = null;
  */

  // Download a new board
  try {
    const f = await fetch("/board.json", {cache: 'no-store'});
    const board = await f.json() as Board;
    if (!board_.value || boardLetters(board) != boardLetters(board_.value)) {
      board_.value = board;
    }
  } catch(e) {
    error.value = e as string;
  }
});

const wordListOpen = ref(false);

function toggleWordList() {
  wordListOpen.value = !wordListOpen.value
}

function boardLetters(b: Board): string {
  return b.cells.reduce((p, c) => { return [...p, c.join("")] }).join("");
}

const wordsByLength = computed((): [string, WordInBoard][][] => {
  if (!board.value) {
    return []
  }
  var res = [] as [string, WordInBoard][][]
  Object.entries(board.value.words).forEach(([word, wordInBoard]) => {
    res[word.length] ??= [];
    res[word.length].push([word, wordInBoard]);
  })
  for (let i = 0; i < res.length; ++i) {
    if (res[i] && res[i].length) {
      res[i] = res[i].sort((a, b) => { return a[0].length - b[0].length });
    }
  }
  return res;
});

function wordWithStars(w: string): string {
  function xStars(i: number): string {
    var s = ""
    for (; i > 0; --i) {
      s += "*"
    }
    return s
  }


  switch(w.length) {
    case 4: return w.substring(0, 1) + xStars(3)
    case 5: return w.substring(0, 1) + xStars(4)
    case 6: return w.substring(0, 2) + xStars(4)
    case 7: return w.substring(0, 2) + xStars(4) + w.substring(6)
    default:
      return w.substring(0, 2) + xStars(w.length - 4) + w.substring(w.length - 2)
  }
}

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

function partialWord(word: string) {
  wordMessage.value = word;
}

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
  max-width: min(600px, 80vw, 80vh);
  aspect-ratio: 1 / 1;

  margin: 20px 40px;
  font-size: 40px;
}

.only-if-small {
  display: none;
}

.wordlist {
  z-index: 500;
  background: white;
  border: 1px solid black;
  overflow:scroll;
  padding: 4px;

  h6 {
    margin: 0;
    margin-top: 20px;
  }
}

@media screen and (min-width: 500px) {
  .wordlist {
    width: 25%;
    height: auto;
    display: block !important;
    visibility: visible !important;
  }

  .game-wrap {
    width: 75%;
    height: auto;
  }
}

@media screen and (max-width: 500px) {
  .game {
    font-size: 30px;
  }

  .only-if-small {
    display: block;
  }

  .wordlist {
    /* make modal */
    position: fixed;
    top: 60px;
    left: 20px;
    width: calc(100vw - 40px);
    height: calc(100vh - 80px);
    overflow: scroll;
    overscroll-behavior: contain;
  }

  .game-wrap {
    width: 100%;
    height: auto;
  }
}

@media screen and (max-height: 500px) {
  .game {
    font-size: 30px;
  }
}
</style>
