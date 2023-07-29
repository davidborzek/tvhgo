import { useTranslation } from 'react-i18next';

import styles from './TwoFactorAuthDisableModal.module.scss';

import * as Yup from 'yup';
import { useFormik } from 'formik';
import { useDeactivateTwoFactorAuth } from '../../../hooks/2fa';
import Button from '../../../components/Button/Button';
import FormGroup from '../../../components/Form/FormGroup/FormGroup';
import Input from '../../../components/Input/Input';
import Modal from '../../../components/Modal/Modal';
import Form from '../../../components/Form/Form';

type Props = {
  visible: boolean;
  onClose: () => void;
  onFinish: () => void;
};

const TwoFactorAuthDisableModal = ({ visible, onClose, onFinish }: Props) => {
  const { deactivateTwoFactorAuth, loading } = useDeactivateTwoFactorAuth();
  const { t } = useTranslation();

  const close = () => {
    formik.resetForm();
    onClose();
  };

  const validationSchema = Yup.object().shape({
    code: Yup.string().required(t('verification_code_required') || ''),
  });

  const formik = useFormik({
    initialValues: {
      code: '',
    },
    validationSchema,
    onSubmit: ({ code }) => {
      deactivateTwoFactorAuth(code).then(() => {
        onFinish();
        close();
      });
    },
  });

  return (
    <Modal
      onClose={close}
      visible={visible}
      maxWidth="30rem"
      disableBackdropClose
    >
      <div className={styles.content}>
        <h3 className={styles.headline}>{t('disable_two_factor_auth')}</h3>
        <Form onSubmit={formik.handleSubmit}>
          <FormGroup
            direction="column"
            info={t('two_factor_auth_disable_info')}
          >
            <Input
              name="code"
              label={t('verification_code')}
              value={formik.values.code}
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              error={formik.touched.code ? formik.errors.code : undefined}
              fullWidth
            />
            <Button
              disabled={loading}
              label={t('disable')}
              style="red"
              type="submit"
            />
          </FormGroup>
        </Form>
      </div>
    </Modal>
  );
};

export default TwoFactorAuthDisableModal;
