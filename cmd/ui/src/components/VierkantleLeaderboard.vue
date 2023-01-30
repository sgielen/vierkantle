<template>
  <q-virtual-scroll
    type="table"
    ref="virtualList"
    style="height: 60vh;"
    :items="scoreBoard"
    :virtual-scroll-item-size="48"
    :virtual-scroll-sticky-size-start="48"
    @virtual-scroll="ensureItems"
  >
    <template v-slot:before>
      <thead class="thead-sticky text-left">
        <tr>
          <th>Positie</th>
          <th>Naam</th>
          <th>Woorden/Tijd</th>
        </tr>
      </thead>
    </template>

    <template v-slot:default="{ item: score, index }: { item: GetScoresResponse_Score, index: number }">
      <tr
        :class="{'ourScore': index == ourScore}"
      >
        <td>#{{ index + 1 }}</td>

        <template v-if="!score">
          <td><q-skeleton type="text" /></td>
          <td><q-skeleton type="text" /></td>
        </template>
        <template v-else>
          <td>
            <template v-if="score.name == ''">
              <span class="anonymous">Anoniem</span>
            </template>
            <template v-else>
              {{ score.name }}
            </template>

            <template v-if="score && score.teamSize > 1">
              (+{{ score.teamSize - 1 }})
            </template>
          </td>
          <td>
            {{ score.words }} / {{ formatTime(score.seconds) }}
          </td>
        </template>
      </tr>
    </template>
  </q-virtual-scroll>
</template>

<script setup lang="ts">
import { GetScoresResponse, GetScoresResponse_Score } from 'src/services/vierkantle';
import { ref, reactive, onMounted } from 'vue'
import { createChannel, createClient } from 'nice-grpc-web';
import { VierkantleServiceDefinition, VierkantleServiceClient } from '../services/vierkantle';
import { QVirtualScroll } from 'quasar';

const props = defineProps<{
  backend: string,
  anonymousId: number,
}>();

const fetchAhead = 10;
const ourScore = ref(0);
const scoreBoard = reactive([] as (GetScoresResponse_Score | undefined)[]);
const virtualList = ref<QVirtualScroll>();

onMounted(async () => {
  const scores = await ensureItems({ from: 0, to: fetchAhead });

  if (scores && scores.yourScore) {
    await ensureItems({ from: scores.yourScore - fetchAhead, to: scores.yourScore + fetchAhead});
    virtualList.value?.scrollTo(scores.yourScore, "center-force");
    setTimeout(() => {
      virtualList.value?.scrollTo(scores.yourScore, "center-force");
    }, 100);
  }
});

function getScores(index: number, amount: number): Promise<GetScoresResponse> {
  const channel = createChannel(props.backend);
  const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);
  return client.getScores({
    index,
    amount,
    myAnonymousId: props.anonymousId,
  });
}

function updateScoreboard(res: GetScoresResponse) {
  if (res.yourScore != 0) {
    ourScore.value = res.yourScore;
  }
  Object.entries(res.scores).forEach(([k, v]) => {
    const key = typeof k === "string" ? Number(k) : k;
    scoreBoard[key] = v
  });
  if (scoreBoard.length < res.totalScores) {
    scoreBoard[res.totalScores - 1] = undefined;
  }
}

function formatTime(seconds: number): string {
  const min = Math.floor(seconds / 60);
  const sec = seconds % 60;
  return min + ":" + sec.toFixed(0).padStart(2, "0");
}

async function ensureItems({ from, to }: { from: number, to: number }): Promise<GetScoresResponse | undefined> {
  if (from < 0) {
    from = 0;
  }
  let needFetchFrom: number | undefined;
  for(let i = from; i <= to; ++i) {
    if (!scoreBoard[i]) {
      needFetchFrom = i;
      break;
    }
  }

  if (needFetchFrom !== undefined) {
    const scores = await getScores(needFetchFrom, to - from);
    updateScoreboard(scores);
    return scores;
  }
}
</script>

<style scoped lang="scss">
.thead-sticky tr > * {
  position: sticky;
  opacity: 1;
  z-index: 1;
  background: white;
}

.thead-sticky tr:last-child > * {
  top: 0
}

.ourScore {
  background-color: lightgreen;
}
</style>
