export function parseTime(ts: number): string {
  return new Date(ts * 1000).toLocaleTimeString(undefined, {
    hour: '2-digit',
    minute: '2-digit',
    hourCycle: 'h23',
  });
}

export function parseDate(ts: number) {
  return new Date(ts * 1000).toLocaleDateString(undefined, {
    day: '2-digit',
    month: '2-digit',
  });
}

export function parseDatetime(startsAt: number, endsAt: number) {
  return `${parseDate(startsAt)} • ${parseTime(startsAt)} - ${parseTime(
    endsAt
  )}`;
}
