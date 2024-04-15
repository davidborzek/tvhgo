import { ComponentPropsWithoutRef } from 'react';
import { c } from '@/utils/classNames';
import styles from './TableRow.module.scss';

export default function TableRow({
  children,
  ...props
}: ComponentPropsWithoutRef<'tr'>) {
  return (
    <tr className={c(styles.row, props.className)} {...props}>
      {children}
    </tr>
  );
}
