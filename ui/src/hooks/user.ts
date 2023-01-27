import { UpdateUser } from './../clients/api/api.types';
import { useTranslation } from 'react-i18next';
import { useLoading } from './../contexts/LoadingContext';
import { useState, useEffect } from 'react';
import { getUser, updateUser } from '../clients/api/api';
import { UserResponse } from '../clients/api/api.types';
import { toast } from 'react-toastify';

export const useFetchUser = () => {
  const NOTIFICATION_ID = 'manageUser';

  const notifyError = (message?: string | null) => {
    toast.error(message, {
      toastId: NOTIFICATION_ID,
      updateId: NOTIFICATION_ID,
    });
  };

  const notifySuccess = (message?: string | null) => {
    toast.success(message, {
      toastId: NOTIFICATION_ID,
      updateId: NOTIFICATION_ID,
    });
  };

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

  const update = async (opts: UpdateUser, msg?: string | null) => {
    setIsLoading(true);
    await updateUser(opts)
      .then((user) => {
        notifySuccess(msg ? t(msg) : t('user_updated_successfully'));
        setUser(user);
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  return { update, user, error };
};
