<template>
  <q-header elevated>
    <q-toolbar>
      <q-btn
        class="only-if-small q-mx-sm"
        dense
        color="secondary"
        icon="spellcheck"
        aria-label="Word list"
        @click="toggleWordList"
      />

      <q-btn
        class="q-mx-sm"
        dense
        color="secondary"
        icon="groups"
        aria-label="Multiplayer"
        @click="toggleMultiplayer"
      />

      <q-toolbar-title>
        Vierkantle
      </q-toolbar-title>
    </q-toolbar>
  </q-header>

  <q-page-container>
    <q-page>
      <q-dialog v-model="multiplayerOpen">
        <q-card style="width: 650px">
          <q-card-section>
            <p class="text-h6">Multiplayer-modus</p>
          </q-card-section>
          <q-separator />
          <q-card-section v-if="multiplayerError">
            <span style="color: red">{{ multiplayerError }}</span>
          </q-card-section>
          <template v-if="!multiplayer">
            <q-card-section>
              Vul je spelernaam in, en join of maak een team:
              <q-input dense outlined label="Je naam" v-model="playerName" @update:model-value="changePlayerName" />
            </q-card-section>
            <q-separator />
            <q-card-section>
              <q-input dense outlined v-model="token" label="Code" />
              <q-btn @click="joinTeam" label="Join een team" color="primary" />
            </q-card-section>
            <q-card-section>
              Of:
              <q-btn @click="createTeam" label="Maak een team" color="primary" />
            </q-card-section>
          </template>
          <template v-else>
            <q-card-section>
              Je speelt in een team met {{ otherPlayers }}. Nodig meer mensen
              uit via de volgende URL:<br />
              <code
                @click="setClipboard(multiplayerUrl)"
                style="cursor: pointer;"
              >{{ multiplayerUrl }}</code> (<span>{{ copyText }}</span>)
            </q-card-section>
            <q-separator />
            <q-card-section>
              <q-btn @click="multiplayerOpen = false" color="primary" label="Sluiten" />
              <q-btn @click="stopMultiplayer" label="Verlaat team" />
            </q-card-section>
          </template>
        </q-card>
      </q-dialog>

      <div class="game-page row items-center q-ma-xl">
        <div class="wordlist" v-show="wordListOpen">
          <template v-for="(words, length) in wordsByLength" :key="length">
            <template v-if="words && words.filter(f => !f[1].bonus).length">
              <h6>{{ length }} letters</h6>
              <template v-for="([word, wstate]) in words" :key="word">
                <template v-if="!wstate.bonus">
                  <template v-if="wstate.guessed">
                    <code>{{ word }}</code><br/>
                  </template>
                  <template v-else>
                    <code>{{ wordWithStars(word) }}</code><br />
                  </template>
                </template>
              </template>
            </template>
          </template>
        </div>
        <div class="game-wrap items-center justify-evenly">
          <VierkantleTop v-if="board" :board="board" :partialWord="partialWord" :lastWords="lastWords" />
          <div class="row items-center justify-evenly">
            <div v-if="error">{{ error }}</div>
            <div v-else-if="!board">Loading board...</div>
            <div class="game" v-else>
              <VierkantleBoard
                :board="board"
                @word="word(null, $event)"
                @partialWord="partialWord = $event"
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
import { computed, onMounted, ref, reactive } from 'vue';
import { StorageSerializers, useStorage } from '@vueuse/core';
import { Multiplayer } from 'src/services/multiplayer';
import { TeamStreamServerMessage } from 'src/services/vierkantle';
import VierkantleTop from 'src/components/VierkantleTop.vue';

const board_ = useStorage<Board | undefined>("board", undefined, undefined, { serializer: StorageSerializers.object });

const board = computed(() => {
  return board_.value;
});

const error = ref("");


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

const multiplayerOpen = ref(false);

