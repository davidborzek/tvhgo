import { RecIcon } from '@/assets';
import { c } from '@/utils/classNames';
import styles from './EventRecordButton.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  dvrUuid?: string;
  pending?: boolean;
  onClick: () => void;
};

function EventRecordButton({ dvrUuid, pending, onClick }: Props) {
  const { t } = useTranslation();

  const getText = () => {
    if (dvrUuid) {
      return t('modify_recording');
    }
    return t('record');
  };

  return (
    <button
      className={c(styles.EventRecordButton, pending ? styles.pending : '')}
      onClick={onClick}
    >
      <RecIcon className={styles.icon} />
      {getText()}
    </button>
  );
}

export default EventRecordButton;
