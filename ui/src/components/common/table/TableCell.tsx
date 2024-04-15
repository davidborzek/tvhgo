import { ComponentPropsWithoutRef } from 'react';
import { c } from '@/utils/classNames';
import styles from './TableCell.module.scss';

export default function TableCell({
  children,
  className,
  ...props
}: ComponentPropsWithoutRef<'td'>) {
  return (
    <td className={c(styles.cell, className)} {...props}>
      {children}
    </td>
  );
}
