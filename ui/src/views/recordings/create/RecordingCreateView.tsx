import Button from '@/components/common/button/Button';
import Form from '@/components/common/form/Form';
import FormGroup from '@/components/common/form/FormGroup/FormGroup';
import Input from '@/components/common/input/Input';
import { useTranslation } from 'react-i18next';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';
import * as Yup from 'yup';

import styles from './RecordingCreateView.module.scss';
import { useCreateRecording } from '@/hooks/recording';
import { useFormik } from 'formik';
import { useNotification } from '@/hooks/notification';

const RecordingCreateView = () => {
  const { t } = useTranslation();
  const { state } = useLocation();
  const navigate = useNavigate();
  const { createRecording, pending } = useCreateRecording();

  const { notifyError, dismissNotification } = useNotification(
    'formikCreateRecording'
  );

  const validationSchema = Yup.object().shape({
    title: Yup.string().required(t('title_required')),
    channel: Yup.string().required(t('channel_required')),
    startsAt: Yup.date().required(t('starts_at_required')),
    endsAt: Yup.date().required(t('ends_at_required')),
    startPadding: Yup.number(),
    endPadding: Yup.number(),
  });

  const formik = useFormik({
    initialValues: {
      title: '',
      extraText: '',
      comment: '',
      startPadding: 0,
      endPadding: 0,
      startsAt: '',
      endsAt: '',
      channel: (state?.channel?.name as string) || '',
    },
    validateOnChange: false,
    validateOnBlur: false,
    validationSchema,
    onSubmit: (opts) => {
      createRecording({
        title: opts.title,
        extraText: opts.extraText,
        startsAt: new Date(opts.startsAt).getTime() / 1000,
        endsAt: new Date(opts.endsAt).getTime() / 1000,
        channelId: state?.channel?.id,
        startPadding: opts.startPadding,
        endPadding: opts.endPadding,
        comment: opts.comment,
      }).then(() => navigate('/recordings'));
    },
    enableReinitialize: true,
  });

  return (
    <div className={styles.root}>
      <h1>{t('create_recording')}</h1>
      <Form
        className={styles.form}
        onSubmit={(e) => {
          if (Object.keys(formik.errors).length > 0) {
            notifyError(Object.values(formik.errors)[0]);
          } else {
            dismissNotification();
          }

          formik.handleSubmit(e);
        }}
      >
        <FormGroup>
          <FormGroup direction="column">
            <Input
              label={t('title')}
              value={formik.values.title}
              name="title"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              disabled={pending}
              fullWidth
            />
            <Input
              label={t('extra_text')}
              value={formik.values.extraText}
              name="extraText"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              disabled={pending}
              fullWidth
            />
            <Input
              label={t('comment')}
              value={formik.values.comment}
              name="comment"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              disabled={pending}
              fullWidth
            />
            <Input
              label={t('channel')}
              name="channel"
              value={state?.channel?.name}
              onClick={() => navigate('select-channel')}
              onBlur={formik.handleBlur}
              disabled={pending}
              hideCarret
              fullWidth
              readonly
            />
          </FormGroup>

          <FormGroup direction="column">
            <Input
              name="startsAt"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              value={formik.values.startsAt}
              disabled={pending}
              type="datetime-local"
              label={t('starts_at')}
              fullWidth
            />
            <Input
              type="datetime-local"
              label={t('ends_at')}
              value={formik.values.endsAt}
              name="endsAt"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              disabled={pending}
              fullWidth
            />
            <Input
              label={t('recording_minutes_before_start')}
              value={formik.values.startPadding}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              disabled={pending}
              name="startPadding"
              type="number"
              fullWidth
            />
            <Input
              label={t('recording_minutes_after_end')}
              value={formik.values.endPadding}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              disabled={pending}
              name="endPadding"
              type="number"
              fullWidth
            />
          </FormGroup>
        </FormGroup>

        <FormGroup direction="column">
          <Button
            type="submit"
            label={t('create')}
            className={styles.button}
            loading={pending}
          />
        </FormGroup>
      </Form>

      <Outlet />
    </div>
  );
};

export default RecordingCreateView;
