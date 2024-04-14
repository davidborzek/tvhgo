import { ReactElement } from 'react';
import { NavLink, useMatch } from 'react-router-dom';

import styles from './NavigationItem.module.scss';
import { c } from '@/utils/classNames';
import { INavigationItem } from '../types';

type Props = {
  to: string;
  icon?: ReactElement;
  title: string;
  items?: INavigationItem[];
  topLevel?: boolean;
};

function NavigationItem({ to, icon, title, items, topLevel }: Props) {
  const match = useMatch(to + '/' + '*');

  return (
    <div
      className={c(
        styles.container,
        topLevel ? styles.topLevel : '',
        !!match ? styles.active : ''
      )}
    >
      <NavLink
        className={({ isActive }) =>
          c(
            styles.root,
            isActive ? styles.active : '',
            topLevel ? styles.topLevel : ''
          )
        }
        to={to}
      >
        {icon || null}
        <span title={title}>{title}</span>
      </NavLink>

      {items && match && (
        <div className={styles.subItems}>
          {items.map(({ icon, title, to }) => (
            <NavigationItem icon={icon} title={title} to={to} />
          ))}
        </div>
      )}
    </div>
  );
}

export default NavigationItem;
