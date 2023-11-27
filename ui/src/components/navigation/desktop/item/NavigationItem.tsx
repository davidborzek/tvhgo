import { ReactElement } from 'react';
import { NavLink } from 'react-router-dom';

import styles from './NavigationItem.module.scss';
import { c } from '@/utils/classNames';

type Props = {
  to: string;
  icon: ReactElement;
  title: string;
};

function NavigationItem({ to, icon, title }: Props) {
  return (
    <NavLink
      className={({ isActive }) =>
        c(styles.root, isActive ? styles.active : '')
      }
      to={to}
    >
      {icon}
      <span title={title}>{title}</span>
    </NavLink>
  );
}

export default NavigationItem;
