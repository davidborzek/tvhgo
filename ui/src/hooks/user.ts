import { UpdateUser } from './../clients/api/api.types';
import { useTranslation } from 'react-i18next';
import { useLoading } from './../contexts/LoadingContext';
import { useState, useEffect } from 'react';
import { getUser, updateUser } from '../clients/api/api';
import { UserResponse } from '../clients/api/api.types';

export const useFetchUser = () => {
  const { t } = useTranslation();
  const [user, setUser] = useState<UserResponse>();
  const [error, setError] = useState<string | null>(null);

  const { setIsLoading } = useLoading();

  useEffect(() => {
    setIsLoading(true);
    getUser()
      .then(setUser)
      .catch(() => setError(t('unexpected')))
      .finally(() => setIsLoading(false));
  }, []);

  const update = async (opts: UpdateUser) => {
    setIsLoading(true);
    await updateUser(opts)
      .then(setUser)
      .finally(() => setIsLoading(false));
  };

  return { update, user, error };
};
