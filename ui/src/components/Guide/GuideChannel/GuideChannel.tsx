import Image from '../../Image/Image';
import styles from './GuideChannel.module.scss';

type Props = {
  name: string;
  picon: string;
  number: number;
};

function GuideChannel({ name, picon, number }: Props) {
  return (
    <div className={styles.channel} tabIndex={0}>
      <Image title={name} className={styles.picon} src={picon} />
      <span className={styles.number}>{number}</span>
    </div>
  );
}

export default GuideChannel;
