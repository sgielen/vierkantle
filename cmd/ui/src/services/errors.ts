// eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/explicit-module-boundary-types
export function errorToString(e: any): string {
  console.log(JSON.stringify(e, null, 2));
  if (typeof e === 'string') {
    return e;
  }
  const errstr = e.details || e.error || e.message;
  if (errstr) {
    return errstr;
  }
  return e.code || e;
}
