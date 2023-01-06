import { Recording } from '../../clients/api/api.types';
import { parseDatetime } from '../../utils/time';
import styles from './RecordingListItem.module.scss';

type Props = {
  recording: Recording;
  onClick: () => void;
};

function renderTitle(title: string, subtitle?: string) {
  return `${title}${subtitle ? ` (${subtitle})` : ''}`;
}

function RecordingListItem({ recording, onClick }: Props) {
  return (
    <div className={styles.RecordingListItem} onClick={onClick} tabIndex={0}>
      <span className={styles.title}>
        {renderTitle(recording.title, recording.subtitle)}
      </span>
      <span className={styles.secondary}>
        {parseDatetime(recording.startsAt, recording.endsAt)}
      </span>
    </div>
  );
}

export default RecordingListItem;
