import React from 'react';
import { forwardRef } from 'react';
import { c } from '../../utils/classNames';

import styles from './Input.module.scss';
import { useTranslation } from 'react-i18next';
import { Copy } from '../../assets';
import { useNotification } from '../../hooks/notification';

type Props = {
  name?: string;
  type?: React.HTMLInputTypeAttribute;
  label?: string | null;
  placeholder?: string | null;
  value?: string | number;
  onChange?: React.ChangeEventHandler<HTMLInputElement>;
  onBlur?: React.FocusEventHandler<HTMLInputElement>;
  className?: string;
  required?: boolean;
  disabled?: boolean;
  selecTextOnFocus?: boolean;
  error?: string;
  maxWidth?: string | number;
  fullWidth?: boolean;
  showCopyButton?: boolean;
};

function Input(props: Props, ref: React.LegacyRef<HTMLInputElement>) {
  const { t } = useTranslation();
  const { notifySuccess } = useNotification('input');

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

      <div className={styles.innerContainer}>
        <input
          type={props.type}
          className={c(
            styles.input,
            props.disabled ? styles.disabled : '',
            props.error ? styles.error : ''
          )}
          name={props.name}
          placeholder={props.placeholder || undefined}
          value={props.value}
          onChange={props.onChange}
          onBlur={props.onBlur}
          required={props.required}
          ref={ref}
          disabled={props.disabled}
          onFocus={(evt) => {
            if (props.selecTextOnFocus) {
              evt.target.select();
            }
          }}
        />
        {props.showCopyButton ? (
          <span
            title={t('copy')}
            className={styles.copyButton}
            onClick={() => {
              if (props.value) {
                navigator.clipboard
                  .writeText(`${props.value}`)
                  .then(() => notifySuccess(t('copied_to_clipboard')));
              }
            }}
          >
            <Copy />
          </span>
        ) : (
          <></>
        )}
      </div>

      {props.error ? (
        <div className={styles.errorMessage}>{props.error}</div>
      ) : (
        <></>
      )}
    </div>
  );
}

export default forwardRef(Input);
