import { useTranslation } from 'react-i18next';
import styles from './TwoFactorAuthSettingsOverview.module.scss';
import { TwoFactorAuthSettings } from '@/clients/api/api.types';
import Button from '../Button/Button';
import Headline from '../Headline/Headline';

type Props = {
  settings: TwoFactorAuthSettings | null;
  onDisable: () => void;
  onEnable: () => void;
};

const TwoFactorAuthSettingsOverview = ({
  settings,
  onDisable,
  onEnable,
}: Props) => {
  const { t } = useTranslation();

  return (
    <div className={styles.settings}>
      <Headline>{t('two_factor_auth')}</Headline>
      {settings?.enabled ? (
        <Button style="red" label={t('disable')} onClick={onDisable} />
      ) : (
        <Button label={t('enable')} onClick={onEnable} />
      )}
    </div>
  );
};

export default TwoFactorAuthSettingsOverview;
