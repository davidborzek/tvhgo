import { useState } from 'react';
import { useLocation } from 'react-router-dom';
import { BurgerMenuIcon } from '@/assets';

import { INavigationItem } from '@/components/navigation/desktop/types';

import styles from './MobileNavigation.module.scss';
import MobileNavigationItem from './MobileNavigationItem';

type Props = {
  items: INavigationItem[];
};

function MobileNavigation({ items }: Props) {
  const [open, setOpen] = useState(false);
  const location = useLocation();

  const getPageTitle = () => {
    return (
      items.find((item) => item.to === location.pathname)?.title || 'tvhgo'
    );
  };

  return (
    <div className={styles.root}>
      <div className={styles.bar}>
        <BurgerMenuIcon
          className={styles.menuIcon}
          onClick={() => setOpen(!open)}
        />
        <span>{getPageTitle()}</span>
      </div>
      <div className={open ? styles.opened : styles.closed}>
        {items.map(({ title, icon, to }) => (
          <MobileNavigationItem
            title={title}
            key={title}
            icon={icon}
            to={to}
            onClick={() => setOpen(false)}
          />
        ))}
      </div>
    </div>
  );
}

export default MobileNavigation;
