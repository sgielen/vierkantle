<template>
  <q-card style="width: 650px">
    <q-card-section>
      <span class="text-h6">Wachtrij</span>
    </q-card-section>
    <q-separator />
    <q-card-section class="q-pa-md">
      <template v-if="error">
        <span style="color: red">{{ error }}</span>
      </template>
      <q-table
        :rows="rows"
        row-key="id"
        :columns="columns"
        :rows-per-page-options="[1000]"
        :hide-pagination="true"
        selection="multiple"
        v-model:selected="selected"
        @row-click="(_, r) => open(r.id)"
      />
    </q-card-section>
    <q-separator />
    <q-card-section>
      <div class="row q-gutter-sm">
        <q-input dense outlined style="max-width: 160px;" v-model="newName" />
        <q-btn color="secondary" :disable="selected.length != 1" @click="rename">Rename board</q-btn>
        <q-btn color="negative" :disable="selected.length != 1" @click="replace">Rename and replace board</q-btn>
      </div>
    </q-card-section>
    <q-separator />
    <q-card-section>
      <div class="row q-gutter-sm">
        <q-btn color="primary" :disable="!selected.length" @click="downloadSelection">Download {{ selected.length }} bord(en)</q-btn>
        <q-btn color="negative" :disable="!selected.length" @click="removeSelection">Haal {{ selected.length }} bord(en) uit wachtrij</q-btn>
      </div>
    </q-card-section>
  </q-card>
</template>

<script setup lang="ts">
import { createChannel, createClient } from 'nice-grpc-web';
import { QTableProps } from 'quasar';
import { errorToString } from 'src/services/errors';
import { GetBoardFromQueueResponse, ListBoardQueueResponse_Board, VierkantleServiceClient, VierkantleServiceDefinition } from 'src/services/vierkantle';
import { onMounted, ref } from 'vue';

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import { Writer } from '@transcend-io/conflux';

const props = defineProps<{
  backend: string,
  currentBoardFn: () => Uint8Array,
}>();

const emit = defineEmits<{
  (event: 'open', board: GetBoardFromQueueResponse): void
}>();

const channel = createChannel(props.backend);
const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);

const error = ref<string>();

const selected = ref([] as ListBoardQueueResponse_Board[]);

const rows = ref([] as ListBoardQueueResponse_Board[]);

onMounted(async () => {
  await listBoardQueue();
});

async function listBoardQueue() {
  try {
    const boards = await client.listBoardQueue({});
    rows.value = boards.boards;
  } catch(e) {
    error.value = errorToString(e);
  }
}

async function open(id: number) {
  try {
    error.value = "";
    const board = await client.getBoardFromQueue({ id });
    emit("open", board);
  } catch(e) {
    error.value = errorToString(e);
  }
}

const newName = ref("");

async function rename() {
  try {
    error.value = "";
    const board = await client.getBoardFromQueue({ id: selected.value[0].id });
    await client.updateBoardInQueue({
      id: selected.value[0].id,
      boardName: newName.value,
      board: board.board,
    });
    await listBoardQueue();
  } catch(e) {
    error.value = errorToString(e);
  }
}

async function replace() {
  try {
    error.value = "";
    await client.updateBoardInQueue({
      id: selected.value[0].id,
      boardName: newName.value || selected.value[0].boardName,
      board: props.currentBoardFn(),
    });
    await listBoardQueue();
  } catch(e) {
    error.value = errorToString(e);
  }
}

async function downloadSelection() {
  let file: File;
  if (selected.value.length == 0) {
    return;
  } else if (selected.value.length == 1) {
    const board = await client.getBoardFromQueue({ id: selected.value[0].id });
    const blob = await new Response(board.board).blob();
    file = new File([blob], selected.value[0].boardName + ".json");
  } else {
    interface ZipFile {
      name: string;
      lastModified: Date;
      stream: () => ReadableStream<Uint8Array>;
    }
    const { readable, writable } = new Writer() as {
      readable: ReadableStream<Uint8Array>,
      writable: WritableStream<ZipFile>
    };
    try {
      error.value = "";
      const writer = writable.getWriter();

      for(let i = 0; i < selected.value.length; ++i) {
        const board = await client.getBoardFromQueue({ id: selected.value[i].id });
        writer.write({
          name: board.boardName + ".json",
          lastModified: new Date(0),
          stream: () => new Response(board.board).body!,
        });
      }
      const response = new Response(readable);
      const [blob] = await Promise.all([response.blob(), writer.close()]);

      file = new File([blob], "boards.zip");
    } catch(e) {
      console.log(e);
      error.value = errorToString(e);
      return;
    }
  }

  const link = document.createElement("a");
  link.style.display = "none";
  link.href = URL.createObjectURL(file);
  link.download = file.name;

  document.body.appendChild(link);
  link.click();

  // To make this work on Firefox we need to wait a little while before removing
  // it.
  setTimeout(() => {
    URL.revokeObjectURL(link.href);
    link.parentNode?.removeChild(link);
  }, 0);
}

async function removeSelection() {
  try {
    error.value = "";

    await client.removeBoardsFromQueue({
      ids: selected.value.map((b) => b.id),
    })
    await listBoardQueue();
  } catch(e) {
    error.value = errorToString(e);
  }
}

const columns: QTableProps['columns'] = [
  {
    name: "id",
    label: "ID",
    field: "id",
  },
  {
    name: "boardName",
    label: "Name",
    field: "boardName",

  },
  {
    name: "username",
    label: "Author",
    field: "username",
  },
];
</script>
