import { ComponentPropsWithoutRef } from 'react';
import { c } from '@/utils/classNames';
import styles from './TableHeadCell.module.scss';

export default function TableHeadCell({
  children,
  className,
  ...props
}: ComponentPropsWithoutRef<'th'>) {
  return (
    <th className={c(styles.tableHeadCell, className)} {...props}>
      {children}
    </th>
  );
}
