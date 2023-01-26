import React from 'react';
import { forwardRef } from 'react';
import { c } from '../../utils/classNames';

import styles from './Input.module.scss';

type Props = {
  name?: string;
  type?: React.HTMLInputTypeAttribute;
  label?: string | null;
  placeholder?: string;
  value?: string | number;
  onChange?: React.ChangeEventHandler<HTMLInputElement>;
  onBlur?: React.FocusEventHandler<HTMLInputElement>;
  className?: string;
  required?: boolean;
  disabled?: boolean;
  error?: string;
  maxWidth?: string | number;
  fullWidth?: boolean;
};

function Input(props: Props, ref: React.LegacyRef<HTMLInputElement>) {
  return (
    <div
      className={c(
        styles.inputContainer,
        props.className ? props.className : ''
      )}
      style={{
        maxWidth: props.maxWidth,
        width: props.fullWidth ? '100%' : 'fit-content',
      }}
    >
      {props.label ? (
        <label className={styles.inputLabel} htmlFor={props.name}>
          {props.label}
        </label>
      ) : (
        <></>
      )}
      <input
        type={props.type}
        className={c(
          styles.input,
          props.disabled ? styles.disabled : '',
          props.error ? styles.error : ''
        )}
        name={props.name}
        placeholder={props.placeholder}
        value={props.value}
        onChange={props.onChange}
        onBlur={props.onBlur}
        required={props.required}
        ref={ref}
        disabled={props.disabled}
      />
      {props.error ? (
        <div className={styles.errorMessage}>{props.error}</div>
      ) : (
        <></>
      )}
    </div>
  );
}

export default forwardRef(Input);
