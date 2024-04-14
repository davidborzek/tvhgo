import { PropsWithChildren, ReactElement, useEffect, useState } from 'react';
import { Theme, ThemeContext } from '@/contexts/ThemeContext';

const LOCAL_STORAGE_KEY = 'tvhgo_theme';

const saved = localStorage.getItem(LOCAL_STORAGE_KEY);

export function ThemeProvider({
  children,
}: PropsWithChildren<unknown>): ReactElement {
  const [theme, setTheme] = useState<Theme>(Theme.DARK);

  useEffect(() => {
    if (!Object.values(Theme).includes(saved as Theme)) {
      return;
    }

    setTheme(saved as Theme);
  }, []);

  useEffect(() => {
    document.body.setAttribute('data-theme', theme);
  }, [theme]);

  function setThemePersistent(theme: Theme) {
    setTheme(theme);
    localStorage.setItem(LOCAL_STORAGE_KEY, theme);
  }

  return (
    <ThemeContext.Provider
      value={{
        theme,
        setTheme: setThemePersistent,
      }}
    >
      {children}
    </ThemeContext.Provider>
  );
}
