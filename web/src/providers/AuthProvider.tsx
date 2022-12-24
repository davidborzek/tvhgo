import { PropsWithChildren, ReactElement, useEffect, useState } from "react";
import { getUser, ApiError } from "../clients/api/api";
import {AuthContext} from "../contexts/AuthContext";

export default function AuthProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const [username, setUsername] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    getUser()
      .then((response) => {
        setUsername(response.username);
      })
      .catch((error) => {
        if (error instanceof ApiError && error.code === 401) {
          setUsername(null);
        } else {
          setError(error.message);
        }
      })
      .finally(() => setIsLoading(false));
  }, []);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Unexpected error: {error}</div>;
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
