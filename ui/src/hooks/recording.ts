import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { toast } from 'react-toastify';
import {
  cancelRecording,
  getRecordings,
  GetRecordingsQuery,
  recordByEvent,
  stopRecording,
} from '../clients/api/api';
import { Recording } from '../clients/api/api.types';

export const useManageRecordingByEvent = () => {
  const NOTIFICATION_ID = 'manageRecordingByEvent';

  const notifyError = (message?: string | null) => {
    toast.error(message, {
      toastId: NOTIFICATION_ID,
      updateId: NOTIFICATION_ID,
    });
  };

  const notifySuccess = (message?: string | null) => {
    toast.success(message, {
      toastId: NOTIFICATION_ID,
      updateId: NOTIFICATION_ID,
    });
  };

  const { t } = useTranslation();
  const [pending, setPending] = useState(false);

  const createRecording = async (eventId: number) => {
    setPending(true);
    return await recordByEvent(eventId)
      .then(() => {
        notifySuccess(t('recording_created'));
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => {
        setPending(false);
      });
  };

  const _cancelRecording = async (dvrId: string) => {
    setPending(true);
    return await cancelRecording(dvrId)
      .then(() => {
        notifySuccess(t('recording_canceled'));
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => {
        setPending(false);
      });
  };

  const _stopRecording = async (dvrId: string) => {
    setPending(true);
    return await stopRecording(dvrId)
      .then(() => {
        notifySuccess(t('recording_stopped'));
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => {
        setPending(false);
      });
  };

  return {
    createRecording,
    cancelRecording: _cancelRecording,
    stopRecording: _stopRecording,
    pending,
  };
};

export const useFetchRecordings = (q?: GetRecordingsQuery) => {
  const { t } = useTranslation();

  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [recordings, setRecordings] = useState<Recording[]>([]);
  const [status, setStatus] = useState(q?.status);

  const fetch = () => {
    setLoading(true);
    getRecordings({ ...q, status })
      .then(setRecordings)
      .catch(() => setError(t('unexpected')))
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    fetch();
  }, [status]);

  return { recordings, error, loading, fetch, setStatus, status };
};
