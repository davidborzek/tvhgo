import { NavLink, useMatch } from 'react-router-dom';

import { INavigationItem } from '../types';
import { ReactElement } from 'react';
import { c } from '@/utils/classNames';
import styles from './NavigationItem.module.scss';

type Props = {
  to: string;
  icon?: ReactElement;
  title: string;
  items?: INavigationItem[];
  topLevel?: boolean;
  requiredRoles?: string[];
  roles?: string[];
};

function NavigationItem({
  to,
  icon,
  title,
  items,
  topLevel,
  requiredRoles = [],
  roles = [],
}: Props) {
  const match = useMatch(`${to}/` + `*`);

  if (
    requiredRoles.length > 0 &&
    !requiredRoles.some((role) => roles.includes(role))
  ) {
    return null;
  }

  return (
    <div
      className={c(
        styles.container,
        topLevel ? styles.topLevel : '',
        match ? styles.active : ''
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
          {items.map(({ icon, title, to, requiredRoles }) => (
            <NavigationItem
              icon={icon}
              title={title}
              to={to}
              requiredRoles={requiredRoles}
              roles={roles}
            />
          ))}
        </div>
      )}
    </div>
  );
}

export default NavigationItem;
