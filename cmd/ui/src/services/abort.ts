import { ref } from 'vue';

// This function primarily returns a newUniqueCall function.  When called, it
// creates and returns a new abortSignal. Then, when called again after this, it
// first aborts the previously returned abourtSignal, then creates and returns a
// new one.  This way, you can call newUniqueCall() before every RPC, passing
// the abortSignal into the RPC; when a new RPC should be called, this first
// aborts the old one, before starting the new one.
//
// You can call useUniqueCall() multiple times and store the newUniqueCall in
// another variable to allow various sets of unique calls.
export function useUniqueCall() {
  const currentAbortController = ref<AbortController>();
  const newUniqueCall = (reason?: unknown) => {
    currentAbortController.value?.abort(reason);
    currentAbortController.value = new AbortController();
    return currentAbortController.value.signal;
  }

  return { newUniqueCall };
}

export function isAbortError(error: unknown): boolean {
  return error instanceof Object && 'name' in error && error.name === "AbortError";
}
