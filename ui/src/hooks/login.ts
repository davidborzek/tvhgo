import { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useTranslation } from 'react-i18next';
import { getUser, login, ApiError } from '../clients/api/api';
import { useNotification } from './notification';

type LoginFunc = (username: string, password: string, code?: string) => void;

const useLogin = () => {
  const { t } = useTranslation();
  const { setUser } = useAuth();
  const [loading, setLoading] = useState(false);
  const [twoFactorRequired, setTwoFactorRequired] = useState(false);

  const { notifyError, dismissNotification } = useNotification('loginError');

  const fetchUser = () => {
    setLoading(true);
    getUser()
      .then(setUser)
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => setLoading(false));
  };

  const handleLogin: LoginFunc = (username, password, code) => {
    dismissNotification();
    setLoading(true);

    login(username, password, code)
      .then(fetchUser)
      .catch((error) => {
        if (error instanceof ApiError && error.code == 401) {
          switch (error.message) {
            case 'two factor auth is required':
              setTwoFactorRequired(true);
              break;
            case 'invalid two factor code provided':
              notifyError(t('invalid_verification_code'));
              break;
            default:
              setTwoFactorRequired(false);
              notifyError(t('invalid_login'));
              break;
          }
        } else {
          notifyError(t('unexpected'));
          setTwoFactorRequired(false);
        }
      })
      .finally(() => setLoading(false));
  };

  return { login: handleLogin, loading, twoFactorRequired };
};

export default useLogin;
