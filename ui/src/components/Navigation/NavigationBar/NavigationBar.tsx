import React from 'react';
import NavigationItem from '../NavigationItem/NavigationItem';
import { INavigationItem } from '../types';

import styles from './NavigationBar.module.scss';

type Props = {
  items: INavigationItem[];
};

function NavigationBar({ items }: Props) {
  return (
    <div className={styles.root}>
      <div className={styles.head}>
        <img className={styles.image} src="/img/tvhgo.png" />
      </div>
      <div className={styles.items}>
        {items.map(({ icon, title, to }) => (
          <NavigationItem icon={icon} title={title} to={to} key={title} />
        ))}
      </div>
    </div>
  );
}

export default NavigationBar;
