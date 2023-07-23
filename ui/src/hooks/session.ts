import { toast } from 'react-toastify';
import { useLoading } from '../contexts/LoadingContext';
import { useEffect, useState } from 'react';
import { ApiError, deleteSession, getSessions } from '../clients/api/api';
import { useTranslation } from 'react-i18next';
import { Session } from '../clients/api/api.types';

export const useManageSessions = () => {
  const NOTIFICATION_ID = 'manageSessions';

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

  const [error, setError] = useState<string | null>(null);
  const [sessions, setSessions] = useState<Array<Session>>([]);

  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const _getSessions = async () => {
    setIsLoading(true);
    return await getSessions()
      .then((sessions) => {
        setSessions(sessions);
      })
      .catch(() => {
        setError(t('unexpected'));
      })
      .finally(() => setIsLoading(false));
  };

  const _revokeSession = async (id: number) => {
    setIsLoading(true);
    return await deleteSession(id)
      .then(() => {
        _getSessions();
        notifySuccess(t('session_revoked'));
      })
      .catch((error) => {
        if (
          error instanceof ApiError &&
          error.code === 400 &&
          error.message === 'current session cannot be revoked'
        ) {
          notifyError(t('current_session_cannot_be_revoked'));
        } else {
          notifyError(t('unexpected'));
        }
      })
      .finally(() => setIsLoading(false));
  };

  useEffect(() => {
    _getSessions();
  }, []);

  return {
    getSessions: _getSessions,
    revokeSession: _revokeSession,
    sessions,
    error,
  };
};
