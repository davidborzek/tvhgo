import { useFormik } from 'formik';
import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import * as Yup from 'yup';
import { useTranslation } from 'react-i18next';

import Error from '@/components/Error/Error';
import Input from '@/components/Input/Input';
import {
  useFetchRecording,
  useManageRecordingByEvent,
} from '@/hooks/recording';
import Button from '@/components/Button/Button';
import FormGroup from '@/components/Form/FormGroup/FormGroup';
import EventChannelInfo from '@/components/Event/EventChannelInfo/EventChannelInfo';
import Form from '@/components/Form/Form';
import PairList from '@/components/PairList/PairList';
import Pair from '@/components/PairList/Pair/Pair';
import PairValue from '@/components/PairList/PairValue/PairValue';
import PairKey from '@/components/PairList/PairKey/PairKey';
import DeleteConfirmationModal from '@/components/DeleteConfirmationModal/DeleteConfirmationModal';
import { getRecordingUrl } from '@/clients/api/api';
import ButtonLink from '@/components/Button/ButtonLink';

import styles from './RecordingDetailView.module.scss';

function RecordingDetailView() {
  const { t } = useTranslation();
  const { id } = useParams();
  const navigate = useNavigate();

  const { recording, error, fetch } = useFetchRecording();
  const {
    cancelRecording,
    stopRecording,
    removeRecording,
    updateRecording,
    pending,
  } = useManageRecordingByEvent();

  const [confirmationModalVisible, setConfirmationModalVisible] =
    useState<boolean>(false);

  const parseStatus = () => {
    switch (recording?.status) {
      case 'scheduled':
        return t('recording_scheduled');
      case 'recording':
        return t('recording_running');
      case 'completed':
        return t('recording_completed');
    }
    return t('recording_unknown');
  };

  const getCancelButtonLabel = () => {
    switch (recording?.status) {
      case 'scheduled':
        return t('cancel_recording');
      case 'recording':
        return t('stop_recording');
      case 'completed':
        return t('delete_recording');
    }
    return '';
  };

  const getConfirmationButtonLabel = () => {
    switch (recording?.status) {
      case 'scheduled':
        return t('cancel');
      case 'recording':
        return t('stop');
      case 'completed':
        return t('delete');
    }
    return '';
  };

  const getConfirmationModalTitle = () => {
    switch (recording?.status) {
      case 'scheduled':
        return t('confirm_cancel_recording');
      case 'recording':
        return t('confirm_stop_recording');
      case 'completed':
        return t('confirm_delete_recording');
    }
    return '';
  };

  const validationSchema = Yup.object().shape({
    startPadding: Yup.number(),
    endPadding: Yup.number(),
  });

  const handleDeleteOrStop = () => {
    if (recording?.status === 'recording') {
      stopRecording(recording.id, () => {
        setConfirmationModalVisible(false);
        fetch(recording.id);
      });
    } else if (recording?.status === 'scheduled') {
      cancelRecording(recording?.id, () => {
        navigate(-1);
      });
    } else if (recording?.status === 'completed') {
      removeRecording(recording.id, () => {
        navigate(-1);
      });
    }
  };

  const formik = useFormik({
    initialValues: {
      startPadding: 0,
      endPadding: 0,
    },
    validationSchema,
    onSubmit: ({ endPadding, startPadding }) => {
      if (!recording) {
        return;
      }

      updateRecording(recording.id, {
        endPadding,
        startPadding,
      });
    },
  });

  useEffect(() => {
    if (id) {
      fetch(id);
    }
  }, [id]);

  useEffect(() => {
    if (recording) {
      formik.setValues({
        startPadding: recording.startPadding,
        endPadding: recording.endPadding,
      });
    }
  }, [recording]);

  const renderCancelButton = () => {
    if (
      recording?.status === 'recording' ||
      recording?.status === 'scheduled' ||
      recording?.status === 'completed'
    ) {
      return (
        <Button
          type="button"
          label={getCancelButtonLabel()}
          style="red"
          onClick={() => setConfirmationModalVisible(true)}
        />
      );
    }
  };

  const renderDownloadButton = () => {
    if (recording?.status === 'completed') {
      return (
        <ButtonLink
          href={getRecordingUrl(recording.id)}
          download
          label={t('download')}
        />
      );
    }
  };

  const renderTimeForm = () => {
    if (
      recording?.status === 'scheduled' ||
      recording?.status === 'recording'
    ) {
      return (
        <Form
          onSubmit={formik.handleSubmit}
          className={styles.timeForm}
          maxWidth="20rem"
        >
          <FormGroup
            heading={t('recording_time')}
            info={t('recording_duration_padding_info')}
            direction="column"
          >
            <Input
              label={t('recording_minutes_before_start')}
              value={formik.values.startPadding}
              name="startPadding"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              disabled={recording.status === 'recording' || pending}
              type="number"
              selectTextOnFocus
            />
            <Input
              label={t('recording_minutes_after_end')}
              value={formik.values.endPadding}
              name="endPadding"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              type="number"
              disabled={pending}
              selectTextOnFocus
            />
          </FormGroup>
          <Button type="submit" label={t('save')} disabled={pending} />
        </Form>
      );
    }
  };

  if (error) {
    return <Error message={error} />;
  }

  if (!recording) {
    return <></>;
  }

  return (
    <div className={styles.RecordingDetailView}>
      <DeleteConfirmationModal
        visible={confirmationModalVisible}
        onClose={() => setConfirmationModalVisible(false)}
        onConfirm={handleDeleteOrStop}
        title={getConfirmationModalTitle()}
        buttonTitle={getConfirmationButtonLabel()}
        pending={pending}
      />

      <EventChannelInfo
        channelName={recording.channelName}
        picon={`/api/picon/${recording.piconId}`}
      />
      <h2>{recording.title}</h2>
      <PairList>
        <Pair>
          <PairKey>{t('subtitle')}</PairKey>
          <PairValue>{recording.subtitle}</PairValue>
        </Pair>
        <Pair>
          <PairKey>{t('airs')}</PairKey>
          <PairValue>{t('event_datetime', { event: recording })}</PairValue>
        </Pair>
        <Pair>
          <PairKey>{t('status')}</PairKey>
          <PairValue>{parseStatus()}</PairValue>
        </Pair>
        <Pair>
          <PairKey>{t('description')}</PairKey>
          <PairValue>{recording.description}</PairValue>
        </Pair>
      </PairList>

      <div className={styles.actions}>
        {renderCancelButton()}
        {renderDownloadButton()}
      </div>

      {renderTimeForm()}
    </div>
  );
}

export default RecordingDetailView;
