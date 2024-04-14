import { useTranslation } from 'react-i18next';
import {
  Outlet,
  ScrollRestoration,
  useLocation,
  useNavigation,
} from 'react-router-dom';

import { GuideIcon, RecordingsIcon, SettingsIcon, TvIcon } from '@/assets';
import { INavigationItem } from '@/components/navigation/types';
import { c } from '@/utils/classNames';

import styles from './DashboardView.module.scss';
import { useEffect, useState } from 'react';
import Header from '@/components/header/Header';
import NavigationBar from '@/components/navigation/bar/NavigationBar';

export function Component() {
  const { t } = useTranslation();
  const { state } = useNavigation();
  const { pathname } = useLocation();

  const [expanded, setExpanded] = useState(false);

  useEffect(() => {
    setExpanded(false);
  }, [pathname]);

  const navigationItems: INavigationItem[] = [
    { icon: <TvIcon />, title: t('channels'), to: '/channels' },
    { icon: <GuideIcon />, title: t('guide'), to: '/guide' },
    { icon: <RecordingsIcon />, title: t('recordings'), to: '/recordings' },
    {
      icon: <SettingsIcon />,
      title: t('settings'),
      to: '/settings',
      items: [
        { title: t('General'), to: '/settings/general' },
        { title: t('Security'), to: '/settings/security' },
      ],
    },
  ];

  return (
    <div className={c(styles.root)}>
      <ScrollRestoration />
      <Header onToggle={() => setExpanded(!expanded)} />
      <div className={c(styles.navigation, expanded ? styles.expanded : '')}>
        <NavigationBar items={navigationItems} />
      </div>
      <main
        className={c(
          styles.main,
          state === 'loading' ? styles.loading : undefined
        )}
      >
        <Outlet />
      </main>
    </div>
  );
}

Component.displayName = 'DashboardView';
