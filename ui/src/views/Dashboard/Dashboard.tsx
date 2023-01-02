import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';
import {
  GuideIcon,
  LogoutIcon,
  RecordingsIcon,
  SettingsIcon,
  TvIcon,
} from '../../assets';
import MobileNavigation from '../../components/MobileNavigation/MobileNavigation';
import NavigationBar from '../../components/Navigation/NavigationBar/NavigationBar';
import { INavigationItem } from '../../components/Navigation/types';
import { c } from '../../utils/classNames';
import styles from './Dashboard.module.scss';

function Dashboard() {
  const { t } = useTranslation();

  const navigationItems: INavigationItem[] = [
    { icon: <TvIcon />, title: t('channels'), to: '/' },
    { icon: <GuideIcon />, title: t('guide'), to: '/guide' },
    { icon: <RecordingsIcon />, title: t('recordings'), to: '/recordings' },
    { icon: <SettingsIcon />, title: t('settings'), to: '/settings' },
    { icon: <LogoutIcon />, title: t('logout'), to: '/logout' },
  ];

  return (
    <div className={c(styles.root, styles.desktop)}>
      <div className={styles.mobileNavBar}>
        <MobileNavigation items={navigationItems} />
      </div>
      <div className={styles.desktopNavBar}>
        <NavigationBar items={navigationItems} />
      </div>
      <main className={styles.main}>
        <Outlet />
      </main>
    </div>
  );
}

export default Dashboard;
