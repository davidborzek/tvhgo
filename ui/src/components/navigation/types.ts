import { ReactElement } from 'react';

export type INavigationItem = {
  to: string;
  title: string;
  icon?: ReactElement;
  items?: INavigationItem[];
};
