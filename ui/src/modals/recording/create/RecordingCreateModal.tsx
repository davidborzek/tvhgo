import * as Yup from 'yup';

import { Channel, DVRConfig } from '@/clients/api/api.types';
import { getChannels, getDVRConfigs } from '@/clients/api/api';
import { useLoaderData, useNavigate } from 'react-router-dom';

import Button from '@/components/common/button/Button';
import Dropdown from '@/components/common/dropdown/Dropdown';
import Form from '@/components/common/form/Form';
import FormGroup from '@/components/common/form/FormGroup/FormGroup';
import Input from '@/components/common/input/Input';
import Modal from '@/components/common/modal/Modal';
import { RecordingsViewRefreshStates } from '@/views/recordings/RecordingsView/RecordingsView';
import styles from './RecordingCreateModal.module.scss';
import { useCreateRecording } from '@/hooks/recording';
import { useFormik } from 'formik';
import { useTranslation } from 'react-i18next';

export const loader = async () => {
  return Promise.all([
    getChannels({
      // eslint-disable-next-line camelcase
      sort_key: 'name',
    }),
    getDVRConfigs(),
  ]);
};

export const Component = () => {
  const navigate = useNavigate();
  const { t } = useTranslation();

  const [createRecording, pending] = useCreateRecording();

  const close = (refresh?: boolean) => {
    const state = refresh ? RecordingsViewRefreshStates.CREATED : undefined;

    navigate('/dvr/recordings', { state });
  };

  const [channels, dvrProfiles] = useLoaderData() as [
    Array<Channel>,
    Array<DVRConfig>,
  ];

  const validationSchema = Yup.object().shape({
    title: Yup.string().required(t('title_required')),
    channel: Yup.string().required(t('channel_required')),
    profile: Yup.string().required(t('profile_required')),
    startsAt: Yup.string().required(t('starts_at_required')),
    endsAt: Yup.string().required(t('ends_at_required')),
    startPadding: Yup.number(),
    endPadding: Yup.number(),
  });

  const formik = useFormik({
    initialValues: {
      title: '',
      extraText: '',
      comment: '',
      channel: channels[0]?.id || '',
      profile: dvrProfiles[0]?.id || '',
      startsAt: '',
      endsAt: '',
      startPadding: 0,
      endPadding: 0,
    },
    onSubmit: ({
      title,
      extraText,
      comment,
      channel,
      profile,
      startsAt,
      endsAt,
      startPadding,
      endPadding,
    }) => {
      createRecording({
        title,
        extraText,
        comment,
        channelId: channel,
        configId: profile,
        startsAt: new Date(startsAt).getTime() / 1000,
        endsAt: new Date(endsAt).getTime() / 1000,
        startPadding,
        endPadding,
      }).then(() => {
        close(true);
      });
    },
    validationSchema,
  });

  return (
    <Modal
      disableBackdropClose
      disableEscapeClose
      visible
      onClose={close}
      maxWidth="25rem"
    >
      <div className={styles.content}>
        <h3 className={styles.headline}>{t('create_recording')}</h3>

        <Form onSubmit={formik.handleSubmit}>
          <FormGroup heading={t('details')} direction="column">
            <Input
              name="title"
              label={t('title')}
              placeholder={t('recording_title_placeholder')}
              value={formik.values.title}
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              error={formik.touched.title ? formik.errors.title : undefined}
              fullWidth
              disabled={pending}
            />
            <Input
              name="extraText"
              label={t('extra_text')}
              placeholder={t('recording_extra_text_placeholder')}
              value={formik.values.extraText}
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              fullWidth
              disabled={pending}
            />
            <Input
              label={t('comment')}
              value={formik.values.comment}
              placeholder={t('recording_comment_placeholder')}
              name="comment"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              fullWidth
              disabled={pending}
            />
            <Dropdown
              label={t('channel')}
              name="channel"
              value={formik.values.channel}
              options={channels.map((channel) => ({
                title: channel.name,
                value: channel.id,
              }))}
              onChange={(changed) => {
                formik.setFieldValue('channel', changed);
              }}
              fullWidth
            />
            <Dropdown
              label={t('profile')}
              name="profile"
              value={formik.values.profile}
              options={dvrProfiles.map((profile) => ({
                title: profile.name || t('default_profile'),
                value: profile.id,
              }))}
              onChange={(changed) => {
                formik.setFieldValue('profile', changed);
              }}
              fullWidth
            />
          </FormGroup>
          <FormGroup heading={t('recording_time')} direction="column">
            <Input
              name="startsAt"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              value={formik.values.startsAt}
              type="datetime-local"
              label={t('starts_at')}
              error={
                formik.touched.startsAt ? formik.errors.startsAt : undefined
              }
              fullWidth
              disabled={pending}
            />
            <Input
              type="datetime-local"
              label={t('ends_at')}
              value={formik.values.endsAt}
              name="endsAt"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={formik.touched.endsAt ? formik.errors.endsAt : undefined}
              fullWidth
              disabled={pending}
            />
          </FormGroup>
          <FormGroup
            info={t('recording_duration_padding_info')}
            direction="column"
          >
            <Input
              label={t('recording_minutes_before_start')}
              value={formik.values.startPadding}
              name="startPadding"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              type="number"
              selectTextOnFocus
              fullWidth
              disabled={pending}
            />
            <Input
              label={t('recording_minutes_after_end')}
              value={formik.values.endPadding}
              name="endPadding"
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              type="number"
              selectTextOnFocus
              fullWidth
              disabled={pending}
            />
          </FormGroup>
          <Button
            disabled={pending}
            label={t('create')}
            type="submit"
            loading={pending}
          />
        </Form>
      </div>
    </Modal>
  );
};

Component.displayName = 'RecordingCreateModal';
