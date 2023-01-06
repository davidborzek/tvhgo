import { PropsWithChildren } from 'react';
import styles from './PairList.module.scss';

type Props = {
  maxWidth?: string | number;
  minWidth?: string | number;
};

function PairList({ children, maxWidth, minWidth }: PropsWithChildren<Props>) {
  return (
    <div className={styles.PairList} style={{ maxWidth, minWidth }}>
      {children}
    </div>
  );
}

export default PairList;
