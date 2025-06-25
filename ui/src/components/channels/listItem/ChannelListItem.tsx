import Button from '@/components/common/button/Button';
import { EpgEvent } from '@/clients/api/api.types';
import Image from '@/components/common/image/Image';
import styles from './ChannelListItem.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  event: EpgEvent;
  onClick: (id: string) => void;
  onWatch: (event: EpgEvent) => void;
};

function ChannelListItem({ event, onClick, onWatch }: Props) {
  const { t } = useTranslation();

  const handleWatchClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onWatch(event);
  };

  return (
    <div className={styles.channel} onClick={() => onClick(event.channelId)}>
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
      <div className={styles.actions}>
        <Button
          onClick={handleWatchClick}
          style="blue"
          size="small"
          label={t('watch')}
        />
      </div>
    </div>
  );
}

export default ChannelListItem;
