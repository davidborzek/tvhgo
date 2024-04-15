import { PropsWithChildren } from 'react';
import { c } from '@/utils/classNames';
import styles from './Tooltip.module.scss';

type TooltipProps = {
  direction?: 'top' | 'bottom' | 'left' | 'right';
  text: string;
  className?: string;
  disabled?: boolean;
};

export default function Tooltip({
  children,
  text,
  direction = 'right',
  disabled = false,
  className,
}: PropsWithChildren<TooltipProps>) {
  return (
    <div className={c(styles.wrapper)}>
      {children}

      {!disabled && (
        <div className={c(styles.tooltip, styles[direction], className)}>
          {text}
        </div>
      )}
    </div>
  );
}
