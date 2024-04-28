import * as Yup from 'yup';

import Button from '@/components/common/button/Button';
import Form from '@/components/common/form/Form';
import Input from '@/components/common/input/Input';
import Modal from '@/components/common/modal/Modal';
import { UserListRefreshStates } from '@/views/settings/users/states';
import styles from './UserCreateModal.module.scss';
import { useCreateUser } from '@/hooks/user';
import { useFormik } from 'formik';
import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

export const Component = () => {
  const navigate = useNavigate();
  const { t } = useTranslation();
  const { create } = useCreateUser();

  const validationSchema = Yup.object().shape({
    name: Yup.string().required(t('input_required')),
    email: Yup.string().required(t('input_required')).email(t('email_invalid')),
    username: Yup.string().required(t('input_required')),
    password: Yup.string()
      .required(t('input_required'))
      .min(8, t('password_min_chars_error')),
    passwordRepeat: Yup.string()
      .required(t('passwords_do_not_match'))
      .oneOf([Yup.ref('password')], t('passwords_do_not_match')),
  });

  const formik = useFormik({
    initialValues: {
      username: '',
      password: '',
      passwordRepeat: '',
      name: '',
      email: '',
    },
    validationSchema,
    onSubmit: ({ name, username, email, password }) => {
      create({ displayName: name, username, email, password }).then(() => {
        close(true);
      });
    },
  });

  const close = (refresh?: boolean) => {
    const state = refresh ? UserListRefreshStates.CREATED : undefined;

    formik.resetForm();
    navigate('/settings/users', { state });
  };

  return (
    <Modal
      disableBackdropClose
      disableEscapeClose
      visible
      onClose={close}
      maxWidth="25rem"
    >
      <div className={styles.content}>
        <h3 className={styles.headline}>{t('create_user')}</h3>

        <Form onSubmit={formik.handleSubmit}>
          <Input
            name="name"
            label={t('name')}
            placeholder={'John Doe'}
            value={formik.values.name}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
            error={formik.touched.name ? formik.errors.name : undefined}
            fullWidth
          />
          <Input
            name="username"
            label={t('username')}
            placeholder={'johndoe'}
            value={formik.values.username}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
            error={formik.touched.username ? formik.errors.username : undefined}
            fullWidth
          />
          <Input
            name="email"
            label={t('email')}
            placeholder={'mail@example.com'}
            value={formik.values.email}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
            error={formik.touched.email ? formik.errors.email : undefined}
            fullWidth
          />
          <Input
            name="password"
            label={t('password')}
            placeholder={'********'}
            type="password"
            value={formik.values.password}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
            error={formik.touched.password ? formik.errors.password : undefined}
            fullWidth
          />
          <Input
            label={t('password_repeat')}
            placeholder={'********'}
            value={formik.values.passwordRepeat}
            name="passwordRepeat"
            type="password"
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={
              formik.touched.passwordRepeat
                ? formik.errors.passwordRepeat
                : undefined
            }
            fullWidth
          />
          <Button disabled={false} label={t('create')} type="submit" />
        </Form>
      </div>
    </Modal>
  );
};

Component.displayName = 'UserCreateModal';
