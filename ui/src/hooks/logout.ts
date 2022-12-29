import { useState } from "react";
import { useTranslation } from "react-i18next";
import { logout } from "../clients/api/api";
import { useAuth } from "../contexts/AuthContext";

const useLogout = () => {
  const { t } = useTranslation("errors");
  const authContext = useAuth();
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const _logout = () => {
    setLoading(true);
    logout()
      .catch(() => {
        setError(t("unexpected"));
      })
      .then(() => authContext.setUser(null))
      .finally(() => setLoading(false));
  };

  return { logout: _logout, loading, error };
};

export default useLogout;
