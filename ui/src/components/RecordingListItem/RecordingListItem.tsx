import { Recording } from '../../clients/api/api.types';
import styles from './RecordingListItem.module.scss';

type Props = {
  recording: Recording;
};

function renderDate(ts: number) {
  return new Date(ts * 1000).toLocaleDateString(undefined, {
    day: '2-digit',
    month: '2-digit',
  });
}

function renderTime(ts: number) {
  return new Date(ts * 1000).toLocaleTimeString(undefined, {
    hour: '2-digit',
    minute: '2-digit',
  });
}

function renderRecordingDatetime(startsAt: number, endsAt: number) {
  return `${renderDate(startsAt)} • ${renderTime(startsAt)} - ${renderTime(
    endsAt
  )}`;
}

function renderTitle(title: string, subtitle?: string) {
  return `${title}${subtitle ? ` (${subtitle})` : ''}`;
}

function RecordingListItem({ recording }: Props) {
  return (
    <div className={styles.RecordingListItem}>
      <span className={styles.title}>
        {renderTitle(recording.title, recording.subtitle)}
      </span>
      <span className={styles.secondary}>
        {renderRecordingDatetime(recording.startsAt, recording.endsAt)}
      </span>
    </div>
  );
}

export default RecordingListItem;
