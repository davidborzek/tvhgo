import { useState } from "react";
import { useAuth } from "../contexts/AuthContext";
import { useTranslation } from "react-i18next";
import { getUser, login, ApiError } from "../clients/api/api";

type LoginFunc = (username: string, password: string) => void;

const useLogin = () => {
  const { t } = useTranslation("errors");
  const { setUser } = useAuth();
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const fetchUser = () => {
    setLoading(true);
    getUser()
      .then(({ username }) => {
        setUser(username);
      })
      .catch(() => {
        setError(t("unexpected"));
      })
      .finally(() => setLoading(false));
  };

  const _login: LoginFunc = (username, password) => {
    setLoading(true);
    login(username, password)
      .catch((error) => {
        if (error instanceof ApiError && error.code == 401) {
          setError(t("invalid_login"));
        } else {
          setError(t("unexpected"));
        }
      })
      .then(fetchUser)
      .finally(() => setLoading(false));
  };

  return { login: _login, loading, error };
};

export default useLogin;
