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
        icon="leaderboard"
        aria-label="Scorebord"
        @click="toggleLeaderboard"
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
    <q-page class="game-page">
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
              <q-input dense outlined label="Je naam" v-model="multiplayerName" @update:model-value="changeMultiplayerName" />
            </q-card-section>
            <q-separator />
            <q-card-section>
              <q-input dense outlined v-model="multiplayerToken" label="Code" />
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

      <q-dialog v-model="leaderboardOpen">
        <q-card style="width: 650px">
          <q-card-section>
            <div class="row justify-between">
              <div><p class="text-h6">Scorebord</p></div>
              <div>
                <q-toggle
                  v-model="enableSubmitScores"
                  label="Stuur scores"
                />
              </div>
            </div>
          </q-card-section>
          <q-separator />
          <q-card-section class="q-pa-md">
            <template v-if="username">
              Je bent ingelogd als <strong>{{ username }}</strong>.
            </template>
            <template v-else>
              Je bent niet ingelogd, dus je score zal getoond worden als <em>Anoniem</em>.<br />
              <q-btn @click="startLogin" dense color="primary">Inloggen</q-btn>
              <span class="q-mx-sm">of</span>
              <q-btn @click="startRegister" dense color="primary">je naam registreren?</q-btn>
            </template>
          </q-card-section>
          <q-separator />
          <q-card-section>
            <VierkantleLeaderboard
              :backend="backendAddress"
              :anonymousId="anonymousId"
              :boardName="boardName"
            />
          </q-card-section>
        </q-card>
      </q-dialog>

      <q-dialog v-model="loginOpen">
        <q-card style="width: 650px">
          <q-card-section>
            <p class="text-h6">Registreren</p>
          </q-card-section>
          <q-separator />
          <q-card-section class="q-pa-md">
            <span style="color: red">*</span> Vul je gebruikersnaam in. Deze is te zien voor andere gebruikers op het scorebord:<br />
            <q-input dense outlined v-model="registerUsername" /><br />
            Optioneel: vul je e-mailadres in. Deze gebruik je om weer toegang te krijgen tot je naam op een ander
            apparaat of wanneer je automatisch bent uitgelogd.<br />
            <q-input dense outlined v-model="registerEmail" /><br />
            <q-btn @click="register" :disabled="!registerUsername.trim().length" color="primary">Registreren</q-btn>
          </q-card-section>
          <template v-if="error">
            <q-separator />
            <q-card-section class="q-pa-md"><strong>{{ error }}</strong></q-card-section>
          </template>
        </q-card>
      </q-dialog>

      <div class="row items-between q-mt-md q-mx-xl">
        <div class="wordlist" v-show="wordListOpen">
          <WordList v-if="board" :words="board.words" />
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
          <div class="row justify-between">
            <div class="q-px-xl">{{ timeSpent }}</div>
            <div></div>
          </div>
        </div>
      </div>
    </q-page>
  </q-page-container>
</template>

<script setup lang="ts">
import VierkantleBoard from 'components/VierkantleBoard.vue';
import { Board } from 'src/components/models';
import { computed, onMounted, ref, reactive, watchEffect } from 'vue';
import { StorageSerializers, useStorage } from '@vueuse/core';
import { Multiplayer } from 'src/services/multiplayer';
import VierkantleTop from 'src/components/VierkantleTop.vue';
import { createChannel, createClient } from 'nice-grpc-web';
import { VierkantleServiceDefinition, VierkantleServiceClient, TeamStreamServerMessage } from '../services/vierkantle';
import WordList from 'src/components/WordList.vue';
import VierkantleLeaderboard from 'src/components/VierkantleLeaderboard.vue';
import { errorToString } from 'src/services/errors';

const board_ = useStorage<Board | undefined>("board", undefined, undefined, { serializer: StorageSerializers.object });
const anonymousId = useStorage("anonymousId", Math.floor(Math.random() * 4294967295 /* UINT32_MAX */));
const seconds = useStorage("seconds", 0);
const enableSubmitScores = useStorage("enableSubmitScores", true);

const boardName = ref("");
const board = computed(() => {
  return board_.value;
});

const timeSpent = computed(() => {
  const min = Math.floor(seconds.value / 60);
  const sec = seconds.value % 60;
  return min + ":" + sec.toFixed(0).padStart(2, "0");
});

const error = ref("");
const backendAddress = window.location.origin + "/api";

const channel = createChannel(backendAddress);
const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);

const username = ref<string>();

onMounted(async () => {
  seconds.value ??= 0;
  setInterval(() => {
    if (document.hasFocus() && !boardIsDone.value) {
      seconds.value! += 1;
    }
  }, 1000)

  // Download the board to check if there is a new one available
  try {
    const boardResponse = await client.getBoard({
      timezoneOffsetMinutes: new Date().getTimezoneOffset(),
    });
    boardName.value = boardResponse.name ?? "";
    const board = JSON.parse(new TextDecoder().decode(boardResponse.board));
    if (!board_.value || boardLetters(board) != boardLetters(board_.value)) {
      board_.value = board;
      seconds.value = 0;
    }
  } catch(e) {
    error.value = errorToString(e);
  }

  // Keep checking whether we're logged in (this also refreshes the cookie)
  setInterval(sendWhoami, 60 * 1000);
  await sendWhoami();
});

