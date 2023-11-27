import styles from './Error.module.scss';

type Props = {
  message: string;
};

function Error({ message }: Props) {
  return <div className={styles.Error}>{message}</div>;
}

export default Error;
