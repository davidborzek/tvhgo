import React, { ReactElement } from 'react';
import { c } from '@/utils/classNames';

import styles from './Button.module.scss';

export type ButtonStyle = 'red' | 'blue';

type Props = {
  label?: string;
  icon?: ReactElement;
  type?: 'submit' | 'reset' | 'button';
  disabled?: boolean;
  loading?: boolean;
  loadingLabel?: string | null;
  className?: string;
  style?: ButtonStyle;
  testID?: string;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
};

export const getStyleClass = (style?: ButtonStyle) => {
  switch (style) {
    case 'red':
      return styles.red;
  }
  return '';
};

function Button(props: Props) {
  const disabled = props.disabled || props.loading;

  const getLabel = () => {
    if (props.loading) {
      return props.loadingLabel || '...';
    }

    return props.label || '';
  };

  return (
    <button
      type={props.type}
      disabled={disabled}
      className={c(
        styles.button,
        getStyleClass(props.style),
        disabled ? styles.disabled : '',
        props.className ? props.className : ''
      )}
      onClick={props.onClick}
      data-testid={props.testID}
    >
      {props.icon}
      {getLabel()}
    </button>
  );
}

export default Button;
