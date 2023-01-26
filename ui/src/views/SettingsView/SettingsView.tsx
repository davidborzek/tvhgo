import { useFormik } from 'formik';
import { useTranslation } from 'react-i18next';
import { useNavigate } from 'react-router-dom';
import Button from '../../components/Button/Button';
import Dropdown, { Option } from '../../components/Dropdown/Dropdown';
import Error from '../../components/Error/Error';
import Form from '../../components/Form/Form';
import FormGroup from '../../components/Form/FormGroup/FormGroup';
import Input from '../../components/Input/Input';
import { Theme, useTheme } from '../../contexts/ThemeContext';
import formik from '../../hooks/formik';
import { useFetchUser } from '../../hooks/user';
import i18n from '../../i18n/i18n';
import styles from './SettingsView.module.scss';
import * as Yup from 'yup';

function SettingsView() {
  const navigate = useNavigate();
  const { t } = useTranslation();

  const { setTheme, theme } = useTheme();

  const { update, user, error } = useFetchUser();

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

  const passwordChangeValidationSchema = Yup.object().shape({
    password: Yup.string()
      .required(t('input_required') || '')
      .min(8, t('password_min_chars_error') || ''),
    passwordRepeat: Yup.string()
      .required(t('passwords_do_not_match') || '')
      .oneOf([Yup.ref('password'), null], t('passwords_do_not_match') || ''),
  });

  const passwordChangeFormik = useFormik({
    initialValues: {
      password: '',
      passwordRepeat: '',
    },
    validationSchema: passwordChangeValidationSchema,
    onSubmit: async ({ password }) => {
      await update({ password });

      passwordChangeFormik.resetForm();
    },
  });

  if (error) {
    return <Error message={t('unexpected')} />;
  }

  return (
    <div className={styles.Settings}>
      <div className={styles.heading}>
        <h1>{t('settings')}</h1>
      </div>
      <div className={styles.row}>
        <Form
          onSubmit={passwordChangeFormik.handleSubmit}
          className={styles.section}
        >
          <Input
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
            fullWidth
          />
          <Input
            label={t('password_repeat')}
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
            fullWidth
          />
          <div>
            <Button type="submit" label={t('save')} />
          </div>
        </Form>
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
    </div>
  );
}

export default SettingsView;