function toggleMultiplayer() {
  multiplayerOpen.value = !multiplayerOpen.value;
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

const partialWord = ref("");
const lastWords = ref([] as string[]);

function word(who: string | null, word: string) {
  // A word was guessed. If who is null, it was guessed by us, otherwise by
  // someone else.
  // If it was guessed by us, set partialWord to the result, e.g.:
  // - nope: x
  // - bonus: x
  // - x (you already had it)
  // If a new word was discovered, add it to lastWords. If it was guessed by
  // us, do *not* also set it in partialWord.
  const isOurs = !who;
  who = isOurs ? "" : `${who}: `;

  if (isOurs) {
    partialWord.value = "";
  }

  if (word.length < 4) {
    partialWord.value = "te kort: " + word
    return
  }

  const wordInBoard = board.value!.words[word];
  if (!wordInBoard) {
    if (isOurs) {
      partialWord.value = "nope: " + word
    }
    return
  }

  const maybeBonus = wordInBoard.bonus ? "bonus: " : "";

  if (wordInBoard.guessed) {
    if (isOurs) {
      partialWord.value = `${maybeBonus}${word} (had je al)`
    }
    return
  }

  wordInBoard.guessed = true
  lastWords.value.push(`${who}${maybeBonus}${word}`);
  setTimeout(() => {
    lastWords.value.shift();
  }, 4000);

  if (multiplayer.value && isOurs) {
    multiplayer.value.sendWord(word);
  }
}

const playerName = useStorage("name", "");
const token = useStorage("token", "");
const multiplayer = ref<Multiplayer>()
const multiplayerError = ref<string>();

function onMessage(message: TeamStreamServerMessage): void {
  if (message.error) {
    console.log(message.error);
  } else if (message.word) {
    word(message.word.player, message.word.word);
  }
}

const multiplayerUrl = computed(() => {
  console.log(window.location);
  return window.location.origin + "/#" + multiplayer.value?.token;
})

const copyText = ref("klik om te kopiëren");

async function setClipboard(v: string) {
  await navigator.clipboard.writeText(v);
  copyText.value = "gekopieerd!";
  await new Promise((resolve) => setTimeout(resolve, 1000));
  copyText.value = "klik om te kopiëren";
}

onMounted(async () => {
  if (window.location.hash.length > 2) {
    token.value = window.location.hash.substring(1);
    window.location.hash = "";
  }
  if (token.value && !playerName.value) {
    // They have a token, but not a player name, so let's get that first
    setTimeout(() => {
      multiplayerOpen.value = true;
    }, 100);
    return;
  }
  if (playerName.value && token.value) {
    // Try to optionally connect to the same multiplayer team. If it does not
    // succeed, it's probably an old team, just forget the token.
    const m = reactive(new Multiplayer(playerName.value, onMessage, token.value));
    try {
      await m.connect();
      multiplayer.value = m as Multiplayer;
    } catch(e) {
      if (typeof e == "string" && e.includes("team not found")) {
        token.value = "";
      }
    }
  }
})

const otherPlayers = computed(() => {
  if (!multiplayer.value) {
    return "niemand anders";
  }
  const list = multiplayer.value.players.filter((f) => f != multiplayer.value!.name);
  if (list.length == 0) {
    return "niemand anders";
  } else if (list.length == 1) {
    return list[0];
  } else {
    const last = list.pop();
    return list.join(", ") + " en " + last;
  }
});

async function changePlayerName() {
  if (multiplayer.value && playerName.value != multiplayer.value.name) {
    await multiplayer.value.changePlayerName(playerName.value);
  }
}

async function createTeam() {
  multiplayerError.value = undefined;
  if (playerName.value && !multiplayer.value) {
    const m = new Multiplayer(playerName.value, onMessage);
    try {
      await m.connect();
      multiplayer.value = m;
      token.value = m.token;
    } catch(e) {
      if (typeof e === "string") {
        multiplayerError.value = e;
      } else {
        multiplayerError.value = JSON.stringify(e, null, 0);
      }
    }
  }
}

async function joinTeam() {
  multiplayerError.value = undefined;
  if (token.value && playerName.value) {
    const m = new Multiplayer(playerName.value, onMessage, token.value);
    try {
      await m.connect();
      multiplayer.value = m;
    } catch(e) {
      if (typeof e === "string") {
        multiplayerError.value = e;
      } else {
        multiplayerError.value = JSON.stringify(e, null, 0);
      }
    }
  }
}

async function stopMultiplayer() {
  multiplayerError.value = undefined;
  multiplayer.value = undefined;
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

    &:first-of-type {
      margin-top: 0;
    }
  }
}

@media screen and (min-width: 500px) {
  .game-page {
    margin: 12px 48px;
  }

  .wordlist {
    width: 25%;
    height: auto;
    max-height: calc(100vh - 80px);
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
