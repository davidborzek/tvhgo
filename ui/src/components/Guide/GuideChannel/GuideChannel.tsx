import Image from '@/components/Image/Image';
import styles from './GuideChannel.module.scss';

type Props = {
  name: string;
  picon: string;
  number: number;
  onClick: () => void;
};

function GuideChannel({ name, picon, number, onClick }: Props) {
  return (
    <div className={styles.container}>
      <div className={styles.channel} tabIndex={0} onClick={onClick}>
        <Image title={name} className={styles.picon} src={picon} />
        <span className={styles.number}>{number}</span>
      </div>
    </div>
  );
}

export default GuideChannel;
