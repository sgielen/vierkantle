<template>
  <q-header elevated>
    <q-toolbar>
      <q-btn
        class="only-if-small topbar-button-margin"
        dense
        color="secondary"
        icon="spellcheck"
        aria-label="Word list"
        @click="toggleWordList"
      />

      <q-btn
        class="topbar-button-margin"
        dense
        color="secondary"
        icon="leaderboard"
        aria-label="Scorebord"
        @click="toggleLeaderboard"
      />

      <q-btn
        class="topbar-button-margin"
        dense
        color="secondary"
        icon="groups"
        aria-label="Multiplayer"
        @click="toggleMultiplayer"
      />

      <q-btn
        class="topbar-button-margin"
        dense
        color="secondary"
        icon="share"
        aria-label="Delen"
        @click="toggleShare"
      />

      <q-toolbar-title>
        Vierkantle
      </q-toolbar-title>

      <q-toggle v-model="dark" icon="dark_mode" color="secondary" />
    </q-toolbar>
  </q-header>

  <q-page-container>
    <q-page class="game-page">
      <q-dialog v-model="multiplayerOpen">
        <q-card style="width: 650px">
          <q-card-section>
            <span class="text-h6">Multiplayer-modus</span>
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
              <div><span class="text-h6">Scorebord</span></div>
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
              Je bent ingelogd als <strong>{{ username }}</strong>.<br />
              <q-btn @click="logout" dense>Uitloggen</q-btn>
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

      <q-dialog v-model="shareOpen">
        <q-card style="width: 650px">
          <q-card-section>
            <span class="text-h6">Delen</span>
          </q-card-section>
          <q-separator />
          <q-card-section class="q-pa-md">
            <div class="share q-pa-md q-ma-md">{{ shareText }}</div>
            <q-btn icon="share" @click="share">&nbsp;Delen</q-btn>
            <div v-if="isClipboard">Gekopieerd naar klembord!</div>
          </q-card-section>
          <template v-if="board?.madeBy || board?.description">
            <q-separator />
            <q-card-section class="q-pa-md">
              <template v-if="board.madeBy && board.description">
                Het bord van vandaag is gemaakt door <strong>{{ board.madeBy }}</strong> en die
                schrijft erbij:
              </template>
              <template v-else-if="board.madeBy">
                Het bord van vandaag is gemaakt door <strong>{{ board.madeBy }}</strong>.
              </template>
              <template v-if="board.description">
                <div class="q-ma-sm quote">{{ board.description }}</div>
              </template>
            </q-card-section>
          </template>
          <q-separator />
          <q-card-section class="q-pa-md">
            Krijg je er geen genoeg van? Elke dag om 12:00 's middags is er weer
            een nieuwe Vierkantle. Of speel in de tussentijd de <a
            href="https://squaredle.app" target="_blank">Squaredle</a> (Engels),
            waarop Vierkantle is gebaseerd!
          </q-card-section>
        </q-card>
      </q-dialog>

      <q-dialog v-model="preDescriptionOpen" v-if="board?.preDescription">
        <q-card style="width: 650px">
          <q-card-section>
            <span class="text-h6">Speciaal bord</span>
          </q-card-section>
          <q-separator />
          <q-card-section class="q-pa-md">
            Het bord van vandaag is gemaakt door <strong>{{ board.madeBy }}</strong> en die
            schrijft erbij:
            <div class="q-ma-sm quote">{{ board.preDescription }}</div>
            <q-btn color="primary" @click="preDescriptionOpen = false">Starten</q-btn>
          </q-card-section>
        </q-card>
      </q-dialog>

      <q-dialog v-model="registerOpen">
        <q-card style="width: 650px">
          <q-card-section>
            <span class="text-h6">Registreren</span>
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

      <q-dialog v-model="loginOpen">
        <q-card style="width: 650px">
          <template v-if="loginStep == 1">
            <q-card-section>
              <div class="text-h6">Inloggen</div>
              Je kan opnieuw inloggen op je account op twee manieren.
            </q-card-section>
            <q-separator />
            <q-card-section class="q-pa-md">
              Weet je je gebruikersnaam nog?<br />
              <q-input dense outlined v-model="loginUsername" /><br />
              <q-btn @click="startLoginUsername" :disabled="!loginUsername.trim().length" color="primary">Start inloggen</q-btn>
            </q-card-section>
            <q-separator />
            <q-card-section class="q-pa-md">
              Zo niet, gebruik dan je e-mailadres. Dit werkt uiteraard alleen als je een e-mailadres
              hebt ingevuld bij het registreren.<br />
              <q-input dense outlined v-model="loginEmail" /><br />
              <q-btn @click="startLoginEmail" :disabled="!loginEmail.trim().length" color="primary">Start inloggen</q-btn>
            </q-card-section>
          </template>
          <template v-else-if="loginStep == 2">
            <q-card-section>
              <div class="text-h6">Inloggen</div>
              Je hebt een e-mail ontvangen met een link. Open deze link om weer in te loggen.
            </q-card-section>
          </template>
          <template v-else-if="loginStep == 3">
            <q-card-section>
              <div class="text-h6">Inloggen</div>
              Je bent succesvol ingelogd als {{ username }}!
            </q-card-section>
          </template>
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
            <div class="game" ref="gameRef" v-else>
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
import { computed, onMounted, ref, reactive, watchEffect, watch } from 'vue';
import { StorageSerializers, useStorage, useMediaQuery } from '@vueuse/core';
import { Multiplayer } from 'src/services/multiplayer';
import VierkantleTop from 'src/components/VierkantleTop.vue';
import { createChannel, createClient } from 'nice-grpc-web';
import { VierkantleServiceDefinition, VierkantleServiceClient, TeamStreamServerMessage } from '../services/vierkantle';
import WordList from 'src/components/WordList.vue';
import VierkantleLeaderboard from 'src/components/VierkantleLeaderboard.vue';
import { errorToString } from 'src/services/errors';
import { confetti, ConfettiOptions } from 'tsparticles-confetti';
import { useQuasar } from 'quasar';
import { useWhoami } from 'src/services/whoami';

