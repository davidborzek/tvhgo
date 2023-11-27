import { Close } from '@/assets';
import styles from './ModalCloseButton.module.scss';

type Props = {
  onClick?: () => void;
};

const ModalCloseButton = ({ onClick }: Props) => {
  return (
    <button className={styles.button} onClick={onClick}>
      <Close />
    </button>
  );
};

export default ModalCloseButton;
