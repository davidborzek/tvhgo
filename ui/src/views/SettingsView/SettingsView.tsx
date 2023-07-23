import { useTranslation } from 'react-i18next';
import { useSearchParams } from 'react-router-dom';
import styles from './SettingsView.module.scss';
import Tabs from '../../components/Tabs/Tabs';
import SecuritySettings from './SecuritySettings';
import GeneralSettings from './GeneralSettings';

enum TabView {
  GENERAL = 0,
  SECURITY = 1,
}

function SettingsView() {
  const [searchParams, setSearchParams] = useSearchParams();

  const { t } = useTranslation();

  const getActiveTab = () => {
    const tab = searchParams.get('tab');
    return tab ? (parseInt(tab) as TabView) : TabView.GENERAL;
  };

  const setActiveTab = (tab: TabView) => {
    setSearchParams((prev) => {
      prev.set('tab', `${tab}`);
      return prev;
    });
  };

  const activeTab = getActiveTab();

  const renderTab = () => {
    switch (activeTab) {
      case TabView.GENERAL:
        return <GeneralSettings />;

      case TabView.SECURITY:
        return <SecuritySettings />;
    }
  };

  return (
    <div className={styles.Settings}>
      <div className={styles.heading}>
        <Tabs
          tabs={[
            {
              label: t('general'),
              active: activeTab === TabView.GENERAL,
            },
            {
              label: t('security'),
              active: activeTab === TabView.SECURITY,
            },
          ]}
          onChange={setActiveTab}
        />
      </div>
      <div className={styles.content}>{renderTab()}</div>
    </div>
  );
}

export default SettingsView;
