import { createContext, useContext } from "react";

export type AuthContextProps = {
  username: string | null;
  setUser: (username: string | null) => void;
};

export const AuthContext = createContext<AuthContextProps>({
  username: null,
  setUser: () => {
    throw new Error("not implemented");
  },
});

export const useAuth = () => {
  return useContext(AuthContext);
};

