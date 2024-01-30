import { useEffect, useState } from 'react';
import { TwoFactorAuthSettings } from '@/clients/api/api.types';
import { useLoading } from '@/contexts/LoadingContext';
import {
  ApiError,
  activateTwoFactorAuth,
  deactivateTwoFactorAuth,
  getTwoFactorAuthSettings,
  setupTwoFactorAuth,
} from '@/clients/api/api';
import { useTranslation } from 'react-i18next';
import { useNotification } from './notification';

export const useTwoFactorAuthSettings = () => {
  const [twoFactorAuthSettings, setTwoFactorAuthSettings] =
    useState<TwoFactorAuthSettings | null>(null);

  const { setIsLoading } = useLoading();

  const fetchTwoFactorAuthSettings = async () => {
    setIsLoading(true);
    return await getTwoFactorAuthSettings()
      .then(setTwoFactorAuthSettings)
      .finally(() => setIsLoading(false));
  };

  useEffect(() => {
    fetchTwoFactorAuthSettings();
  }, []);

  return {
    fetchTwoFactorAuthSettings,
    twoFactorAuthSettings,
  };
};

export const useSetupTwoFactorAuth = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('setupTwoFactorAuth');

  const { t } = useTranslation();

  const [loading, setLoading] = useState<boolean>(false);
  const [twoFactorUrl, setTwoFactorUrl] = useState<string | null>(null);

  return {
    setupTwoFactorAuth: async (password: string) => {
      dismissNotification();
      setLoading(true);
      return await setupTwoFactorAuth(password)
        .then(({ url }) => setTwoFactorUrl(url))
        .catch((error) => {
          if (
            error instanceof ApiError &&
            error.code === 400 &&
            error.message === 'confirmation password is invalid'
          ) {
            notifyError(t('invalid_current_password'));
          } else if (error.message === 'two factor auth is already enabled') {
            notifyError(t('two_factor_auth_already_enabled'));
          } else {
            notifyError(t('unexpected'));
          }

          throw error;
        })
        .finally(() => setLoading(false));
    },
    activateTwoFactorAuth: async (password: string, code: string) => {
      dismissNotification();
      setLoading(true);

      return await activateTwoFactorAuth(password, code)
        .then(() => {
          notifySuccess(t('two_factor_auth_successfully_enabled'));
        })
        .catch((error) => {
          if (error instanceof ApiError && error.code === 400) {
            if (error.message === 'confirmation password is invalid') {
              notifyError(t('invalid_current_password'));
            } else if (error.message === 'invalid two factor code provided') {
              notifyError(t('invalid_verification_code'));
            } else {
              notifyError(t('unexpected'));
            }
          } else {
            notifyError(t('unexpected'));
          }

          throw error;
        })
        .finally(() => setLoading(false));
    },
    setTwoFactorUrl,
    twoFactorUrl,
    loading,
  };
};

export const useDeactivateTwoFactorAuth = () => {
  const { notifyError, notifySuccess, dismissNotification } = useNotification(
    'deactivateTwoFactorAuth'
  );

  const { t } = useTranslation();

  const [loading, setLoading] = useState<boolean>(false);

  return {
    deactivateTwoFactorAuth: async (code: string) => {
      dismissNotification();
      setLoading(true);

      return await deactivateTwoFactorAuth(code)
        .then(() => notifySuccess(t('two_factor_auth_successfully_disabled')))
        .catch((error) => {
          if (
            error instanceof ApiError &&
            error.code === 400 &&
            error.message === 'invalid two factor code provided'
          ) {
            notifyError(t('invalid_verification_code'));
          } else {
            notifyError(t('unexpected'));
          }

          throw error;
        })
        .finally(() => setLoading(false));
    },
    loading,
  };
};
