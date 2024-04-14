import React, { forwardRef } from 'react';

import { Copy } from '@/assets';
import { c } from '@/utils/classNames';
import styles from './Input.module.scss';
import { useNotification } from '@/hooks/notification';
import { useTranslation } from 'react-i18next';

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
  selectTextOnFocus?: boolean;
  error?: string;
  maxWidth?: string | number;
  fullWidth?: boolean;
  showCopyButton?: boolean;
  ellipsis?: boolean;
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
            props.ellipsis ? styles.ellipsis : '',
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
            if (props.selectTextOnFocus) {
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
