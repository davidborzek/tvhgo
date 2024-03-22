import { useTranslation } from 'react-i18next';
import { Outlet } from 'react-router-dom';

import TabNavigation from '@/components/common/tabs/TabNavigation';

import styles from './SettingsView.module.scss';

export function Component() {
  const { t } = useTranslation();

  return (
    <div className={styles.Settings}>
      <div className={styles.heading}>
        <TabNavigation
          tabs={[
            {
              title: t('general'),
              to: '/settings/general',
            },
            {
              title: t('security'),
              to: '/settings/security',
            },
          ]}
        />
      </div>
      <div className={styles.content}>
        <Outlet />
      </div>
    </div>
  );
}

Component.displayName = 'SettingsView';
