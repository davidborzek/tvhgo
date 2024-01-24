import { UserResponse } from '@/clients/api/api.types';
import { useAuth } from '@/contexts/AuthContext';
import { Theme, useTheme } from '@/contexts/ThemeContext';
import { useUpdateUser } from '@/hooks/user';
import { cleanup, render } from '@testing-library/react';
import { useNavigate } from 'react-router-dom';
import { afterEach, beforeEach, expect, test, vi } from 'vitest';
import GeneralSettingsView from './GeneralSettingsView';
import { userEvent } from '@testing-library/user-event';
import i18n from 'i18next';

vi.mock('@/contexts/AuthContext');
vi.mock('@/contexts/ThemeContext');
vi.mock('@/hooks/user');
vi.mock('react-router-dom');
vi.mock('i18next');

const user: UserResponse = {
  id: 1,
  username: 'someUsername',
  displayName: 'Some User',
  email: 'some@email.com',
  createdAt: 0,
  updatedAt: 0,
};

afterEach(() => {
  vi.restoreAllMocks();
  cleanup();
});

const setThemeMock = vi.fn();
const navigateMock = vi.fn();
const updateUserMock = vi.fn();

beforeEach(() => {
  vi.mocked(useAuth).mockReturnValue({
    setUser: vi.fn(),
    user,
  });

  vi.mocked(useTheme).mockReturnValue({
    theme: Theme.DARK,
    setTheme: setThemeMock,
  });

  vi.mocked(useUpdateUser).mockReturnValue({
    update: updateUserMock,
  });

  vi.mocked(useNavigate).mockReturnValue(navigateMock);
});

test('should render', () => {
  const document = render(<GeneralSettingsView />);
  expect(document.asFragment()).toMatchSnapshot();
});

test('should update user', async () => {
  const document = render(<GeneralSettingsView />);

  const newUsername = 'newUsername';
  const newEmail = 'new@email.com';
  const newDisplayName = 'New Name';

  const usernameInput = document.container.querySelector(
    'input[name=username]'
  ) as Element;
  const emailInput = document.container.querySelector(
    'input[name=email]'
  ) as Element;
  const displayNameInput = document.container.querySelector(
    'input[name=displayName]'
  ) as Element;
  const saveButton = document.getByTestId('save_user');

  await userEvent.clear(usernameInput);
  await userEvent.type(usernameInput, newUsername);
  await userEvent.clear(emailInput);
  await userEvent.type(emailInput, newEmail);
  await userEvent.clear(displayNameInput);
  await userEvent.type(displayNameInput, newDisplayName);

  expect(usernameInput).toHaveValue(newUsername);
  expect(emailInput).toHaveValue(newEmail);
  expect(displayNameInput).toHaveValue(newDisplayName);

  await userEvent.click(saveButton);

  expect(updateUserMock).toHaveBeenCalledWith({
    username: newUsername,
    email: newEmail,
    displayName: newDisplayName,
  });
});

test('should logout', async () => {
  const document = render(<GeneralSettingsView />);

  const logoutButton = document.getByTestId('logout_button');
  await userEvent.click(logoutButton);

  expect(navigateMock).toHaveBeenCalledWith('/logout');
});

test('should change theme', async () => {
  const document = render(<GeneralSettingsView />);

  const themeDropdown = document.getByTestId('theme_dropdown');
  await userEvent.selectOptions(themeDropdown, Theme.LIGHT);

  expect(setThemeMock).toHaveBeenCalledWith(Theme.LIGHT);
});

test('should change language', async () => {
  const changeLanguageMock = vi.spyOn(i18n, 'changeLanguage');

  const document = render(<GeneralSettingsView />);

  const languageDropdown = document.getByTestId('language_dropdown');
  await userEvent.selectOptions(languageDropdown, 'de');

  expect(changeLanguageMock).toHaveBeenCalledWith('de', expect.anything());
});
