export function c(...classes: (string | undefined)[]) {
  return classes.filter((cn) => cn !== undefined).join(' ');
}
