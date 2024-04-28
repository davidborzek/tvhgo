import { ApiError, deleteSession, deleteUserSession } from '@/clients/api/api';

import { useNotification } from './notification';
import { useRevalidator } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

export const useManageSessions = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('manageSessions');
  const revalidator = useRevalidator();

  const { t } = useTranslation();

  const _revokeSession = async (id: number) => {
    dismissNotification();

    return await deleteSession(id)
      .then(() => {
        notifySuccess(t('session_revoked'));
        revalidator.revalidate();
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
      });
  };

  const _revokeUserSession = async (userId: number, sessionId: number) => {
    dismissNotification();

    return await deleteUserSession(userId, sessionId)
      .then(() => {
        notifySuccess(t('session_revoked'));
        revalidator.revalidate();
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
      });
  };

  return {
    revokeSession: _revokeSession,
    revokeUserSession: _revokeUserSession,
  };
};
