import { ComponentPropsWithRef } from 'react';
import { c } from '@/utils/classNames';
import styles from './Badge.module.scss';

type BadgeProps = ComponentPropsWithRef<'span'> & {
  color?: 'default' | 'success' | 'failure' | 'warning' | 'indigo' | 'pink';
};

export default function Badge({
  children,
  className,
  color = 'default',
  ...rest
}: BadgeProps) {
  return (
    <span
      className={c(styles.badge, color && styles[color], className)}
      {...rest}
    >
      {children}
    </span>
  );
}
