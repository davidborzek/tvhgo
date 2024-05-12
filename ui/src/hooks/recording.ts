import { CreateRecordingOpts, UpdateRecording } from '@/clients/api/api.types';
import {
  cancelRecording,
  cancelRecordings,
  createRecording,
  recordByEvent,
  removeRecording,
  removeRecordings,
  stopRecording,
  stopRecordings,
  updateRecording,
} from '@/clients/api/api';

import { useNotification } from './notification';
import { useState } from 'react';
import { useTranslation } from 'react-i18next';

export const useManageRecordingByEvent = () => {
  const { notifyError, notifySuccess, dismissNotification } = useNotification(
    'manageRecordingByEvent'
  );

  const { t } = useTranslation();
  const [pending, setPending] = useState(false);

  const createRecording = async (eventId: number, configId?: string) => {
    dismissNotification();
    setPending(true);
    return await recordByEvent(eventId, configId)
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
        if (success) success();
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
        if (success) success();
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
        if (success) success();
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

export const useCreateRecording = (): [
  (opts: CreateRecordingOpts) => Promise<void>,
  boolean,
] => {
  const { notifyError, notifySuccess, dismissNotification } =
    useNotification('createRecording');

  const { t } = useTranslation();
  const [pending, setPending] = useState(false);

  const _createRecording = (opts: CreateRecordingOpts) => {
    dismissNotification();
    setPending(true);

    return createRecording(opts)
      .then(() => {
        notifySuccess(t('recording_created_successfully'));
      })
      .catch((err) => {
        notifyError(t('unexpected'));
        throw err;
      })
      .finally(() => {
        setPending(false);
      });
  };

  return [_createRecording, pending];
};
