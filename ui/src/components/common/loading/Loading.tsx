import styles from './Loading.module.scss';

function Loading() {
  return (
    <div className={styles.Loading}>
      <div className={styles.ldsEllipsis}>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
      </div>
    </div>
  );
}

export default Loading;
