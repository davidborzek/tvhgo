import React from 'react';
import { c } from '../../utils/classNames';

import styles from './Button.module.scss';

export type ButtonStyle = 'red' | 'blue';

type Props = {
  label: string;
  type?: 'submit' | 'reset' | 'button';
  disabled?: boolean;
  loading?: boolean;
  loadingLabel?: string | null;
  className?: string;
  style?: ButtonStyle;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
};

function Button(props: Props) {
  const disabled = props.disabled || props.loading;

  const getStyleClass = () => {
    switch (props.style) {
      case 'red':
        return styles.red;
    }
    return '';
  };

  return (
    <button
      type={props.type}
      disabled={disabled}
      className={c(
        styles.button,
        getStyleClass(),
        disabled ? styles.disabled : '',
        props.className ? props.className : ''
      )}
      onClick={props.onClick}
    >
      {props.loading ? props.loadingLabel || '...' : props.label}
    </button>
  );
}

export default Button;
