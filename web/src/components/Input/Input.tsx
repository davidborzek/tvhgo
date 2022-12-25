import React from "react";
import { forwardRef } from "react";

import styles from "./Input.module.scss";

type Props = {
  name?: string;
  type?: React.HTMLInputTypeAttribute;
  label?: string | null;
  value?: string | number;
  onChange?: React.ChangeEventHandler<HTMLInputElement>;
  onBlur?: React.FocusEventHandler<HTMLInputElement>;
  className?: string;
  required?: boolean;
  error?: string;
};

function Input(props: Props, ref: React.LegacyRef<HTMLInputElement>) {
  return (
    <div
      className={`${styles.inputContainer} ${
        props.className ? props.className : ""
      }`}
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
        className={`${styles.input} ${props.error ? styles.error : ""}`}
        name={props.name}
        value={props.value}
        onChange={props.onChange}
        onBlur={props.onBlur}
        required={props.required}
        ref={ref}
      />
      {props.error ? (
        <div className={styles.errorMessage}>
          {props.error}
        </div>
      ) : (
        <></>
      )}
    </div>
  );
}

export default forwardRef(Input);
