import { UpdateUser, UpdateUserPassword } from '@/clients/api/api.types';
import { useTranslation } from 'react-i18next';
import { useLoading } from '@/contexts/LoadingContext';
import { ApiError, updateUser, updateUserPassword } from '@/clients/api/api';
import { useAuth } from '@/contexts/AuthContext';
import { useNotification } from './notification';

export const useUpdateUser = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('manageUser');

  const { setUser } = useAuth();
  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const update = async (opts: UpdateUser) => {
    dismissNotification();
    setIsLoading(true);

    await updateUser(opts)
      .then((user) => {
        notifySuccess(t('user_updated_successfully'));
        setUser(user);
      })
      .catch((error) => {
        if (error instanceof ApiError && error.code === 400) {
          if (error.message === 'username already exists') {
            notifyError(t('username_already_exists'));
          } else if (error.message === 'email already exists') {
            notifyError(t('email_already_exists'));
          }
          return;
        }

        notifyError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  return { update };
};

export const useUpdateUserPassword = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('changeUserPassword');

  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const updatePassword = async (opts: UpdateUserPassword) => {
    dismissNotification();
    setIsLoading(true);

    return await updateUserPassword(opts)
      .then(() => {
        notifySuccess(t('password_updated_successfully'));
        return true;
      })
      .catch((error) => {
        if (
          error instanceof ApiError &&
          error.code === 400 &&
          error.message === 'current password is invalid'
        ) {
          notifyError(t('invalid_current_password'));
          return false;
        }

        notifyError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  return { updatePassword };
};
