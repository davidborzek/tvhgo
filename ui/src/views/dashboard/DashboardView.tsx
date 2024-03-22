import { useTranslation } from 'react-i18next';
import { Outlet, ScrollRestoration, useNavigation } from 'react-router-dom';

import { GuideIcon, RecordingsIcon, SettingsIcon, TvIcon } from '@/assets';
import MobileNavigation from '@/components/navigation/mobile/MobileNavigation';
import NavigationBar from '@/components/navigation/desktop/bar/NavigationBar';
import { INavigationItem } from '@/components/navigation/types';
import { c } from '@/utils/classNames';

import styles from './DashboardView.module.scss';

export function Component() {
  const { t } = useTranslation();
  const { state } = useNavigation();

  const navigationItems: INavigationItem[] = [
    { icon: <TvIcon />, title: t('channels'), to: '/channels' },
    { icon: <GuideIcon />, title: t('guide'), to: '/guide' },
    { icon: <RecordingsIcon />, title: t('recordings'), to: '/recordings' },
    { icon: <SettingsIcon />, title: t('settings'), to: '/settings' },
  ];

  return (
    <div className={c(styles.root)}>
      <ScrollRestoration />
      <div className={styles.mobileNavBar}>
        <MobileNavigation items={navigationItems} />
      </div>
      <div className={styles.desktopNavBar}>
        <NavigationBar items={navigationItems} />
      </div>
      <main
        className={c(styles.main, state === 'loading' ? styles.loading : undefined)}
      >
        <Outlet />
      </main>
    </div>
  );
}

Component.displayName = 'DashboardView';
