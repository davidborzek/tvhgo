import { ComponentPropsWithoutRef } from 'react';
import { c } from '@/utils/classNames';
import styles from './TableHead.module.scss';

export default function TableHead({
  children,
  ...props
}: ComponentPropsWithoutRef<'thead'>) {
  return (
    <thead className={c(styles.tableHead, props.className)} {...props}>
      <tr>{children}</tr>
    </thead>
  );
}
