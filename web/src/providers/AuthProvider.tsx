import { PropsWithChildren, ReactElement, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { toast } from "react-toastify";
import { getUser, ApiError } from "../clients/api/api";
import { AuthContext } from "../contexts/AuthContext";

const NOTIFICATION_ID = "authError";

const notify = (message?: string | null) => {
  toast.error(message, {
    toastId: NOTIFICATION_ID,
    updateId: NOTIFICATION_ID,
  });
};

export default function AuthProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const { t } = useTranslation("errors");

  const [username, setUsername] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    getUser()
      .then((response) => {
        setUsername(response.username);
      })
      .catch((error) => {
        if (error instanceof ApiError && error.code === 401) {
          setUsername(null);
        } else {
          notify(t("unexpected"));
        }
      })
      .finally(() => setIsLoading(false));
  }, []);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <AuthContext.Provider
      value={{
        username,
        setUser: setUsername,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
