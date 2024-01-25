import { useTranslation } from 'react-i18next';
import { useRef } from 'react';
import * as Yup from 'yup';
import { useFormik } from 'formik';

import styles from './SettingsView.module.scss';
import Button from '@/components/common/button/Button';
import Dropdown, { Option } from '@/components/common/dropdown/Dropdown';
import Input from '@/components/common/input/Input';
import { useAuth } from '@/contexts/AuthContext';
import { useTheme, Theme } from '@/contexts/ThemeContext';
import useFormikErrorFocus from '@/hooks/formik';
import { useUpdateUser } from '@/hooks/user';
import i18n from 'i18next';
import { useNavigate } from 'react-router-dom';
import Form from '@/components/common/form/Form';
import { TestIds } from '@/__test__/ids';

const GeneralSettingsView = () => {
  const { t } = useTranslation();
  const { user } = useAuth();
  const navigate = useNavigate();

  const { setTheme, theme } = useTheme();
  const { update } = useUpdateUser();

  const themeOptions: Option[] = [
    { title: t('dark'), value: 'dark' },
    { title: t('light'), value: 'light' },
  ];

  const languageOptions: Option[] = [
    { title: t('english'), value: 'en' },
    { title: t('german'), value: 'de' },
  ];

  const handleChangeLanguage = (lang: string) => {
    i18n.changeLanguage(lang, () => window.location.reload());
  };

  const usernameRef = useRef<HTMLInputElement>(null);
  const emailRef = useRef<HTMLInputElement>(null);
  const displayNameRef = useRef<HTMLInputElement>(null);

  const userSettingsValidationSchema = Yup.object().shape({
    username: Yup.string().required(t('input_required')),
    email: Yup.string().required(t('input_required')),
    displayName: Yup.string().required(t('input_required')),
  });

  const userSettingsFormik = useFormik({
    initialValues: {
      username: user?.username || '',
      email: user?.email || '',
      displayName: user?.displayName || '',
    },
    validationSchema: userSettingsValidationSchema,
    onSubmit: async ({ username, email, displayName }) => {
      await update({ username, email, displayName });
    },
    enableReinitialize: true,
  });

  useFormikErrorFocus(
    userSettingsFormik,
    usernameRef,
    emailRef,
    displayNameRef
  );

  return (
    <>
      <div className={styles.row}>
        <Form
          onSubmit={userSettingsFormik.handleSubmit}
          className={styles.section}
        >
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
          <div>
            <Button
              label={t('logout')}
              style="red"
              onClick={() => navigate('/logout')}
              testID={TestIds.LOGOUT_BUTTON}
            />
          </div>
        </div>
      </div>
    </>
  );
};

export default GeneralSettingsView;
