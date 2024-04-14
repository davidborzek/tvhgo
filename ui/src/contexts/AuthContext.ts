import { createContext, useContext } from 'react';

import { UserResponse } from '@/clients/api/api.types';

export type AuthContextProps = {
  user: UserResponse | null;
  setUser: (user: UserResponse | null) => void;
};

export const AuthContext = createContext<AuthContextProps>({
  user: null,
  setUser: () => {
    throw new Error('not implemented');
  },
});

export const useAuth = () => {
  return useContext(AuthContext);
};