const $q = useQuasar();
const dark = computed({
  get() {
    return $q.dark.isActive;
  },
  set(status: boolean) {
    $q.dark.set(status);
  },
});

const board_ = useStorage<Board | null>("board", null, undefined, { serializer: StorageSerializers.object });
const anonymousId = useStorage("anonymousId", Math.floor(Math.random() * 4294967295 /* UINT32_MAX */));
const seconds = useStorage("seconds", 0);
const enableSubmitScores = useStorage("enableSubmitScores", true);

const board = computed(() => {
  return board_.value;
});
const boardName = computed(() => {
  return board.value?.name ?? "";
})

const timeSpent = computed(() => {
  let sec = Math.floor(seconds.value);
  const min = Math.floor(sec / 60);
  sec = sec % 60;
  return min + ":" + sec.toFixed(0).padStart(2, "0");
});

const error = ref("");
const backendAddress = window.location.origin + "/api";

const channel = createChannel(backendAddress);
const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);

const { username, updateWhoami } = useWhoami(client);

const isWideScreen = useMediaQuery(`(min-width: 500px)`);

const wordListDialogVisible = computed(() => {
  // On wide screens, the word list dialog is never visible.
  if (isWideScreen.value) {
    return false;
  }

  // Otherwise, the word list dialog is visible depending on whether it's opened.
  return wordListOpen.value;
})

const anyDialogOpen = computed(() => {
  return wordListDialogVisible.value || multiplayerOpen.value || leaderboardOpen.value || shareOpen.value || preDescriptionOpen.value || loginOpen.value || registerOpen.value;
})

onMounted(async () => {
  seconds.value ??= 0;
  setInterval(() => {
    if (document.hasFocus() && !boardIsDone.value && !anyDialogOpen.value) {
      seconds.value! += 0.1;
    }
  }, 100)

  // Download the board to check if there is a new one available
  try {
    const boardResponse = await client.getBoard({
      timezoneOffsetMinutes: new Date().getTimezoneOffset(),
    });
    const board = JSON.parse(new TextDecoder().decode(boardResponse.board));
    board.name = boardResponse.name;
    if (!board_.value || boardLetters(board) != boardLetters(board_.value)) {
      board_.value = board;
      seconds.value = 0;
    }
  } catch(e) {
    error.value = errorToString(e);
  }

  if (board.value?.preDescription) {
    preDescriptionOpen.value = true;
  }

  await updateWhoami();
});

const preDescriptionOpen = ref(false);

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

const shareOpen = ref(false);
const shareText = computed(() => {
  let text = "https://vierkantle.nl/ " + boardName.value + "\n";
  text += wordsGuessed.value + "/" + wordsTotal.value + " woorden in " + timeSpent.value;
  return text;
});
const isClipboard = ref(false);
async function share() {
  try {
    navigator.clipboard.writeText(shareText.value);
    isClipboard.value = true;
    setTimeout(() => isClipboard.value = false, 5000);
  } catch(e) {}
  try {
    await navigator.share({
      text: shareText.value,
    });
  } catch(e) {}
}
function toggleShare() {
  shareOpen.value = !shareOpen.value;
}
const gameRef = ref<HTMLDivElement>();
function boardDoneAnimation() {
  const count = 200;

  const gameRect = gameRef.value!.getBoundingClientRect();
  const gameMidX = gameRect.left + (gameRect.width / 2);
  const gameBotY = gameRect.bottom - 100;
  const defaults: ConfettiOptions = {
    origin: {
      x: gameMidX / window.innerWidth,
      y: gameBotY / window.innerHeight,
    },
  };

  const fire = (particleRatio: number, opts: ConfettiOptions) => {
    confetti(
      Object.assign({}, defaults, opts, {
        particleCount: Math.floor(count * particleRatio),
      })
    );
  }

  const fireworks = () => {
    const factor = Math.random();
    fire(0.25, {
      spread: 10 + factor * 30,
      startVelocity: 55,
    });

    fire(0.2, {
      spread: 60,
    });

    fire(0.35, {
      spread: 10 + factor * 120,
      decay: 0.91,
      scalar: 0.8,
    });

    fire(0.1, {
      spread: 120,
      startVelocity: 25,
      decay: 0.92,
      scalar: 1.2,
    });

    fire(0.1, {
      spread: 20 + factor * 250,
      startVelocity: 45,
    });
  }

  fireworks();

  let fireworksCount = 0;

  const interval = setInterval(() => {
    if (fireworksCount == 3) {
      clearInterval(interval);
    }
    fireworks();
    fireworksCount++;
  }, 1200);

  setTimeout(() => shareOpen.value = true, 3000);
}

