import { useRef } from 'react';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import { useTranslation } from 'react-i18next';

import useLogin from '@/hooks/login';
import Input from '@/components/common/input/Input';
import Button from '@/components/common/button/Button';
import LoginCard from '@/components/login/card/LoginCard';
import useFormikErrorFocus from '@/hooks/formik';
import LoginFooter from '@/components/login/footer/LoginFooter';
import FormGroup from '@/components/common/form/FormGroup/FormGroup';

import styles from './LoginView.module.scss';

const GITHUB_URL = 'https://github.com/davidborzek/tvhgo';

export default function LoginView() {
  const { t } = useTranslation();
  const { login, loading, twoFactorRequired } = useLogin();

  const usernameRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);

  const loginSchema = Yup.object().shape({
    username: Yup.string().required(t('username_required')),
    password: Yup.string().required(t('password_required')),
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
      <LoginCard onSubmit={twoFactorFormik.handleSubmit}>
        <FormGroup info={t('two_factor_auth_info')} direction="column">
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
      </LoginCard>
    );
  };

  const renderLoginForm = () => {
    return (
      <LoginCard onSubmit={loginFormik.handleSubmit}>
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
      </LoginCard>
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
