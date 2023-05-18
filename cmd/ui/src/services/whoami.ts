import { Ref, ref, watchEffect } from 'vue';
import { VierkantleServiceClient, WhoamiResponse } from './vierkantle';

export function useWhoami(client: VierkantleServiceClient) {
  const whoami = ref<WhoamiResponse>();
  const curriedSendWhoami = () => {
    return sendWhoami(client, whoami);
  };

  // Keep checking whether we're logged in (this also refreshes the cookie)
  setInterval(curriedSendWhoami, 60 * 1000);

  const username = ref<string>();
  watchEffect(() => {
    username.value = whoami.value?.username;
  });

  return { username, whoami, updateWhoami: curriedSendWhoami }
}

async function sendWhoami(client: VierkantleServiceClient, response: Ref<WhoamiResponse | undefined>, critical?: boolean) {
  try {
    response.value = await client.whoami({});
  } catch(e) {
    if (critical) {
      response.value = undefined;
    } else {
      // This may be a temporary error, or a race condition with the
      // finishLogin() RPC. Try whoami RPC again, but with critical=true, before
      // setting username.value = undefined.
      setTimeout(() => sendWhoami(client, response, true), 500);
    }
  }
}
