import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import Button from '../../components/Button/Button';
import Dropdown, { Option } from '../../components/Dropdown/Dropdown';
import { Theme, useTheme } from '../../contexts/ThemeContext';
import i18n from '../../i18n/i18n';
import styles from './SettingsView.module.scss';

function SettingsView() {
  const navigate = useNavigate();
  const { t } = useTranslation();

  const { setTheme, theme } = useTheme();

  const themeOptions: Option[] = [
    { title: t('dark'), value: 'dark' },
    { title: t('light'), value: 'light' },
  ];

  const languageOptions: Option[] = [
    { title: t('english'), value: 'en' },
    { title: t('german'), value: 'de' },
  ];

  const handleChangeLanguage = (lang: string) => {
    i18n.changeLanguage(lang, () => window.location.reload());
  };

  return (
    <div className={styles.Settings}>
      <div>
        <Dropdown
          label={t('appearance')}
          value={theme}
          options={themeOptions}
          onChange={(theme) => setTheme(theme as Theme)}
          maxWidth="10rem"
          fullWidth
        />
      </div>
      <div>
        <Dropdown
          label={t('language')}
          value={i18n.language}
          options={languageOptions}
          onChange={handleChangeLanguage}
          maxWidth="10rem"
          fullWidth
        />
      </div>
      <div>
        <Button
          label={t('logout')}
          style="red"
          onClick={() => navigate('/logout')}
        />
      </div>
    </div>
  );
}

export default SettingsView;
