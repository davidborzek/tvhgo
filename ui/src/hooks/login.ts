import { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useTranslation } from 'react-i18next';
import { getUser, login, ApiError } from '../clients/api/api';
import { toast } from 'react-toastify';

type LoginFunc = (username: string, password: string) => void;

const NOTIFICATION_ID = 'loginError';

const useLogin = () => {
  const { t } = useTranslation();
  const { setUser } = useAuth();
  const [loading, setLoading] = useState(false);

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

  const handleLogin: LoginFunc = (username, password) => {
    toast.dismiss(NOTIFICATION_ID);
    setLoading(true);

    login(username, password)
      .then(fetchUser)
      .catch((error) => {
        if (error instanceof ApiError && error.code == 401) {
          notify(t('invalid_login'));
        } else {
          notify(t('unexpected'));
        }
      })
      .finally(() => setLoading(false));
  };

  return { login: handleLogin, loading };
};

export default useLogin;
