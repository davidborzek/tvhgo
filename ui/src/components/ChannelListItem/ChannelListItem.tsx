import styles from './ChannelListItem.module.scss';

import Image from '../Image/Image';

import { EpgEvent } from '../../clients/api/api.types';

type Props = {
  event: EpgEvent;
};

function parseTime(ts: number): string {
  return new Date(ts * 1000).toLocaleTimeString(undefined, {
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
  });
}

function ChannelListItem({ event }: Props) {
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
        <span className={styles.eventTitle}>
          {parseTime(event.startsAt)} - {parseTime(event.endsAt)}
        </span>
      </div>
    </div>
  );
}

export default ChannelListItem;
