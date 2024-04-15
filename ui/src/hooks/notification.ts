import { toast } from 'react-toastify';
import { useCallback } from 'react';

export const useNotification = (notificationId: string) => {
  const notifyError = useCallback(
    (message?: string | null) => {
      toast.error(message, {
        toastId: notificationId,
        updateId: notificationId,
      });
    },
    [notificationId]
  );

  const notifySuccess = useCallback(
    (message?: string | null) => {
      toast.success(message, {
        toastId: notificationId,
        updateId: notificationId,
      });
    },
    [notificationId]
  );

  const dismissNotification = () => toast.dismiss(notificationId);

  return { notifyError, notifySuccess, dismissNotification };
};
