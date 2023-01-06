import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { toast } from 'react-toastify';
import {
  ApiError,
  cancelRecording,
  getRecording,
  getRecordings,
  GetRecordingsQuery,
  recordByEvent,
  removeRecording,
  stopRecording,
  updateRecording,
} from '../clients/api/api';
import { Recording, UpdateRecording } from '../clients/api/api.types';

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

  const _cancelRecording = async (id: string, success?: () => void) => {
    setPending(true);
    return await cancelRecording(id)
      .then(() => {
        notifySuccess(t('recording_canceled'));
        success && success();
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => {
        setPending(false);
      });
  };

  const _stopRecording = async (id: string, success?: () => void) => {
    setPending(true);
    return await stopRecording(id)
      .then(() => {
        notifySuccess(t('recording_stopped'));
        success && success();
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => {
        setPending(false);
      });
  };

  const _removeRecording = async (id: string, success?: () => void) => {
    setPending(true);
    return await removeRecording(id)
      .then(() => {
        notifySuccess(t('recording_stopped'));
        success && success();
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => {
        setPending(false);
      });
  };

  const _updateRecording = async (id: string, opts: UpdateRecording) => {
    setPending(true);
    return await updateRecording(id, opts)
      .then(() => {
        notifySuccess(t('recording_updated'));
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
    removeRecording: _removeRecording,
    updateRecording: _updateRecording,
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

export const useFetchRecording = () => {
  const { t } = useTranslation();

  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [recording, setRecording] = useState<Recording>();

  const fetch = async (id: string) => {
    setLoading(true);
    getRecording(id)
      .then(setRecording)
      .catch((error) => {
        if (error instanceof ApiError && error.code === 404) {
          setError(t('not_found'));
        } else {
          setError(t('unexpected'));
        }
      })
      .finally(() => setLoading(false));
  };

  return { error, loading, recording, fetch };
};
