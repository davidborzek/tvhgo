import Image from '@/components/Image/Image';
import styles from './EventChannelInfo.module.scss';

type Props = {
  channelName: string;
  picon: string;
};

function EventChannelInfo({ channelName, picon }: Props) {
  return (
    <div className={styles.EventChannelInfo}>
      <span className={styles.channelName}>{channelName}</span>
      <Image title={channelName} className={styles.picon} src={picon} />
    </div>
  );
}

export default EventChannelInfo;
