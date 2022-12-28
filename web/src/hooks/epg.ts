import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { getEpgEvents, GetEpgEventsQuery } from "../clients/api/api";
import { EpgEvent } from "../clients/api/api.types";

export const useFetchEpg = (q?: GetEpgEventsQuery) => {
  const initialOffset = q?.offset || 0;

  const { t } = useTranslation("errors");

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
        setError(t("unexpected"));
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
