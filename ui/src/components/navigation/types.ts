import { ReactElement } from 'react';

export type INavigationItem = {
  to: string;
  title: string;
  requiredRoles?: string[];
  icon?: ReactElement;
  items?: INavigationItem[];
};
