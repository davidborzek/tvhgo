import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import {
  getEpgChannelEvents,
  GetEpgChannelEventsQuery,
  getEpgEvents,
  GetEpgEventsQuery,
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
