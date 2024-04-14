import { GuideIcon, RecordingsIcon, SettingsIcon, TvIcon } from '@/assets';
import {
  Outlet,
  ScrollRestoration,
  useLocation,
  useNavigation,
} from 'react-router-dom';
import { useEffect, useState } from 'react';

import Header from '@/components/header/Header';
import { INavigationItem } from '@/components/navigation/types';
import NavigationBar from '@/components/navigation/bar/NavigationBar';
import { c } from '@/utils/classNames';
import styles from './DashboardView.module.scss';
import { useTranslation } from 'react-i18next';

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
