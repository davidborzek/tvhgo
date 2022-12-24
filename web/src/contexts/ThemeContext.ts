import { createContext } from "react";

export enum Theme {
  LIGHT = "light",
  DARK = "dark",
}

type ThemeContextProps = {
  theme: Theme;
  setTheme: (theme: Theme) => void;
};

export const ThemeContext = createContext<ThemeContextProps>({
  theme: Theme.LIGHT,
  setTheme: () => {
    throw new Error("not implemented");
  },
});
