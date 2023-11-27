import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { logout } from '@/clients/api/api';
import { useAuth } from '@/contexts/AuthContext';
import { useLoading } from '@/contexts/LoadingContext';

const useLogout = () => {
  const { t } = useTranslation();
  const authContext = useAuth();
  const { setIsLoading } = useLoading();
  const [error, setError] = useState<string | null>(null);

  const _logout = () => {
    setIsLoading(true);
    logout()
      .catch(() => {
        setError(t('unexpected'));
      })
      .then(() => authContext.setUser(null))
      .finally(() => setIsLoading(false));
  };

  return { logout: _logout, error };
};

export default useLogout;
