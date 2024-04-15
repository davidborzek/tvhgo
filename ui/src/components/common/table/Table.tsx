import { ComponentPropsWithoutRef } from 'react';
import { c } from '@/utils/classNames';
import styles from './Table.module.scss';

export default function Table({
  children,
  className,
  ...props
}: ComponentPropsWithoutRef<'table'>) {
  return (
    <table className={c(styles.table, className)} {...props}>
      {children}
    </table>
  );
}
