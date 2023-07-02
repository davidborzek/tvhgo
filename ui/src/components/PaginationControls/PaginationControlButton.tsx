import { c } from '../../utils/classNames';
import styles from './PaginationControlButton.module.scss';

type Props = {
  disabled?: boolean;
  label?: string | null | undefined;
  onClick?: () => void;
};

const PaginationControlButton = ({ disabled, label, onClick }: Props) => {
  return (
    <button
      className={c(styles.button, disabled ? styles.disabled : '')}
      disabled={disabled}
      onClick={onClick}
    >
      {label}
    </button>
  );
};

export default PaginationControlButton;
