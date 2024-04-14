import {
  PropsWithChildren,
  ReactElement,
  useEffect,
  useRef,
  useState,
} from 'react';
import LoadingBar, { LoadingBarRef } from 'react-top-loading-bar';

import { LoadingContext } from '@/contexts/LoadingContext';
import { Outlet, useNavigation } from 'react-router-dom';

export default function LoadingProvider(): ReactElement {
  const { state } = useNavigation();
  const [isLoading, setIsLoading] = useState(false);

  const ref = useRef<LoadingBarRef>(null);

  useEffect(() => {
    isLoading ? ref.current?.continuousStart() : ref.current?.complete();
  }, [isLoading]);

  useEffect(() => {
    state === 'loading'
      ? ref.current?.continuousStart()
      : ref.current?.complete();
  }, [state]);

  return (
    <LoadingContext.Provider
      value={{
        isLoading,
        setIsLoading,
      }}
    >
      <LoadingBar ref={ref} color="#00FFFF" />
      <Outlet />
    </LoadingContext.Provider>
  );
}
