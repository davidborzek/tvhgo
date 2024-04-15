import { ComponentPropsWithoutRef } from 'react';
import { c } from '@/utils/classNames';
import styles from './Table.module.scss';

export default function Table({
  children,
  ...props
}: ComponentPropsWithoutRef<'table'>) {
  return (
    <table className={c(styles.table, props.className)} {...props}>
      {children}
    </table>
  );
}
