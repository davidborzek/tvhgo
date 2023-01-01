import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ApiError, getChannel } from '../clients/api/api';
import { Channel } from '../clients/api/api.types';

export const useFetchChannel = (id?: string) => {
  const { t } = useTranslation();

  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [channel, setChannel] = useState<Channel>();

  useEffect(() => {
    setLoading(true);
    getChannel(id || '')
      .then(setChannel)
      .catch((err) => {
        if (err instanceof ApiError && err.code === 404) {
          return;
        }

        setError(t('unexpected'));
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  return { channel, error, loading };
};
