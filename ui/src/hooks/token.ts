import { useLoading } from '@/contexts/LoadingContext';
import { useEffect, useState } from 'react';
import { createToken, deleteToken, getTokens } from '@/clients/api/api';
import { useTranslation } from 'react-i18next';
import { Token } from '@/clients/api/api.types';
import { useNotification } from './notification';

export const useManageTokens = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('manageTokens');

  const [error, setError] = useState<string | null>(null);
  const [tokens, setTokens] = useState<Array<Token>>([]);

  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const _getTokens = async () => {
    setIsLoading(true);
    return await getTokens()
      .then((sessions) => {
        setTokens(sessions);
      })
      .catch(() => {
        setError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  const _revokeToken = async (id: number) => {
    dismissNotification();
    setIsLoading(true);

    return await deleteToken(id)
      .then(() => {
        _getTokens();
        notifySuccess(t('session_revoked'));
      })
      .catch((error) => {
        notifyError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  useEffect(() => {
    _getTokens();
  }, []);

  return {
    getTokens: _getTokens,
    revokeToken: _revokeToken,
    tokens,
    error,
  };
};

export const useCreateToken = () => {
  const { notifyError, dismissNotification } = useNotification('createToken');

  const [token, setToken] = useState('');

  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const _createToken = async (name: string) => {
    dismissNotification();
    setIsLoading(true);

    return await createToken(name)
      .then((res) => {
        setToken(res.token);
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  return {
    createToken: _createToken,
    setToken,
    token,
  };
};
