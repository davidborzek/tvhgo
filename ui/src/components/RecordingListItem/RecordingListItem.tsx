import { useTranslation } from 'react-i18next';
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
  const { t } = useTranslation();

  const renderRecIndicator = () => {
    if (recording.status === 'recording') {
      return (
        <div
          title={t('recording_running') || ''}
          className={styles.recIndicator}
        />
      );
    }
  };

  return (
    <div className={styles.RecordingListItem} onClick={onClick} tabIndex={0}>
      <span className={styles.title}>
        {renderTitle(recording.title, recording.subtitle)}
      </span>
      <div className={styles.secondary}>
        {renderRecIndicator()}
        <span>{parseDatetime(recording.startsAt, recording.endsAt)}</span>
      </div>
    </div>
  );
}

export default RecordingListItem;
