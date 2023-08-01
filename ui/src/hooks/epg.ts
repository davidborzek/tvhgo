import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import {
  ApiError,
  getEpg,
  GetEpgChannelEventsQuery,
  getEpgEvent,
  getEpgEvents,
  GetEpgEventsQuery,
  getRelatedEpgEvents,
} from '../clients/api/api';
import { EpgChannel, EpgEvent } from '../clients/api/api.types';
import { useLoading } from '../contexts/LoadingContext';

export const useFetchEvents = (q?: GetEpgEventsQuery) => {
  const { t } = useTranslation();

  const { setIsLoading } = useLoading();

  const [error, setError] = useState<string | null>(null);
  const [total, setTotal] = useState(0);
  const [epg, setEpg] = useState<EpgEvent[] | null>(null);

  useEffect(() => {
    setIsLoading(true);
    getEpgEvents(q)
      .then((events) => {
        setEpg(events.entries);
        setTotal(events.total);
      })
      .catch(() => {
        setError(t('unexpected'));
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, [q?.limit, q?.offset]);

  return { error, events: epg, total };
};

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

export const useFetchEvent = () => {
  const { t } = useTranslation();
  const { setIsLoading } = useLoading();

  const [error, setError] = useState<string | null>(null);
  const [event, setEvent] = useState<EpgEvent>();
  const [relatedEvents, setRelatedEvents] = useState<EpgEvent[]>([]);

  const fetch = async (id: number) => {
    setIsLoading(true);
    return await Promise.all([getEpgEvent(id), getRelatedEpgEvents(id)])
      .then(([event, related]) => {
        setEvent(event);
        setRelatedEvents(related.entries.filter((r) => r.id !== id));
      })
      .catch((error) => {
        if (error instanceof ApiError && error.code === 404) {
          setError(t('not_found'));
        } else {
          setError(t('unexpected'));
        }
      })
      .finally(() => setIsLoading(false));
  };

  return { error, event, relatedEvents, fetch };
};
