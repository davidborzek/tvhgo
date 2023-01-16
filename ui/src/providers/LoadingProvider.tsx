import {
  PropsWithChildren,
  ReactElement,
  useEffect,
  useRef,
  useState,
} from 'react';
import LoadingBar, { LoadingBarRef } from 'react-top-loading-bar';
import { LoadingContext } from '../contexts/LoadingContext';

export default function LoadingProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const [isLoading, setIsLoading] = useState(true);

  const ref = useRef<LoadingBarRef>(null);

  useEffect(() => {
    console.log(isLoading);

    isLoading ? ref.current?.continuousStart() : ref.current?.complete();
  }, [isLoading]);

  return (
    <LoadingContext.Provider
      value={{
        isLoading,
        setIsLoading,
      }}
    >
      <LoadingBar ref={ref} color="#00FFFF" />
      {children}
    </LoadingContext.Provider>
  );
}
