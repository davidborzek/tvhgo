import { cleanup, render } from '@testing-library/react';
import { afterEach, expect, test, vi } from 'vitest';
import SettingsView from './SettingsView';
import { BrowserRouter } from 'react-router-dom';

afterEach(() => {
  vi.restoreAllMocks();
  cleanup();
});

test('should render', () => {
  const document = render(<SettingsView />, { wrapper: BrowserRouter });
  expect(document.asFragment()).toMatchSnapshot();
});
