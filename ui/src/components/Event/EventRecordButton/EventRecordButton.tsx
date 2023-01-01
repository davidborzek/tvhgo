import { useTranslation } from 'react-i18next';
import { RecIcon } from '../../../assets';
import { c } from '../../../utils/classNames';
import styles from './EventRecordButton.module.scss';

type Props = {
  dvrState?: string;
  pending?: boolean;
  onClick: () => void;
};

function EventRecordButton({ dvrState, pending, onClick}: Props) {
  const { t } = useTranslation('event');

  const getText = () => {
    switch (dvrState) {
      case 'scheduled':
        return t('cancel');
      case 'recording':
        return t('stop');
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
