import React, { useState } from "react";
import { BurgerMenuIcon } from "../../assets";
import { INavigationItem } from "../Navigation/types";

import styles from "./MobileNavigation.module.scss";
import MobileNavigationItem from "./MobileNavigationItem";

type Props = {
  items: INavigationItem[];
};

function MobileNavigation({ items }: Props) {
  const [open, setOpen] = useState(false);

  return (
    <div className={styles.root}>
      <div className={styles.bar}>
        <BurgerMenuIcon
          className={styles.menuIcon}
          onClick={() => setOpen(!open)}
        />
        <span>tvhgo</span>
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
