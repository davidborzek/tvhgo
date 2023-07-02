import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { ApiError, getChannel, getEpgEvents } from '../clients/api/api';
import { Channel, EpgEvent } from '../clients/api/api.types';
import { useLoading } from '../contexts/LoadingContext';

export const useFetchChannel = (
  id?: string,
  offset?: number,
  limit?: number
) => {
  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const [error, setError] = useState<string | null>(null);
  const [channel, setChannel] = useState<Channel>();
  const [events, setEvents] = useState<EpgEvent[]>([]);
  const [total, setTotal] = useState(0);

  const fetch = async () => {
    if (!id) {
      return;
    }

    setIsLoading(true);

    return await Promise.all([
      getChannel(id),
      getEpgEvents({ channel: id, offset, limit }),
    ])
      .then(([channel, _events]) => {
        setEvents(_events.entries);
        setChannel(channel);
        setTotal(_events.total);
      })
      .catch((err) => {
        if (err instanceof ApiError && err.code === 404) {
          setError(t('not_found'));
          return;
        }

        setError(t('unexpected'));
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  useEffect(() => {
    fetch();
  }, [id, offset, limit]);

  return { channel, events, total, error };
};
