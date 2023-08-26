import styles from './TabNavigation.module.scss';
import { c } from '../../utils/classNames';
import { NavLink } from 'react-router-dom';
import { useState } from 'react';

export type Tab = {
  title: string;
  to: string;
};

type TabNavigationItemProps = {
  title: string;
  to: string;
};

const TabNavigationItem = (props: TabNavigationItemProps) => {
  const [active, setActive] = useState(false);

  return (
    <div className={c(styles.TabItem, active ? styles.active : '')}>
      <NavLink
        className={({ isActive }) => {
          setActive(isActive);
          return c(styles.button, isActive ? styles.active : '');
        }}
        to={props.to}
      >
        {props.title}
      </NavLink>
    </div>
  );
};

type TabsProps = {
  tabs: Tab[];
};

const TabNavigation = (props: TabsProps) => {
  return (
    <nav className={styles.Tabs}>
      {props.tabs.map((tab, index) => (
        <TabNavigationItem key={index} title={tab.title} to={tab.to} />
      ))}
    </nav>
  );
};

export default TabNavigation;
