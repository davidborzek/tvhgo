import { UpdateUser } from './../clients/api/api.types';
import { useTranslation } from 'react-i18next';
import { useLoading } from './../contexts/LoadingContext';
import { updateUser } from '../clients/api/api';
import { toast } from 'react-toastify';
import { useAuth } from '../contexts/AuthContext';

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

  const { setUser } = useAuth();
  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

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

  return { update };
};
