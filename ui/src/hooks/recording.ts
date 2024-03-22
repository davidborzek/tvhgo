import { useState } from 'react';
import { useTranslation } from 'react-i18next';
import {
  cancelRecording,
  cancelRecordings,
  recordByEvent,
  removeRecording,
  removeRecordings,
  stopRecording,
  stopRecordings,
  updateRecording,
} from '@/clients/api/api';
import { UpdateRecording } from '@/clients/api/api.types';
import { useNotification } from './notification';

export const useManageRecordingByEvent = () => {
  const { notifyError, notifySuccess, dismissNotification } = useNotification(
    'manageRecordingByEvent'
  );

  const { t } = useTranslation();
  const [pending, setPending] = useState(false);

  const createRecording = async (eventId: number) => {
    dismissNotification();
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
    dismissNotification();
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
    dismissNotification();
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
    dismissNotification();
    setPending(true);
    return await removeRecording(id)
      .then(() => {
        notifySuccess(t('recording_removed'));
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
    dismissNotification();
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

export const useManageRecordings = () => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('manageRecordings');

  const { t } = useTranslation();
  const [pending, setPending] = useState(false);

  const stopAndCancelRecordings = async (
    stopIds: string[],
    cancelIds: string[]
  ) => {
    dismissNotification();
    setPending(true);

    const promises = [];
    if (stopIds.length > 0) {
      promises.push(stopRecordings(stopIds));
    }

    if (cancelIds.length > 0) {
      promises.push(cancelRecordings(cancelIds));
    }

    return await Promise.all(promises)
      .then(() => {
        notifySuccess(t('recordings_stopped_canceled'));
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => {
        setPending(false);
      });
  };

  const _removeRecordings = async (ids: string[]) => {
    dismissNotification();
    setPending(true);

    return await removeRecordings(ids)
      .then(() => {
        notifySuccess(t('recordings_removed'));
      })
      .catch(() => {
        notifyError(t('unexpected'));
      })
      .finally(() => {
        setPending(false);
      });
  };

  return {
    stopAndCancelRecordings,
    removeRecordings: _removeRecordings,
    pending,
  };
};
