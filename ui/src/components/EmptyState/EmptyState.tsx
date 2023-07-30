import styles from './EmptyState.module.scss';

type Props = {
  title: string;
  subtitle?: string;
};

const EmptyState = ({ title, subtitle }: Props) => {
  return (
    <div className={styles.emptyState}>
      <span className={styles.title}>{title}</span>
      {subtitle ? <span className={styles.subtitle}>{subtitle}</span> : <></>}
    </div>
  );
};

export default EmptyState;
