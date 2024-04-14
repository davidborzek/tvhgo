import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';

import styles from './SettingsView.module.scss';

export function Component() {
  const { t } = useTranslation();

  return (
    <div className={styles.Settings}>
      <div className={styles.content}>
        <Outlet />
      </div>
    </div>
  );
}

Component.displayName = 'SettingsView';
