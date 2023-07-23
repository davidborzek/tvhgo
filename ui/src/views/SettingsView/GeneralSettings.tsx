import { useTranslation } from 'react-i18next';
import Dropdown, { Option } from '../../components/Dropdown/Dropdown';
import i18n from '../../i18n/i18n';
import { Theme, useTheme } from '../../contexts/ThemeContext';
import { useUpdateUser } from '../../hooks/user';
import { useEffect, useRef } from 'react';
import * as Yup from 'yup';
import styles from './SettingsView.module.scss';
import { useFormik } from 'formik';
import useFormikErrorFocus from '../../hooks/formik';
import { useAuth } from '../../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';
import Form from '../../components/Form/Form';
import Input from '../../components/Input/Input';
import Button from '../../components/Button/Button';

const GeneralSettings = () => {
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
    username: Yup.string().required(t('input_required') || ''),
    email: Yup.string().required(t('input_required') || ''),
    displayName: Yup.string().required(t('input_required') || ''),
  });

  const userSettingsFormik = useFormik({
    initialValues: {
      username: '',
      email: '',
      displayName: '',
    },
    validationSchema: userSettingsValidationSchema,
    onSubmit: async ({ username, email, displayName }) => {
      await update({ username, email, displayName });
    },
  });

  useFormikErrorFocus(
    userSettingsFormik,
    usernameRef,
    emailRef,
    displayNameRef
  );

  useEffect(() => {
    if (user) {
      userSettingsFormik.setValues({
        username: user.username,
        email: user.email,
        displayName: user.displayName,
      });
    }
  }, [user]);

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
            <Button type="submit" label={t('save')} />
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
          />
          <Dropdown
            label={t('language')}
            value={i18n.language}
            options={languageOptions}
            onChange={handleChangeLanguage}
            fullWidth
          />
          <div>
            <Button
              label={t('logout')}
              style="red"
              onClick={() => navigate('/logout')}
            />
          </div>
        </div>
      </div>
    </>
  );
};

export default GeneralSettings;