import { useFormik } from 'formik';
import { useTranslation } from 'react-i18next';
import * as Yup from 'yup';
import QRCode from 'react-qr-code';
import { useNavigate } from 'react-router-dom';

import Button from '@/components/Button/Button';
import FormGroup from '@/components/Form/FormGroup/FormGroup';
import Input from '@/components/Input/Input';
import Modal from '@/components/Modal/Modal';
import { useSetupTwoFactorAuth } from '@/hooks/2fa';
import Form from '@/components/Form/Form';

import styles from './TwoFactorAuthSetupModal.module.scss';

const TwoFactorAuthSetupModal = () => {
  const { t } = useTranslation();
  const navigate = useNavigate();

  const {
    activateTwoFactorAuth,
    setupTwoFactorAuth,
    loading,
    setTwoFactorUrl,
    twoFactorUrl,
  } = useSetupTwoFactorAuth();

  const setupSchema = Yup.object().shape({
    password: Yup.string().required(t('password_required') || ''),
  });

  const setupFormik = useFormik({
    initialValues: {
      password: '',
    },
    validationSchema: setupSchema,
    onSubmit: ({ password }) => {
      setupTwoFactorAuth(password);
    },
  });

  const enableSchema = Yup.object().shape({
    code: Yup.string().required(t('two_factor_code_required') || ''),
  });

  const enableFormik = useFormik({
    initialValues: {
      code: '',
    },
    validationSchema: enableSchema,
    onSubmit: ({ code }) => {
      activateTwoFactorAuth(setupFormik.values.password, code).then(() => {
        close();
      });
    },
  });

  const close = () => {
    setTwoFactorUrl(null);
    enableFormik.resetForm();
    setupFormik.resetForm();
    navigate('/settings/security');
  };

  const renderEnableTwoFactorForm = () => {
    return (
      <Form onSubmit={enableFormik.handleSubmit}>
        <div className={styles.qrCode}>
          <QRCode value={twoFactorUrl || ''} />
          <span>{twoFactorUrl}</span>
        </div>
        <FormGroup direction="column" info={t('two_factor_auth_info')}>
          <Input
            name="code"
            label={t('verification_code')}
            value={enableFormik.values.code}
            onBlur={enableFormik.handleBlur}
            onChange={enableFormik.handleChange}
            error={
              enableFormik.touched.code ? enableFormik.errors.code : undefined
            }
            fullWidth
          />
          <Button disabled={loading} label={t('enable')} type="submit" />
        </FormGroup>
      </Form>
    );
  };

  const renderSetupTwoFactorForm = () => {
    return (
      <Form onSubmit={setupFormik.handleSubmit}>
        <FormGroup direction="column" info={t('two_factor_auth_setup_info')}>
          <Input
            name="password"
            label={t('password')}
            type="password"
            value={setupFormik.values.password}
            onBlur={setupFormik.handleBlur}
            onChange={setupFormik.handleChange}
            error={
              setupFormik.touched.password
                ? setupFormik.errors.password
                : undefined
            }
            fullWidth
          />
          <Button disabled={loading} label={t('next')} type="submit" />
        </FormGroup>
      </Form>
    );
  };

  return (
    <Modal
      onClose={close}
      visible
      maxWidth="30rem"
      disableBackdropClose
    >
      <div className={styles.content}>
        <h3 className={styles.headline}>{t('enable_two_factor_auth')}</h3>
        {twoFactorUrl
          ? renderEnableTwoFactorForm()
          : renderSetupTwoFactorForm()}
      </div>
    </Modal>
  );
};

export default TwoFactorAuthSetupModal;
