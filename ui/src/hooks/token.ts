import { createToken, deleteToken } from '@/clients/api/api';

import { useNotification } from './notification';
import { useRevalidator } from 'react-router-dom';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';

export const useManageTokens = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('manageTokens');
  const { t } = useTranslation();
  const revalidator = useRevalidator();

  const _revokeToken = async (id: number) => {
    dismissNotification();

    return await deleteToken(id)
      .then(() => {
        notifySuccess(t('session_revoked'));
        revalidator.revalidate();
      })
      .catch(() => {
        notifyError(t('unexpected'));
      });
  };

  return {
    revokeToken: _revokeToken,
  };
};

export const useCreateToken = () => {
  const { notifyError, dismissNotification } = useNotification('createToken');

  const [token, setToken] = useState('');

  const { t } = useTranslation();

  const _createToken = async (name: string) => {
    dismissNotification();

    return await createToken(name)
      .then((res) => {
        setToken(res.token);
      })
      .catch(() => {
        notifyError(t('unexpected'));
      });
  };

  return {
    createToken: _createToken,
    setToken,
    token,
  };
};
