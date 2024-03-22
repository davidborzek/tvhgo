import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { getEpg, GetEpgChannelEventsQuery } from '@/clients/api/api';
import { EpgChannel } from '@/clients/api/api.types';
import { useLoading } from '@/contexts/LoadingContext';

export const useFetchEpg = (q?: GetEpgChannelEventsQuery) => {
  const { t } = useTranslation();
  const { setIsLoading } = useLoading();

  const [error, setError] = useState<string | null>(null);
  const [epg, setEpg] = useState<EpgChannel[] | null>(null);
  const [startsAt, setStartsAt] = useState<number | undefined>(q?.startsAt);
  const [endsAt, setEndsAt] = useState<number | undefined>(q?.endsAt);

  const fetch = async () => {
    setIsLoading(true);
    try {
      const result = await getEpg({
        ...q,
        startsAt,
        endsAt,
      });
      setEpg(result);
    } catch (error) {
      setError(t('unexpected'));
    }
    setIsLoading(false);
  };

  useEffect(() => {
    fetch();
  }, [startsAt, endsAt]);

  return {
    error,
    events: epg,
    setEndsAt,
    endsAt,
    setStartsAt,
    startsAt,
  };
};
