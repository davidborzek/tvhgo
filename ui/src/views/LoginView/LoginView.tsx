import { useFormik } from 'formik';
import useLogin from '../../hooks/login';
import * as Yup from 'yup';

import styles from './LoginView.module.scss';
import { useTranslation } from 'react-i18next';
import Input from '../../components/Input/Input';
import Button from '../../components/Button/Button';
import FormCard from '../../components/FormCard/FormCard';
import { useRef } from 'react';
import useFormikErrorFocus from '../../hooks/formik';
import LoginFooter from '../../components/LoginFooter/LoginFooter';

const GITHUB_URL = 'https://github.com/davidborzek/tvhgo';

export default function LoginView() {
  const { t } = useTranslation();
  const { login, loading } = useLogin();

  const usernameRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);

  const loginSchema = Yup.object().shape({
    username: Yup.string().required(t('username_required') || ''),
    password: Yup.string().required(t('password_required') || ''),
  });

  const formik = useFormik({
    initialValues: {
      username: '',
      password: '',
    },
    validationSchema: loginSchema,
    onSubmit: ({ username, password }) => login(username, password),
  });

  useFormikErrorFocus(formik, usernameRef, passwordRef);

  return (
    <div className={styles.Login}>
      <FormCard onSubmit={formik.handleSubmit}>
        <Input
          type="text"
          name="username"
          label={t('username')}
          value={formik.values.username}
          onBlur={formik.handleBlur}
          onChange={formik.handleChange}
          error={formik.touched.username ? formik.errors.username : undefined}
          ref={usernameRef}
        />
        <Input
          type="password"
          name="password"
          label={t('password')}
          value={formik.values.password}
          onBlur={formik.handleBlur}
          onChange={formik.handleChange}
          error={formik.touched.password ? formik.errors.password : undefined}
          ref={passwordRef}
        />
        <Button
          label={t('login')}
          type="submit"
          loading={loading}
          loadingLabel={t('login_pending')}
        />
      </FormCard>
      <LoginFooter
        commitHash={__COMMIT_HASH__}
        githubUrl={GITHUB_URL}
        version={__VERSION__}
      />
    </div>
  );
}
