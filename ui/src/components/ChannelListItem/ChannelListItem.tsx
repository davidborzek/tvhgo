import styles from './ChannelListItem.module.scss';

import Image from '../Image/Image';

import { EpgEvent } from '../../clients/api/api.types';
import { useTranslation } from 'react-i18next';

type Props = {
  event: EpgEvent;
};

function ChannelListItem({ event }: Props) {
  const { t } = useTranslation();

  return (
    <div className={styles.channel}>
      <div className={styles.piconContainer}>
        <Image
          className={styles.picon}
          src={`/api/picon/${event.piconId}`}
          alt=""
          title={event.channelName}
        />
      </div>
      <div className={styles.event}>
        <span className={styles.channelName}>{event.channelName}</span>
        <span className={styles.eventTitle}>{event.title}</span>
        <span className={styles.eventTitle}>{t('event_time', { event })}</span>
      </div>
    </div>
  );
}

export default ChannelListItem;
