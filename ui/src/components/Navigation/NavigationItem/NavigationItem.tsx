import React, { ReactElement } from 'react';
import { NavLink } from 'react-router-dom';
import { c } from '../../../utils/classNames';

import styles from './NavigationItem.module.scss';

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
      {title}
    </NavLink>
  );
}

export default NavigationItem;
