import { deleteDVRConfig } from '@/clients/api/api';
import { useNotification } from './notification';
import { useRevalidator } from 'react-router-dom';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';

export const useDVRConfig = (): [boolean, (id: string) => Promise<void>] => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('dvrConfig');
  const [isPending, setIsPending] = useState(false);
  const { revalidate } = useRevalidator();

  const { t } = useTranslation();

  const _deleteDVRConfig = async (id: string) => {
    dismissNotification();
    setIsPending(true);
    return deleteDVRConfig(id)
      .then(() => {
        notifySuccess(t('dvr_config_deleted'));
        revalidate();
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => setIsPending(false));
  };

  return [isPending, _deleteDVRConfig];
};
