import { PropsWithChildren } from 'react';
import { c } from '../../../utils/classNames';
import styles from './FormGroup.module.scss';

type Props = {
  heading?: string | null;
  info?: string | null;
  direction?: 'row' | 'column';
  maxWidth?: string | number;
};

function FormGroup({
  direction,
  info,
  heading,
  maxWidth,
  children,
}: PropsWithChildren<Props>) {
  return (
    <div className={styles.FormGroup}>
      {heading && <span className={styles.heading}>{heading}</span>}
      <div
        className={c(
          styles.content,
          direction == 'column' ? styles.column : styles.row
        )}
        style={{
          maxWidth,
        }}
      >
        {children}
      </div>
      {info && <span className={styles.info}>{info}</span>}
    </div>
  );
}

export default FormGroup;
