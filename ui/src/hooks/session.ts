import { ApiError, deleteSession } from '@/clients/api/api';
import { useTranslation } from 'react-i18next';
import { useNotification } from './notification';
import { useRevalidator } from 'react-router-dom';

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

  return {
    revokeSession: _revokeSession,
  };
};
