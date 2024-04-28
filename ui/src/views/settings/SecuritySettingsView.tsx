import * as Yup from 'yup';

import {
  Outlet,
  useLoaderData,
  useLocation,
  useNavigate,
  useRevalidator,
} from 'react-router-dom';
import { Session, Token, TwoFactorAuthSettings } from '@/clients/api/api.types';
import {
  getSessionsForCurrentUser,
  getTokens,
  getTwoFactorAuthSettings,
} from '@/clients/api/api';
import { useEffect, useRef } from 'react';

import Button from '@/components/common/button/Button';
import Form from '@/components/common/form/Form';
import Headline from '@/components/common/headline/Headline';
import Input from '@/components/common/input/Input';
import { SecuritySettingsRefreshStates } from './states';
import SessionList from '@/components/settings/sessionList/SessionList';
import TokenList from '@/components/settings/tokenList/TokenList';
import TwoFactorAuthSettingsOverview from '@/components/settings/twoFactorAuthSettings/TwoFactorAuthSettingsOverview';
import styles from './SettingsView.module.scss';
import { useFormik } from 'formik';
import useFormikErrorFocus from '@/hooks/formik';
import { useManageSessions } from '@/hooks/session';
import { useManageTokens } from '@/hooks/token';
import { useTranslation } from 'react-i18next';
import { useUpdateUserPassword } from '@/hooks/user';

export async function loader() {
  return Promise.all([
    getTwoFactorAuthSettings(),
    getSessionsForCurrentUser(),
    getTokens(),
  ]);
}

export const Component = () => {
  const { t } = useTranslation();
  const { revalidate } = useRevalidator();
  const navigate = useNavigate();
  const { state } = useLocation();
  const { revokeSession } = useManageSessions();
  const { revokeToken } = useManageTokens();
  const { updatePassword } = useUpdateUserPassword();

  const [twoFactorAuthSettings, sessions, tokens] = useLoaderData() as [
    TwoFactorAuthSettings,
    Array<Session>,
    Array<Token>,
  ];

  const currentPasswordRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);
  const passwordRepeatRef = useRef<HTMLInputElement>(null);

  const passwordChangeValidationSchema = Yup.object().shape({
    currentPassword: Yup.string().required(t('input_required')),
    password: Yup.string()
      .required(t('input_required'))
      .min(8, t('password_min_chars_error')),
    passwordRepeat: Yup.string()
      .required(t('passwords_do_not_match'))
      .oneOf([Yup.ref('password')], t('passwords_do_not_match')),
  });

  const passwordChangeFormik = useFormik({
    initialValues: {
      currentPassword: '',
      password: '',
      passwordRepeat: '',
    },
    onSubmit: async ({ password, currentPassword }) => {
      updatePassword({ currentPassword, password }).then((success) => {
        // TODO: can this be done better
        currentPasswordRef.current?.blur();
        passwordRef.current?.blur();
        passwordRepeatRef.current?.blur();

        if (success) {
          passwordChangeFormik.resetForm();
        }
      });
    },
    validationSchema: passwordChangeValidationSchema,
  });

  useFormikErrorFocus(
    passwordChangeFormik,
    currentPasswordRef,
    passwordRef,
    passwordRepeatRef
  );

  useEffect(() => {
    switch (state) {
      case SecuritySettingsRefreshStates.TWOFA:
        revalidate();
        break;
      case SecuritySettingsRefreshStates.TOKEN:
        revalidate();
        break;
    }
  }, [state, revalidate]);

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
      <div className={styles.row}>
        <TokenList tokens={tokens} onRevoke={revokeToken} />
      </div>
      <Outlet />
    </>
  );
};

Component.displayName = 'SecuritySettingsView';
