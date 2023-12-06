import { c } from '@/utils/classNames';
import styles from './PaginationControlButton.module.scss';

type Props = {
  disabled?: boolean;
  label?: string | null | undefined;
  onClick?: () => void;
  testID?: string;
};

const PaginationControlButton = ({
  disabled,
  label,
  onClick,
  testID,
}: Props) => {
  return (
    <button
      className={c(styles.button, disabled ? styles.disabled : '')}
      disabled={disabled}
      onClick={onClick}
      data-testid={testID}
    >
      {label}
    </button>
  );
};

export default PaginationControlButton;
