import { PropsWithChildren } from 'react';
import styles from './PairValue.module.scss';

function PairValue({ children }: PropsWithChildren) {
  return <div className={styles.PairValue}>{children}</div>;
}

export default PairValue;
