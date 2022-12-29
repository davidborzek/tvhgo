import React, { ReactElement } from "react";
import { NavLink } from "react-router-dom";
import { ArrowRightIcon } from "../../assets";
import { c } from "../../utils/classNames";

import styles from "./MobileNavigationItem.module.scss";

type Props = {
  title: string;
  icon: ReactElement;
  to: string;
  onClick?: () => void;
};

function MobileNavigationItem({ title, icon, to, onClick }: Props) {
  return (
    <NavLink
      onClick={onClick}
      to={to}
      className={({ isActive }) =>
        c(styles.root, isActive ? styles.active : "")
      }
    >
      {icon}
      {title}
      <ArrowRightIcon className={styles.arrow} />
    </NavLink>
  );
}

export default MobileNavigationItem;
