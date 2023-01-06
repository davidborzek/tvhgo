import { PropsWithChildren } from 'react';
import styles from './Pair.module.scss';

function Pair({ children }: PropsWithChildren) {
  return <div className={styles.Pair}>{children}</div>;
}

export default Pair;
