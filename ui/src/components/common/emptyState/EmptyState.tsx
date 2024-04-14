import { PropsWithChildren } from 'react';
import styles from './EmptyState.module.scss';

type Props = {
  title: string;
  subtitle?: string;
};

const EmptyState = ({
  title,
  subtitle,
  children,
}: PropsWithChildren<Props>) => {
  return (
    <div className={styles.emptyState}>
      <span className={styles.title}>{title}</span>
      {subtitle ? <span className={styles.subtitle}>{subtitle}</span> : <></>}
      {children}
    </div>
  );
};

export default EmptyState;
