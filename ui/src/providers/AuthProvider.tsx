import { PropsWithChildren, ReactElement, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { getUser, ApiError } from '../clients/api/api';
import { UserResponse } from '../clients/api/api.types';
import { AuthContext } from '../contexts/AuthContext';
import { useNotification } from '../hooks/notification';

export default function AuthProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const { t } = useTranslation();

  const [user, setUser] = useState<UserResponse | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  const { notifyError } = useNotification('authError');

  useEffect(() => {
    getUser()
      .then((user) => {
        setUser(user);
      })
      .catch((error) => {
        if (error instanceof ApiError && error.code === 401) {
          setUser(null);
        } else {
          notifyError(t('unexpected'));
        }
      })
      .finally(() => setIsLoading(false));
  }, []);

  if (isLoading) {
    return <></>;
  }

  return (
    <AuthContext.Provider
      value={{
        user,
        setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
