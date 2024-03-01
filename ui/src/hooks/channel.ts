import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import {
  ApiError,
  getChannel,
  getChannels,
  getEpgEvents,
} from '@/clients/api/api';
import { Channel, EpgEvent } from '@/clients/api/api.types';
import { useLoading } from '@/contexts/LoadingContext';

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

export const useGetChannels = (name?: string) => {
  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const [error, setError] = useState<string | null>(null);
  const [channels, setChannels] = useState<Array<Channel>>();

  const fetch = async () => {
    setIsLoading(true);

    return getChannels({ name })
      .then(setChannels)
      .catch((err) => {
        setError(t('unexpected'));
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  useEffect(() => {
    fetch();
  }, [name]);

  return { channels, error };
};
