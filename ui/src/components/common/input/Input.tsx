import React from 'react';
import { forwardRef } from 'react';
import { c } from '@/utils/classNames';

import styles from './Input.module.scss';
import { useTranslation } from 'react-i18next';
import { Copy } from '@/assets';
import { useNotification } from '@/hooks/notification';

type Props = {
  name?: string;
  type?: React.HTMLInputTypeAttribute;
  label?: string | null;
  placeholder?: string | null;
  value?: string | number;
  defaultValue?: string | number;
  onChange?: React.ChangeEventHandler<HTMLInputElement>;
  onBlur?: React.FocusEventHandler<HTMLInputElement>;
  onClick?: React.MouseEventHandler<HTMLInputElement>;
  className?: string;
  required?: boolean;
  disabled?: boolean;
  selectTextOnFocus?: boolean;
  error?: string;
  maxWidth?: string | number;
  fullWidth?: boolean;
  showCopyButton?: boolean;
  ellipsis?: boolean;
  hideCarret?: boolean;
  readonly?: boolean;
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
            props.error ? styles.error : '',
            props.onClick ? styles.clickable : '',
            props.hideCarret ? styles.hideCarret : ''
          )}
          name={props.name}
          placeholder={props.placeholder || undefined}
          value={props.value}
          defaultValue={props.defaultValue}
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
          onClick={props.onClick}
          readOnly={props.readonly}
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
