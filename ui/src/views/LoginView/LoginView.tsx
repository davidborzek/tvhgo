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
import FormGroup from '../../components/Form/FormGroup/FormGroup';

const GITHUB_URL = 'https://github.com/davidborzek/tvhgo';

export default function LoginView() {
  const { t } = useTranslation();
  const { login, loading, twoFactorRequired } = useLogin();

  const usernameRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);

  const loginSchema = Yup.object().shape({
    username: Yup.string().required(t('username_required') || ''),
    password: Yup.string().required(t('password_required') || ''),
  });

  const loginFormik = useFormik({
    initialValues: {
      username: '',
      password: '',
    },
    validationSchema: loginSchema,
    onSubmit: ({ username, password }) => login(username, password),
  });

  useFormikErrorFocus(loginFormik, usernameRef, passwordRef);

  const twoFactorSchema = Yup.object().shape({
    code: Yup.string().required(t('two_factor_code_required') || ''),
  });

  const twoFactorFormik = useFormik({
    initialValues: {
      code: '',
    },
    validationSchema: twoFactorSchema,
    onSubmit: ({ code }) =>
      login(loginFormik.values.username, loginFormik.values.password, code),
  });

  const renderTwoFactorForm = () => {
    return (
      <FormCard onSubmit={twoFactorFormik.handleSubmit}>
        <FormGroup
          info={t('two_factor_auth_info')}
          direction="column"
        >
          <Input
            type="text"
            name="code"
            label={t('verification_code')}
            value={twoFactorFormik.values.code}
            onBlur={twoFactorFormik.handleBlur}
            onChange={twoFactorFormik.handleChange}
            error={
              twoFactorFormik.touched.code
                ? twoFactorFormik.errors.code
                : undefined
            }
            fullWidth
          />
          <Button
            label={t('login')}
            type="submit"
            loading={loading}
            loadingLabel={t('login_pending')}
          />
        </FormGroup>
      </FormCard>
    );
  };

  const renderLoginForm = () => {
    return (
      <FormCard onSubmit={loginFormik.handleSubmit}>
        <Input
          type="text"
          name="username"
          label={t('username')}
          value={loginFormik.values.username}
          onBlur={loginFormik.handleBlur}
          onChange={loginFormik.handleChange}
          error={
            loginFormik.touched.username
              ? loginFormik.errors.username
              : undefined
          }
          ref={usernameRef}
          fullWidth
        />
        <Input
          type="password"
          name="password"
          label={t('password')}
          value={loginFormik.values.password}
          onBlur={loginFormik.handleBlur}
          onChange={loginFormik.handleChange}
          error={
            loginFormik.touched.password
              ? loginFormik.errors.password
              : undefined
          }
          ref={passwordRef}
          fullWidth
        />
        <Button
          label={t('login')}
          type="submit"
          loading={loading}
          loadingLabel={t('login_pending')}
        />
      </FormCard>
    );
  };

  return (
    <div className={styles.Login}>
      {twoFactorRequired ? renderTwoFactorForm() : renderLoginForm()}
      <LoginFooter
        commitHash={__COMMIT_HASH__}
        githubUrl={GITHUB_URL}
        version={__VERSION__}
      />
    </div>
  );
}
