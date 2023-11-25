import { Checkmark, Dash } from '@/assets';
import { c } from '@/utils/classNames';
import styles from './Checkbox.module.scss';

type Props = {
  onChange?: (checked: boolean) => void;
  checked?: boolean;
  disabled?: boolean;
  indeterminate?: boolean;
  className?: string;
};

const Checkbox = (props: Props) => {
  const getMark = () => {
    if (props.checked) {
      return <Checkmark className={styles.mark} />;
    }

    if (props.indeterminate) {
      return <Dash className={styles.mark} />;
    }
  };

  return (
    <div
      className={c(
        styles.container,
        props.className,
        props.checked || props.indeterminate ? styles.checked : '',
        props.disabled ? styles.disabled : ''
      )}
      onClick={(e) => {
        if (!props.disabled && props.onChange && e.target === e.currentTarget) {
          props.onChange(!props.checked);
        }
      }}
    >
      {getMark()}
      <input
        className={c(styles.checkbox)}
        type="checkbox"
        onChange={() => {
          if (props.onChange) {
            props.onChange(!props.checked);
          }
        }}
        checked={props.checked}
        disabled={props.disabled}
      />
    </div>
  );
};

export default Checkbox;
