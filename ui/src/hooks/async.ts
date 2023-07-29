import { useEffect } from 'react';

type PromiseFun = () => Promise<unknown>;

export const usePromiseAll = (values: PromiseFun[]) => {
  useEffect(() => {
    Promise.all(values.map((fn) => fn()));
  }, []);
};
