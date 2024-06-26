import Checkbox from '@/components/common/checkbox/Checkbox';
import { Recording } from '@/clients/api/api.types';
import { TestIds } from '@/__test__/ids';
import styles from './RecordingListItem.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  recording: Recording;
  onClick: () => void;
  onSelection: (selected: boolean) => void;
  selected: boolean;
};

function renderTitle(title: string, subtitle?: string) {
  return `${title}${subtitle ? ` (${subtitle})` : ''}`;
}

function RecordingListItem({
  recording,
  selected,
  onClick,
  onSelection,
}: Props) {
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
    <div className={styles.RecordingListItem} tabIndex={0}>
      <div className={styles.link} onClick={onClick}>
        <span className={styles.title}>
          {renderTitle(recording.title, recording.subtitle)}
        </span>
        <div className={styles.secondary}>
          {renderRecIndicator()}
          <span>
            {recording.channelName} |{' '}
            {t('event_datetime', { event: recording })}
          </span>
        </div>
      </div>

      <Checkbox
        onChange={(checked) => onSelection(checked)}
        checked={selected}
        testId={TestIds.SELECT_RECORDING_CHECKBOX}
      />
    </div>
  );
}

export default RecordingListItem;
