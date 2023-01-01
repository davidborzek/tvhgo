import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import {
  ApiError,
  getEpgChannelEvents,
  GetEpgChannelEventsQuery,
  getEpgEvent,
  getEpgEvents,
  GetEpgEventsQuery,
  getRelatedEpgEvents,
} from '../clients/api/api';
import { EpgChannel, EpgEvent } from '../clients/api/api.types';

export const useFetchEpg = (q?: GetEpgEventsQuery) => {
  const initialOffset = q?.offset || 0;

  const { t } = useTranslation('errors');

  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);
  const [epg, setEpg] = useState<EpgEvent[]>([]);
  const [offset, setOffset] = useState(initialOffset);

  useEffect(() => {
    setLoading(true);
    getEpgEvents({ ...q, offset })
      .then((events) => {
        if (offset !== initialOffset) {
          setEpg([...epg, ...events.entries]);
        } else {
          setEpg(events.entries);
        }

        setTotal(events.total);
      })
      .catch(() => {
        setError(t('unexpected'));
      })
      .finally(() => {
        setLoading(false);
      });
  }, [offset]);

  const increaseOffset = (value: number) => {
    setOffset((oldOffset) => oldOffset + value);
  };

  return { error, loading, events: epg, increaseOffset, offset, total };
};

export const useFetchChannelEvents = (q?: GetEpgChannelEventsQuery) => {
  const { t } = useTranslation('errors');

  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);
  const [epg, setEpg] = useState<EpgChannel[]>([]);
  const [startsAt, setStartsAt] = useState<number | undefined>(q?.startsAt);
  const [endsAt, setEndsAt] = useState<number | undefined>(q?.endsAt);

  const fetch = async () => {
    setLoading(true);
    try {
      const meta = await getEpgChannelEvents({
        ...q,
        limit: 1,
        startsAt,
        endsAt,
      });
      const result = await getEpgChannelEvents({
        ...q,
        limit: meta.total,
        startsAt,
        endsAt,
      });
      setEpg(result.entries);
      setTotal(result.total);
    } catch (error) {
      setError(t('unexpected'));
    }
    setLoading(false);
  };

  useEffect(() => {
    fetch();
  }, [startsAt, endsAt]);

  return {
    error,
    loading,
    events: epg,
    total,
    setEndsAt,
    endsAt,
    setStartsAt,
    startsAt,
  };
};

export const useFetchEvent = () => {
  const { t } = useTranslation('errors');

  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [event, setEvent] = useState<EpgEvent>();
  const [relatedEvents, setRelatedEvents] = useState<EpgEvent[]>([]);

  const fetch = async (id: number) => {
    setLoading(true);
    let eventRes: EpgEvent;
    try {
      eventRes = await getEpgEvent(id);
      setEvent(eventRes);
    } catch (error) {
      if (error instanceof ApiError && error.code === 404) {
        setError(t('not_found'));
      } else {
        setError(t('unexpected'));
      }
      setLoading(false);
      return;
    }

    try {
      const related = await getRelatedEpgEvents(id);
      setRelatedEvents(related.entries.filter((r) => r.id !== eventRes.id));
    } catch (error) {
      setError(t('unexpected'));
    }

    setLoading(false);
  };

  return { error, loading, event, relatedEvents, fetch };
};
