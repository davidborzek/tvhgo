import { ComponentPropsWithoutRef } from 'react';

export default function TableBody({
  children,
  ...props
}: ComponentPropsWithoutRef<'tbody'>) {
  return <tbody {...props}>{children}</tbody>;
}
