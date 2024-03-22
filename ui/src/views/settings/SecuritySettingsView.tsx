import { useFormik } from 'formik';
import { useEffect, useRef } from 'react';
import { useTranslation } from 'react-i18next';
import { useNavigate, useLocation, Outlet, useLoaderData, useRevalidator } from 'react-router-dom';
import * as Yup from 'yup';

import Button from '@/components/common/button/Button';
import Headline from '@/components/common/headline/Headline';
import Input from '@/components/common/input/Input';
import useFormikErrorFocus from '@/hooks/formik';
import { useManageSessions } from '@/hooks/session';
import { useManageTokens } from '@/hooks/token';
import { useUpdateUserPassword } from '@/hooks/user';
import Form from '@/components/common/form/Form';

import styles from './SettingsView.module.scss';
import SessionList from '@/components/settings/sessionList/SessionList';
import TokenList from '@/components/settings/tokenList/TokenList';
import TwoFactorAuthSettingsOverview from '@/components/settings/twoFactorAuthSettings/TwoFactorAuthSettingsOverview';
import {
  getSessions,
  getTokens,
  getTwoFactorAuthSettings,
} from '@/clients/api/api';
import { Token, TwoFactorAuthSettings } from '@/clients/api/api.types';
import { Session } from '@/clients/api/api.types';

export enum SecuritySettingsRefreshStates {
  TWOFA = 'refresh_2fa',
  TOKEN = 'refresh_token',
}

export async function loader() {
  return Promise.all([getTwoFactorAuthSettings(), getSessions(), getTokens()]);
}

export const Component = () => {
  const { t } = useTranslation();
  const revalidator = useRevalidator();
  const navigate = useNavigate();
  const { state } = useLocation();
  const { revokeSession } = useManageSessions();
  const { revokeToken } = useManageTokens();
  const { updatePassword } = useUpdateUserPassword();

  const [twoFactorAuthSettings, sessions, tokens] = useLoaderData() as [TwoFactorAuthSettings, Array<Session>, Array<Token>]

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
    switch (state) {
      case SecuritySettingsRefreshStates.TWOFA:
        revalidator.revalidate();
        break;
      case SecuritySettingsRefreshStates.TOKEN:
        revalidator.revalidate();
        break;
    }
  }, [state]);

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
