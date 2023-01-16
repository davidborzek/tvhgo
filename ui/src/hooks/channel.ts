import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ApiError, getChannel } from '../clients/api/api';
import { Channel } from '../clients/api/api.types';
import { useLoading } from '../contexts/LoadingContext';

export const useFetchChannel = (id?: string) => {
  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const [error, setError] = useState<string | null>(null);
  const [channel, setChannel] = useState<Channel>();

  useEffect(() => {
    setIsLoading(true);
    getChannel(id || '')
      .then(setChannel)
      .catch((err) => {
        if (err instanceof ApiError && err.code === 404) {
          return;
        }

        setError(t('unexpected'));
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, []);

  return { channel, error };
};