async function sendWhoami() {
  try {
    const whoami = await client.whoami({});
    username.value = whoami.username;
  } catch(e) {
    username.value = undefined;
  }
}

const wordListOpen = ref(false);

function toggleWordList() {
  wordListOpen.value = !wordListOpen.value
}

const multiplayerOpen = ref(false);

function toggleMultiplayer() {
  multiplayerOpen.value = !multiplayerOpen.value;
}

const leaderboardOpen = ref(false);

function toggleLeaderboard() {
  leaderboardOpen.value = !leaderboardOpen.value;
}

const boardIsDone = computed(() => {
  return Object.values(board.value?.words ?? {}).find((f) => !f.guessed && !f.bonus) === undefined;
});

function boardLetters(b: Board): string {
  return b.cells.reduce((p, c) => { return [...p, c.join("")] }).join("");
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

  updateScore();

  if (multiplayer.value && isOurs) {
    multiplayer.value.sendWord(word);
  }
}

const loginOpen = ref(false);
const registerUsername = ref<string>("");
const registerEmail = ref<string>("");

function startLogin() {
  // TODO: implement login
  startRegister();
}

function startRegister() {
  leaderboardOpen.value = false;
  loginOpen.value = true;
  error.value = "";
}

async function register() {
  try {
    error.value = "";
    await client.register({
      username: registerUsername.value.trim(),
      email: registerEmail.value.trim(),
    })
    await sendWhoami();
    loginOpen.value = false;
    leaderboardOpen.value = true;

    // Also perform a single SubmitScore so that the user ID of the score is updated
    // if there already was one
    await updateScore();
  } catch(e) {
    error.value = errorToString(e);
  }
}

const multiplayerName = useStorage("name", "");
watchEffect(() => {
  if (username.value) {
    multiplayerName.value = username.value;
    changeMultiplayerName();
  }
});
const multiplayerToken = useStorage("multiplayerToken", "");
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
  return window.location.origin + "/#m=" + multiplayer.value?.token;
})

const copyText = ref("klik om te kopiëren");

async function setClipboard(v: string) {
  await navigator.clipboard.writeText(v);
  copyText.value = "gekopieerd!";
  await new Promise((resolve) => setTimeout(resolve, 1000));
  copyText.value = "klik om te kopiëren";
}

function newMultiplayer(token?: string): Multiplayer {
  return reactive(new Multiplayer(backendAddress, multiplayerName.value, onMessage, token)) as Multiplayer;
}

onMounted(async () => {
  if (window.location.hash.startsWith("#m=")) {
    multiplayerToken.value = window.location.hash.substring(3);
    window.location.hash = "";
  }
  if (multiplayerToken.value && !multiplayerName.value) {
    // They have a token, but not a player name, so let's get that first
    setTimeout(() => {
      multiplayerOpen.value = true;
    }, 100);
    return;
  }
  if (multiplayerName.value && multiplayerToken.value) {
    // Try to optionally connect to the same multiplayer team. If it does not
    // succeed, it's probably an old team, just forget the token.
    const m = newMultiplayer(multiplayerToken.value);
    try {
      await m.connect();
      multiplayer.value = m as Multiplayer;
    } catch(e) {
      if (typeof e == "string" && e.includes("team not found")) {
        multiplayerToken.value = "";
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

async function changeMultiplayerName() {
  if (multiplayer.value && multiplayerName.value != multiplayer.value.name) {
    await multiplayer.value.changePlayerName(multiplayerName.value);
  }
}

async function createTeam() {
  multiplayerError.value = undefined;
  if (multiplayerName.value && !multiplayer.value) {
    const m = newMultiplayer();
    try {
      await m.connect();
      multiplayer.value = m;
      multiplayerToken.value = m.token;
    } catch(e) {
      multiplayerError.value = errorToString(e);
    }
  }
}

async function joinTeam() {
  multiplayerError.value = undefined;
  if (multiplayerToken.value && multiplayerName.value) {
    const m = newMultiplayer(multiplayerToken.value);
    try {
      await m.connect();
      multiplayer.value = m;
    } catch(e) {
      multiplayerError.value = errorToString(e);
    }
  }
}

async function stopMultiplayer() {
  multiplayer.value?.disconnect();
  multiplayerError.value = undefined;
  multiplayer.value = undefined;
}

async function updateScore() {
  if (!enableSubmitScores.value) {
    return;
  }
  try {
    await client.submitScore({
      anonymousId: anonymousId.value,
      teamSize: multiplayer.value ? multiplayer.value.players.length : 0,
      words: Object.values(board.value!.words).filter((f) => f.guessed && !f.bonus).length,
      seconds: seconds.value,
      boardName: boardName.value,
    });
  } catch(e) {
    // ignore error
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
  overflow: scroll;
  padding: 4px;
}

@media screen and (min-width: 500px) {
  .game-page {
    padding: 12px;
  }

  .wordlist {
    width: 25%;
    height: auto;
    max-height: calc(100svh - 128px);
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
    height: calc(100svh - 80px);
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
