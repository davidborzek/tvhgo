import { toast } from 'react-toastify';

export const useNotification = (notificationId: string) => {
  const notifyError = (message?: string | null) => {
    toast.error(message, {
      toastId: notificationId,
      updateId: notificationId,
    });
  };

  const notifySuccess = (message?: string | null) => {
    toast.success(message, {
      toastId: notificationId,
      updateId: notificationId,
    });
  };

  const dismissNotification = () => toast.dismiss(notificationId);

  return { notifyError, notifySuccess, dismissNotification };
};
