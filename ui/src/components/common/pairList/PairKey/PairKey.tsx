import { PropsWithChildren } from 'react';
import styles from './PairKey.module.scss';

function PairKey({ children }: PropsWithChildren) {
  return <div className={styles.PairKey}>{children}</div>;
}

export default PairKey;
