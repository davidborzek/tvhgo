import { Formik, useFormik } from "formik";
import useLogin from "../../hooks/login";
import * as Yup from "yup";

import styles from "./Login.module.scss";
import { useTranslation } from "react-i18next";

const loginSchema = Yup.object().shape({
  username: Yup.string().required("Required"),
  password: Yup.string().required("Required"),
});

export default function Login() {
  const { t } = useTranslation("login");
  const { login, error, loading } = useLogin();

  console.log(error);
  

  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },
    validationSchema: loginSchema,
    onSubmit: async (values) => {
      login(values.username, values.password);
    },
  });

  return (
    <div className={styles.Login}>
      <p>Login</p>
      <form title="Login" onSubmit={formik.handleSubmit}>
        <p>
          <input
            type="text"
            name="username"
            value={formik.values.username}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
          />
        </p>
        <p>
          <input
            type="password"
            name="password"
            value={formik.values.password}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
            className={styles.input}
          />
        </p>
        <p>
          <button disabled={formik.isSubmitting}>{t("login")}</button>
        </p>
      </form>
    </div>
  );
}
