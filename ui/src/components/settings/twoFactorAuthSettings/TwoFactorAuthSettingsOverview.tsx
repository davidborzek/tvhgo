import Button from '@/components/common/button/Button';
import Headline from '@/components/common/headline/Headline';
import { TestIds } from '@/__test__/ids';
import { TwoFactorAuthSettings } from '@/clients/api/api.types';
import styles from './TwoFactorAuthSettingsOverview.module.scss';
import { useTranslation } from 'react-i18next';

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
        <Button
          style="red"
          label={t('disable')}
          onClick={onDisable}
          testID={TestIds.TWOFA_DISABLE_BUTTON}
        />
      ) : (
        <Button
          label={t('enable')}
          onClick={onEnable}
          testID={TestIds.TWOFA_ENABLE_BUTTON}
        />
      )}
    </div>
  );
};

export default TwoFactorAuthSettingsOverview;
