import * as Yup from 'yup';

import Button from '@/components/common/button/Button';
import Form from '@/components/common/form/Form';
import FormGroup from '@/components/common/form/FormGroup/FormGroup';
import Input from '@/components/common/input/Input';
import Modal from '@/components/common/modal/Modal';
import { SecuritySettingsRefreshStates } from '@/views/settings/states';
import styles from './TwoFactorAuthDisableModal.module.scss';
import { useDeactivateTwoFactorAuth } from '@/hooks/2fa';
import { useFormik } from 'formik';
import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

export const Component = () => {
  const { deactivateTwoFactorAuth, loading } = useDeactivateTwoFactorAuth();
  const { t } = useTranslation();
  const navigate = useNavigate();

  const close = (refresh = false) => {
    formik.resetForm();
    navigate('/settings/security', {
      state: refresh ? SecuritySettingsRefreshStates.TWOFA : undefined,
      preventScrollReset: true,
    });
  };

  const validationSchema = Yup.object().shape({
    code: Yup.string().required(t('two_factor_code_required')),
  });

  const formik = useFormik({
    initialValues: {
      code: '',
    },
    validationSchema,
    onSubmit: ({ code }) => {
      deactivateTwoFactorAuth(code).then(() => close(true));
    },
  });

  return (
    <Modal
      visible
      onClose={() => close()}
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

Component.displayName = 'TwoFactorAuthDisableModal';
