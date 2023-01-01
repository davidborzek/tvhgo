import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { toast } from 'react-toastify';
import {
  cancelRecording,
  recordByEvent,
  stopRecording,
} from '../clients/api/api';

const NOTIFICATION_ID = 'createRecordingByEvent';

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

export const useManageRecordingByEvent = () => {
  const { t } = useTranslation(['errors', 'recording']);
  const [pending, setPending] = useState(false);

  const createRecording = async (eventId: number) => {
    setPending(true);
    return await recordByEvent(eventId)
      .then(() => {
        notifySuccess(t('recording_created', { ns: "recording" }));
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
        notifySuccess(t('recording_canceled', { ns: "recording" }));
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
        notifySuccess(t('recording_stopped', { ns: "recording" }));
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
