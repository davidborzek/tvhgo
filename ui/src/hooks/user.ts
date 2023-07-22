import { UpdateUser, UpdateUserPassword } from './../clients/api/api.types';
import { useTranslation } from 'react-i18next';
import { useLoading } from './../contexts/LoadingContext';
import { ApiError, updateUser, updateUserPassword } from '../clients/api/api';
import { toast } from 'react-toastify';
import { useAuth } from '../contexts/AuthContext';

export const useUpdateUser = () => {
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

  const update = async (opts: UpdateUser) => {
    setIsLoading(true);
    await updateUser(opts)
      .then((user) => {
        notifySuccess(t('user_updated_successfully'));
        setUser(user);
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  return { update };
};

export const useUpdateUserPassword = () => {
  const NOTIFICATION_ID = 'changeUserPassword';

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

  const { setIsLoading } = useLoading();

  const updatePassword = async (opts: UpdateUserPassword) => {
    setIsLoading(true);
    await updateUserPassword(opts)
      .then(() => {
        notifySuccess(t('password_updated_successfully'));
      })
      .catch((error) => {
        if (
          error instanceof ApiError &&
          error.code === 400 &&
          error.message === 'current password is invalid'
        ) {
          notifyError(t('invalid_current_password'));
          return;
        }

        notifyError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  return { updatePassword };
};
