import React, { ReactElement } from 'react';

import { c } from '@/utils/classNames';
import styles from './Button.module.scss';

export type ButtonStyle = 'red' | 'blue' | 'text';

type Props = {
  label?: string;
  icon?: ReactElement;
  type?: 'submit' | 'reset' | 'button';
  size?: 'small' | 'medium' | 'large';
  disabled?: boolean;
  quiet?: boolean;
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
    case 'text':
      return styles.text;
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
        props.quiet ? styles.quiet : '',
        props.size ? styles[props.size] : styles.medium,
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
