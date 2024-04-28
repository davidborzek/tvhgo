import {
  ApiError,
  createUser,
  deleteUser,
  updateUser,
  updateUserPassword,
} from '@/clients/api/api';
import {
  CreateUser,
  UpdateUser,
  UpdateUserPassword,
} from '@/clients/api/api.types';

import { useAuth } from '@/contexts/AuthContext';
import { useNotification } from './notification';
import { useRevalidator } from 'react-router-dom';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';

export const useUpdateUser = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('manageUser');

  const { setUser } = useAuth();
  const { t } = useTranslation();

  const update = async (opts: UpdateUser) => {
    dismissNotification();

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
      });
  };

  return { update };
};

export const useUpdateUserPassword = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('changeUserPassword');

  const { t } = useTranslation();

  const updatePassword = async (opts: UpdateUserPassword) => {
    dismissNotification();

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
      });
  };

  return { updatePassword };
};

export const useCreateUser = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('createUser');

  const { t } = useTranslation();

  const create = async (opts: CreateUser) => {
    dismissNotification();

    return await createUser(opts)
      .then(() => {
        notifySuccess(t('user_created_successfully'));
      })
      .catch((error) => {
        if (error instanceof ApiError && error.code === 400) {
          if (error.message === 'username already exists') {
            notifyError(t('username_already_exists'));
          } else if (error.message === 'email already exists') {
            notifyError(t('email_already_exists'));
          }
        } else {
          notifyError(t('unexpected'));
        }

        throw error;
      });
  };

  return { create };
};

export const useDeleteUser = (): [boolean, (id: number) => Promise<void>] => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('deleteUser');
  const [isPending, setIsPending] = useState(false);
  const { revalidate } = useRevalidator();

  const { t } = useTranslation();

  const _deleteUser = async (id: number) => {
    dismissNotification();
    setIsPending(true);
    return deleteUser(id)
      .then(() => {
        notifySuccess(t('user_deleted_successfully'));
        revalidate();
      })
      .catch((error) => {
        if (error instanceof ApiError && error.code === 400) {
          if (error.message === 'current user cannot be deleted') {
            notifyError(t('current_user_cannot_be_deleted'));
          }
        } else {
          notifyError(t('unexpected'));
        }

        throw error;
      })
      .finally(() => setIsPending(false));
  };

  return [isPending, _deleteUser];
};
