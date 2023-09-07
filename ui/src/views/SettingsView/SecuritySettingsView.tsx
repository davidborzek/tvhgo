import { useFormik } from 'formik';
import { useEffect, useRef } from 'react';
import { useTranslation } from 'react-i18next';
import * as Yup from 'yup';
import useFormikErrorFocus from '../../hooks/formik';
import { useUpdateUserPassword } from '../../hooks/user';
import Input from '../../components/Input/Input';
import Button from '../../components/Button/Button';

import styles from './SettingsView.module.scss';
import Form from '../../components/Form/Form';
import { useManageSessions } from '../../hooks/session';
import SessionList from '../../components/SessionList/SessionList';
import { useTwoFactorAuthSettings } from '../../hooks/2fa';
import { usePromiseAll } from '../../hooks/async';
import Headline from '../../components/Headline/Headline';
import TwoFactorAuthSettingsOverview from '../../components/TwoFactorAuthSettingsOverview/TwoFactorAuthSettingsOverview';
import { Outlet, useLocation, useNavigate } from 'react-router-dom';

const SecuritySettingsView = () => {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const location = useLocation();
  const { sessions, getSessions, revokeSession } = useManageSessions();
  const { twoFactorAuthSettings, fetchTwoFactorAuthSettings } =
    useTwoFactorAuthSettings();
  const { updatePassword } = useUpdateUserPassword();

  usePromiseAll([getSessions, fetchTwoFactorAuthSettings]);
  const currentPasswordRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);
  const passwordRepeatRef = useRef<HTMLInputElement>(null);

  const passwordChangeValidationSchema = Yup.object().shape({
    currentPassword: Yup.string().required(t('input_required') || ''),
    password: Yup.string()
      .required(t('input_required') || '')
      .min(8, t('password_min_chars_error') || ''),
    passwordRepeat: Yup.string()
      .required(t('passwords_do_not_match') || '')
      .oneOf([Yup.ref('password')], t('passwords_do_not_match') || ''),
  });

  const passwordChangeFormik = useFormik({
    initialValues: {
      currentPassword: '',
      password: '',
      passwordRepeat: '',
    },
    validationSchema: passwordChangeValidationSchema,
    onSubmit: async ({ password, currentPassword }) => {
      updatePassword({ password, currentPassword }).then((success) => {
        // TODO: can this be done better
        currentPasswordRef.current?.blur();
        passwordRef.current?.blur();
        passwordRepeatRef.current?.blur();
        success && passwordChangeFormik.resetForm();
      });
    },
  });

  useFormikErrorFocus(
    passwordChangeFormik,
    currentPasswordRef,
    passwordRef,
    passwordRepeatRef
  );

  useEffect(() => {
    if (location.pathname === '/settings/security') {
      fetchTwoFactorAuthSettings();
    }
  }, [location.key]);

  return (
    <>
      <div className={styles.row}>
        <Form
          onSubmit={passwordChangeFormik.handleSubmit}
          className={styles.section}
        >
          <Headline>{t('password')}</Headline>
          <Input
            placeholder={t('current_password')}
            label={t('current_password')}
            value={passwordChangeFormik.values.currentPassword}
            name="currentPassword"
            type="password"
            onChange={passwordChangeFormik.handleChange}
            onBlur={passwordChangeFormik.handleBlur}
            error={
              passwordChangeFormik.touched.currentPassword
                ? passwordChangeFormik.errors.currentPassword
                : undefined
            }
            ref={currentPasswordRef}
            fullWidth
          />
          <Input
            placeholder={t('password')}
            label={t('password')}
            value={passwordChangeFormik.values.password}
            name="password"
            type="password"
            onChange={passwordChangeFormik.handleChange}
            onBlur={passwordChangeFormik.handleBlur}
            error={
              passwordChangeFormik.touched.password
                ? passwordChangeFormik.errors.password
                : undefined
            }
            ref={passwordRef}
            fullWidth
          />
          <Input
            label={t('password_repeat')}
            placeholder={t('password_repeat')}
            value={passwordChangeFormik.values.passwordRepeat}
            name="passwordRepeat"
            type="password"
            onChange={passwordChangeFormik.handleChange}
            onBlur={passwordChangeFormik.handleBlur}
            error={
              passwordChangeFormik.touched.passwordRepeat
                ? passwordChangeFormik.errors.passwordRepeat
                : undefined
            }
            ref={passwordRepeatRef}
            fullWidth
          />
          <div>
            <Button type="submit" label={t('save')} />
          </div>
        </Form>
      </div>
      <div className={styles.row}>
        <TwoFactorAuthSettingsOverview
          settings={twoFactorAuthSettings}
          onDisable={() =>
            navigate('/settings/security/two-factor-auth/disable')
          }
          onEnable={() => navigate('/settings/security/two-factor-auth/setup')}
        />
      </div>
      <div className={styles.row}>
        <SessionList sessions={sessions} onRevoke={revokeSession} />
      </div>
      <Outlet />
    </>
  );
};

export default SecuritySettingsView;