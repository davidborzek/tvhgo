import { createContext, useContext } from 'react';

export type LoadingContextProps = {
  isLoading: boolean;
  setIsLoading: (isLoading: boolean) => void;
};

export const LoadingContext = createContext<LoadingContextProps>({
  isLoading: false,
  setIsLoading: () => {
    throw new Error('not implemented');
  },
});

export const useLoading = () => {
  return useContext(LoadingContext);
};