const wordsTotal = computed(() => {
  return Object.values(board.value?.words ?? []).filter((w) => !w.bonus).length;
});

const wordsGuessed = computed(() => {
  if (!board.value) {
    return 0;
  }

  return Object.values(board.value.words).filter((w) => !w.bonus && w.guessed).length;
});

const boardIsDone = computed(() => {
  return Object.values(board.value?.words ?? {}).find((f) => !f.guessed && !f.bonus) === undefined;
});

watch(boardIsDone, (newV, oldV) => {
  if (!oldV && newV) {
    boardDoneAnimation();
  }
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
const loginStep = ref<1 | 2 | 3>(1);
const loginUsername = ref<string>("");
const loginEmail = ref<string>("");

function startLogin() {
  leaderboardOpen.value = false;
  registerOpen.value = false;
  loginOpen.value = true;
  loginStep.value = 1;
  error.value = "";
}

async function startLoginUsername() {
  try {
    error.value = "";
    const response = await client.startLogin({
      username: loginUsername.value.trim(),
      urlPrefix: window.location.origin + "/#l=",
    })
    if (response.found) {
      loginStep.value = 2;
    } else {
      error.value = "Die gebruikersnaam bestaat niet, of heeft geen bijbehorend e-mailadres."
    }
  } catch(e) {
    error.value = errorToString(e);
  }
}

async function startLoginEmail() {
  try {
    error.value = "";
    const response = await client.startLogin({
      email: loginEmail.value.trim(),
      urlPrefix: window.location.origin + "/#l=",
    })
    if (response.found) {
      loginStep.value = 2;
    } else {
      error.value = "Dat e-mailadres is niet gevonden.";
    }
  } catch(e) {
    error.value = errorToString(e);
  }
}

onMounted(async () => {
  if (window.location.hash.startsWith("#l=")) {
    const token = window.location.hash.substring(3);
    window.location.hash = "";
    try {
      error.value = "";
      const response = await client.finishLogin({
        token,
      });
      username.value = response.username;
      loginOpen.value = true;
      loginStep.value = 3;

      // Also perform a single SubmitScore so that the user ID of the score is updated
      // if there already was one
      await updateScore();
    } catch(e) {
      error.value = errorToString(e);
      loginOpen.value = true;
      loginStep.value = 2;
    }
  }
});

const registerOpen = ref(false);
const registerUsername = ref<string>("");
const registerEmail = ref<string>("");

function startRegister() {
  leaderboardOpen.value = false;
  loginOpen.value = false;
  registerOpen.value = true;
  error.value = "";
}

async function register() {
  try {
    error.value = "";
    await client.register({
      username: registerUsername.value.trim(),
      email: registerEmail.value.trim(),
    })
    await updateWhoami();
    loginOpen.value = false;
    leaderboardOpen.value = true;

    // Also perform a single SubmitScore so that the user ID of the score is updated
    // if there already was one
    await updateScore();
  } catch(e) {
    error.value = errorToString(e);
  }
}

async function logout() {
  try {
    error.value = "";
    await client.logout({});
    username.value = undefined;
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
});

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

.share {
  white-space: pre-wrap;
  background-color: #ccc;
  border: 1px solid black;
}

.quote {
  white-space: pre-wrap;
  font-style: italic;
  border-left: 4px solid var(--q-primary);
  padding-left: 6px;
}

body.body--dark {
  header {
    background-color: #0d4174;
  }

  .wordlist {
    background: var(--q-dark-page);
    border: 1px solid white;
  }

  .share {
    background-color: #0d4174;
    border: 1px solid white;
  }
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

.topbar-button-margin {
  margin-left: 8px;
  margin-right: 8px;
}

@media screen and (max-width: 400px) {
  /* Smaller top bar buttons, so the word "Vierkantle" isn't shortened */
  .topbar-button-margin {
    margin-left: 1px;
    margin-right: 1px;
  }

  .q-toolbar__title {
    padding: 0;
    padding-left: 6px;
  }
}
</style>
