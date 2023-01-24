import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import Button from '../../components/Button/Button';
import styles from './SettingsView.module.scss';

function SettingsView() {
  const navigate = useNavigate();
  const { t } = useTranslation();

  return (
    <div className={styles.Settings}>
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
