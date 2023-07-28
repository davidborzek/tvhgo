import { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useTranslation } from 'react-i18next';
import { getUser, login, ApiError } from '../clients/api/api';
import { toast } from 'react-toastify';

type LoginFunc = (username: string, password: string, code?: string) => void;

const NOTIFICATION_ID = 'loginError';

const useLogin = () => {
  const { t } = useTranslation();
  const { setUser } = useAuth();
  const [loading, setLoading] = useState(false);
  const [twoFactorRequired, setTwoFactorRequired] = useState(false);

  const notify = (message?: string | null) => {
    toast.error(message, {
      toastId: NOTIFICATION_ID,
      updateId: NOTIFICATION_ID,
    });
  };

  const fetchUser = () => {
    setLoading(true);
    getUser()
      .then(setUser)
      .catch(() => {
        notify(t('unexpected'));
      })
      .finally(() => setLoading(false));
  };

  const handleLogin: LoginFunc = (username, password, code) => {
    toast.dismiss(NOTIFICATION_ID);
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
              notify(t('invalid 2fa code'));
              break;
            default:
              setTwoFactorRequired(false);
              notify(t('invalid_login'));
              break;
          }
        } else {
          notify(t('unexpected'));
          setTwoFactorRequired(false);
        }
      })
      .finally(() => setLoading(false));
  };

  return { login: handleLogin, loading, twoFactorRequired };
};

export default useLogin;
