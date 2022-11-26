<template>
  <q-page>
    <div class="row items-center justify-evenly">{{ wordsRemaining }} words remaining</div>
    <div class="row items-center justify-evenly">{{ wordMessage }}</div>
    <div class="row items-center justify-evenly">
      <div v-if="error">{{ error }}</div>
      <div v-else-if="!board">Loading board...</div>
      <div class="game" v-else>
        <VierkantleBoard
          :board="board"
          @word="word"
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
import { VierkantleServiceDefinition, VierkantleServiceClient, HelloStreamRequest } from '../services/vierkantle';
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

async function testRpc() {
  console.log("trying normal rpc...");
  try {
    const response = await client.hello({
      m: "world",
    });
    console.log("succeeded with: ", response.m);
  } catch(e) {
    console.log("failed with: ", e);
  }
  console.log("done with normal rpc...");

  class AsyncHelloStreamRequests {
    async *[Symbol.asyncIterator]() {
      for (let i = 0; i < 5; i++) {
        await new Promise((resolve) => setTimeout(resolve, 1000));
        yield HelloStreamRequest.fromPartial({
          m: "world " + i,
        });
      }
    }
  }

  console.log("trying stream rpc...");
  try {
    const stream = client.helloStream(new AsyncHelloStreamRequests)
    for await (const response of stream) {
      console.log("have message", response);
    }
  } catch(e) {
    console.log("failed with: ", e);
  }
  console.log("done with stream rpc...");

}

function word(word: string) {
  testRpc();

  if (word.length < 4) {
    wordMessage.value = "te kort: " + word
    return
  }
  const wordInBoard = board.value!.words[word];
  if (!wordInBoard) {
    wordMessage.value = "nope: " + word
    return
  }
  if (wordInBoard.bonus) {
    wordMessage.value = "bonus: " + word
  } else {
    wordMessage.value = word
  }
  if (wordInBoard.guessed) {
    wordMessage.value += " (had je al)"
  }
  wordInBoard.guessed = true
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
