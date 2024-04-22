import * as Yup from 'yup';

import Dropdown, { Option } from '@/components/common/dropdown/Dropdown';
import { Theme, useTheme } from '@/contexts/ThemeContext';
import { useLoaderData, useNavigate } from 'react-router-dom';

import { AuthInfo } from '@/clients/api/api.types';
import Button from '@/components/common/button/Button';
import Error from '@/components/common/error/Error';
import Form from '@/components/common/form/Form';
import Headline from '@/components/common/headline/Headline';
import Input from '@/components/common/input/Input';
import { TestIds } from '@/__test__/ids';
import { getAuthInfo } from '@/clients/api/api';
import i18n from 'i18next';
import moment from 'moment/min/moment-with-locales';
import styles from './SettingsView.module.scss';
import { useAuth } from '@/contexts/AuthContext';
import { useFormik } from 'formik';
import useFormikErrorFocus from '@/hooks/formik';
import { useRef } from 'react';
import { useTranslation } from 'react-i18next';
import { useUpdateUser } from '@/hooks/user';

export async function loader() {
  return Promise.all([getAuthInfo()]);
}

export const Component = () => {
  const { t } = useTranslation();
  const { user } = useAuth();
  const navigate = useNavigate();

  const { setTheme, theme } = useTheme();
  const { update } = useUpdateUser();

  const [authInfo] = useLoaderData() as [AuthInfo];

  const themeOptions: Option[] = [
    { title: t('dark'), value: 'dark' },
    { title: t('light'), value: 'light' },
  ];

  const languageOptions: Option[] = [
    { title: t('english'), value: 'en' },
    { title: t('german'), value: 'de' },
    { title: t('spanish'), value: 'es' },
  ];

  const timeFormatOptions: Option[] = [
    {
      title: t(`default`),
      value: 'default',
    },
    ...moment.locales().map((locale) => ({
      title: t(`locales.${locale}`),
      value: locale,
    })),
  ];

  const handleChangeLanguage = (lang: string) => {
    i18n.changeLanguage(lang, () => window.location.reload());
  };

  const handleChangeTimeFormat = (locale: string) => {
    if (locale === 'default') {
      localStorage.removeItem('time_locale');
    } else {
      localStorage.setItem('time_locale', locale);
    }

    window.location.reload();
  };

  const usernameRef = useRef<HTMLInputElement>(null);
  const emailRef = useRef<HTMLInputElement>(null);
  const displayNameRef = useRef<HTMLInputElement>(null);

  const userSettingsValidationSchema = Yup.object().shape({
    displayName: Yup.string().required(t('input_required')),
    email: Yup.string().required(t('input_required')),
    username: Yup.string().required(t('input_required')),
  });

  const userSettingsFormik = useFormik({
    enableReinitialize: true,
    initialValues: {
      displayName: user?.displayName || '',
      email: user?.email || '',
      username: user?.username || '',
    },
    onSubmit: async ({ username, email, displayName }) => {
      await update({ displayName, email, username });
    },
    validationSchema: userSettingsValidationSchema,
  });

  useFormikErrorFocus(
    userSettingsFormik,
    usernameRef,
    emailRef,
    displayNameRef
  );

  if (!user) {
    return <Error message={t('unexpected')} />;
  }

  return (
    <>
      <div className={styles.row}>
        <Form
          onSubmit={userSettingsFormik.handleSubmit}
          className={styles.section}
        >
          <Headline>{t('user')}</Headline>
          <Input
            label={t('username')}
            placeholder={t('username')}
            value={userSettingsFormik.values.username}
            name="username"
            onChange={userSettingsFormik.handleChange}
            onBlur={userSettingsFormik.handleBlur}
            error={
              userSettingsFormik.touched.username
                ? userSettingsFormik.errors.username
                : undefined
            }
            ref={usernameRef}
            fullWidth
          />
          <Input
            label={t('email')}
            placeholder={t('email')}
            value={userSettingsFormik.values.email}
            name="email"
            onChange={userSettingsFormik.handleChange}
            onBlur={userSettingsFormik.handleBlur}
            error={
              userSettingsFormik.touched.email
                ? userSettingsFormik.errors.email
                : undefined
            }
            ref={emailRef}
            fullWidth
          />
          <Input
            label={t('display_name')}
            placeholder={t('display_name')}
            value={userSettingsFormik.values.displayName}
            name="displayName"
            onChange={userSettingsFormik.handleChange}
            onBlur={userSettingsFormik.handleBlur}
            error={
              userSettingsFormik.touched.displayName
                ? userSettingsFormik.errors.displayName
                : undefined
            }
            ref={displayNameRef}
            fullWidth
          />
          <div>
            <Button
              type="submit"
              label={t('save')}
              testID={TestIds.SAVE_USER_BUTTON}
            />
          </div>
        </Form>
      </div>
      <div className={styles.row}>
        <div className={styles.section}>
          <Headline>{t('user_interface')}</Headline>
          <Dropdown
            label={t('appearance')}
            value={theme}
            options={themeOptions}
            onChange={(theme) => setTheme(theme as Theme)}
            fullWidth
            testID={TestIds.THEME_DROPDOWN}
          />
          <Dropdown
            label={t('language')}
            value={i18n.language}
            options={languageOptions}
            onChange={handleChangeLanguage}
            fullWidth
            testID={TestIds.LANGUAGE_DROPDOWN}
          />
          <Dropdown
            label={t('time_format')}
            value={localStorage.getItem('time_locale') || t('default')}
            options={timeFormatOptions}
            onChange={handleChangeTimeFormat}
            fullWidth
            testID={TestIds.TIME_FORMAT_DROPDOWN}
          />
          {!authInfo.forwardAuth ? (
            <div>
              <Button
                label={t('logout')}
                style="red"
                onClick={() => navigate('/logout')}
                testID={TestIds.LOGOUT_BUTTON}
              />
            </div>
          ) : null}
        </div>
      </div>
    </>
  );
};

Component.displayName = 'GeneralSettingsView';
