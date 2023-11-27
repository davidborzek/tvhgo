import { PropsWithChildren } from 'react';
import { c } from '@/utils/classNames';
import styles from './Form.module.scss';

type Props = {
  maxWidth?: string | number;
  onSubmit?: React.FormEventHandler<HTMLFormElement>;
  className?: string;
};

function Form({
  children,
  className,
  maxWidth,
  onSubmit,
}: PropsWithChildren<Props>) {
  return (
    <form
      className={c(styles.Form, className)}
      onSubmit={onSubmit}
      style={{ maxWidth }}
    >
      {children}
    </form>
  );
}

export default Form;
