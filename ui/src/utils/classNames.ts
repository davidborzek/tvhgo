export function c(...classes: (string | undefined | null)[]) {
  return classes.filter((cn) => !!cn).join(' ');
}
