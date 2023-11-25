import { useNavigate } from 'react-router-dom';
import Modal from '../../../components/Modal/Modal';
import * as Yup from 'yup';

import styles from './CreateTokenModal.module.scss';
import { useTranslation } from 'react-i18next';
import { useFormik } from 'formik';
import Form from '../../../components/Form/Form';
import FormGroup from '../../../components/Form/FormGroup/FormGroup';
import Input from '../../../components/Input/Input';
import Button from '../../../components/Button/Button';
import { useCreateToken } from '../../../hooks/token';

const CreateTokenModal = () => {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const { createToken, setToken, token } = useCreateToken();

  const validationSchema = Yup.object().shape({
    name: Yup.string().required(t('name_required') || ''),
  });

  const formik = useFormik({
    initialValues: {
      name: '',
    },
    validationSchema,
    onSubmit: ({ name }) => {
      createToken(name);
    },
  });

  const close = () => {
    setToken('');
    formik.resetForm();
    navigate('/settings/security');
  };

  const renderToken = () => {
    return (
      <Form>
        <FormGroup direction="column" info={t('api_token_created_info')}>
          <Input
            name="token"
            label={t('token')}
            value={token}
            fullWidth
            showCopyButton
            selectTextOnFocus
            ellipsis
          />
          <Button
            disabled={false}
            label={t('close')}
            type="button"
            onClick={close}
          />
        </FormGroup>
      </Form>
    );
  };

  const renderForm = () => {
    return (
      <Form onSubmit={formik.handleSubmit}>
        <FormGroup direction="column" info={t('api_token_create_info')}>
          <Input
            name="name"
            label={t('name')}
            value={formik.values.name}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
            error={formik.touched.name ? formik.errors.name : undefined}
            fullWidth
          />
          <Button disabled={false} label={t('create')} type="submit" />
        </FormGroup>
      </Form>
    );
  };

  return (
    <Modal disableBackdropClose visible onClose={close} maxWidth="30rem">
      <div className={styles.content}>
        <h3 className={styles.headline}>{t('api_tokens')}</h3>
        {token ? renderToken() : renderForm()}
      </div>
    </Modal>
  );
};

export default CreateTokenModal;
