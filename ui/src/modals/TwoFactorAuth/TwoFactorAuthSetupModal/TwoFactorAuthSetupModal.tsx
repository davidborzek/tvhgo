import { useFormik } from 'formik';
import { useTranslation } from 'react-i18next';
import QRCode from 'react-qr-code';
import Button from '../../../components/Button/Button';
import FormGroup from '../../../components/Form/FormGroup/FormGroup';
import Input from '../../../components/Input/Input';
import Modal from '../../../components/Modal/Modal';
import { useSetupTwoFactorAuth } from '../../../hooks/2fa';
import styles from './TwoFactorAuthSetupModal.module.scss';

import * as Yup from 'yup';
import Form from '../../../components/Form/Form';

type Props = {
  visible: boolean;
  onClose: () => void;
  onFinish: () => void;
};

const TwoFactorAuthSetupModal = ({ visible, onClose, onFinish }: Props) => {
  const { t } = useTranslation();

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
        onFinish();
        close();
      });
    },
  });

  const close = () => {
    onClose();
    setTwoFactorUrl(null);
    enableFormik.resetForm();
    setupFormik.resetForm();
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
      visible={visible}
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
