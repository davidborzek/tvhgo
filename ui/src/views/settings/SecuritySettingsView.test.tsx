import { Session, Token, TwoFactorAuthSettings } from '@/clients/api/api.types';
import { afterEach, beforeEach, describe, expect, test, vi } from 'vitest';
import { cleanup, render } from '@testing-library/react';
import {
  useLoaderData,
  useLocation,
  useNavigate,
  useRevalidator,
} from 'react-router-dom';

import { SecuritySettingsRefreshStates } from './states';
import { Component as SecuritySettingsView } from './SecuritySettingsView';
import { TestIds } from '@/__test__/ids';
import { useManageSessions } from '@/hooks/session';
import { useManageTokens } from '@/hooks/token';
import { useUpdateUserPassword } from '@/hooks/user';
import { userEvent } from '@testing-library/user-event';

vi.mock('react-router-dom');

vi.mock('@/hooks/session');
vi.mock('@/hooks/2fa');
vi.mock('@/hooks/token');
vi.mock('@/hooks/user');

const sessions: Session[] = [
  {
    clientIp: '10.0.0.1',
    createdAt: 0,
    id: 1,
    lastUsedAt: 0,
    userAgent:
      'Mozilla/5.0 (X11; Linux x86_64; rv:120.0) Gecko/20100101 Firefox/120.0',
    userId: 1,
  },
  {
    clientIp: '10.0.0.2',
    createdAt: 0,
    id: 2,
    lastUsedAt: 0,
    userAgent:
      'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36',
    userId: 1,
  },
];

const tokens: Token[] = [
  {
    createdAt: 0,
    id: 1,
    name: 'My Token 1',
    updatedAt: 0,
  },
  {
    createdAt: 0,
    id: 2,
    name: 'My Token 2',
    updatedAt: 0,
  },
];

const twoFactorAuthEnabled: TwoFactorAuthSettings = {
  enabled: true,
};

const twoFactorAuthDisabled: TwoFactorAuthSettings = {
  enabled: false,
};

afterEach(() => {
  vi.restoreAllMocks();
  cleanup();
});

const navigateMock = vi.fn();
const updatePasswordMock = vi.fn();
const revokeSessionMock = vi.fn();
const revokeTokenMock = vi.fn();
const revalidateMock = vi.fn();

beforeEach(() => {
  vi.mocked(useLocation).mockReturnValue({
    hash: '',
    key: '',
    pathname: '',
    search: '',
    state: null,
  });

  vi.mocked(useNavigate).mockReturnValue(navigateMock);
  vi.mocked(useRevalidator).mockReturnValue({
    revalidate: revalidateMock,
    state: 'idle',
  });

  vi.mocked(useUpdateUserPassword).mockReturnValue({
    updatePassword: updatePasswordMock,
  });
  updatePasswordMock.mockResolvedValue(true);

  vi.mocked(useManageSessions).mockReturnValue({
    revokeSession: revokeSessionMock,
    revokeUserSession: vi.fn(),
  });

  vi.mocked(useManageTokens).mockReturnValue({
    revokeToken: revokeTokenMock,
  });

  vi.mocked(useLoaderData).mockReturnValue([
    twoFactorAuthEnabled,
    sessions,
    tokens,
  ]);
});

test('should render', () => {
  const document = render(<SecuritySettingsView />);
  expect(document.asFragment()).toMatchSnapshot();
});

describe('update password', () => {
  test('should update password', async () => {
    const document = render(<SecuritySettingsView />);

    const currentPassword = '12345678';
    const password = '87654321';

    const currenPasswordInput = document.container.querySelector(
      'input[name=currentPassword]'
    ) as Element;

    const passwordInput = document.container.querySelector(
      'input[name=password]'
    ) as Element;

    const passwordRepeatInput = document.container.querySelector(
      'input[name=passwordRepeat]'
    ) as Element;

    const saveButton = document.getByText('save');

    await userEvent.type(currenPasswordInput, currentPassword);
    await userEvent.type(passwordInput, password);
    await userEvent.type(passwordRepeatInput, password);

    expect(currenPasswordInput).toHaveValue(currentPassword);
    expect(passwordInput).toHaveValue(password);
    expect(passwordRepeatInput).toHaveValue(password);

    await userEvent.click(saveButton);

    expect(updatePasswordMock).toBeCalledWith({ currentPassword, password });

    expect(currenPasswordInput).toHaveValue('');
    expect(passwordInput).toHaveValue('');
    expect(passwordRepeatInput).toHaveValue('');
  });

  test('should not update password for all empty inputs', async () => {
    const document = render(<SecuritySettingsView />);

    const saveButton = document.getByText('save');
    await userEvent.click(saveButton);

    expect(updatePasswordMock).not.toBeCalled();

    const inputRequiredError = document.getAllByText('input_required');
    expect(inputRequiredError).toHaveLength(2);

    const passwordDoNotMatch = document.getByText('passwords_do_not_match');
    expect(passwordDoNotMatch).toBeInTheDocument();
  });

  test('should not update password for some empty inputs', async () => {
    const document = render(<SecuritySettingsView />);

    const password = '87654321';

    const passwordInput = document.container.querySelector(
      'input[name=password]'
    ) as Element;

    const saveButton = document.getByText('save');

    await userEvent.type(passwordInput, password);
    expect(passwordInput).toHaveValue(password);

    await userEvent.click(saveButton);
    expect(passwordInput).toHaveValue(password);

    expect(updatePasswordMock).not.toBeCalled();

    const inputRequiredError = document.getByText('input_required');
    expect(inputRequiredError).toBeInTheDocument();

    const passwordDoNotMatch = document.getByText('passwords_do_not_match');
    expect(passwordDoNotMatch).toBeInTheDocument();
  });

  test('should not update password when passwords do not match', async () => {
    const document = render(<SecuritySettingsView />);

    const currentPassword = '12345678';
    const password = '87654321';

    const currenPasswordInput = document.container.querySelector(
      'input[name=currentPassword]'
    ) as Element;

    const passwordInput = document.container.querySelector(
      'input[name=password]'
    ) as Element;

    const passwordRepeatInput = document.container.querySelector(
      'input[name=passwordRepeat]'
    ) as Element;
    const saveButton = document.getByText('save');

    await userEvent.type(currenPasswordInput, currentPassword);
    await userEvent.type(passwordInput, password);
    await userEvent.type(passwordRepeatInput, 'invalid');

    await userEvent.click(saveButton);
    expect(passwordInput).toHaveValue(password);

    expect(updatePasswordMock).not.toBeCalled();

    const inputRequiredError = document.queryByText('input_required');
    expect(inputRequiredError).not.toBeInTheDocument();

    const passwordDoNotMatch = document.getByText('passwords_do_not_match');
    expect(passwordDoNotMatch).toBeInTheDocument();
  });

  test('should not update password when password is invalid', async () => {
    const document = render(<SecuritySettingsView />);

    const currentPassword = '12345678';
    const password = '1234';

    const currenPasswordInput = document.container.querySelector(
      'input[name=currentPassword]'
    ) as Element;

    const passwordInput = document.container.querySelector(
      'input[name=password]'
    ) as Element;

    const passwordRepeatInput = document.container.querySelector(
      'input[name=passwordRepeat]'
    ) as Element;
    const saveButton = document.getByText('save');

    await userEvent.type(currenPasswordInput, currentPassword);
    await userEvent.type(passwordInput, password);
    await userEvent.type(passwordRepeatInput, password);

    await userEvent.click(saveButton);
    expect(passwordInput).toHaveValue(password);

    expect(updatePasswordMock).not.toBeCalled();

    const inputRequiredError = document.queryByText('input_required');
    expect(inputRequiredError).not.toBeInTheDocument();

    const passwordDoNotMatch = document.queryByText('passwords_do_not_match');
    expect(passwordDoNotMatch).not.toBeInTheDocument();

    const passwordInvalid = document.queryByText('password_min_chars_error');
    expect(passwordInvalid).toBeInTheDocument();
  });
});

describe('twofa settings', () => {
  test('should open disable twofa modal', async () => {
    const document = render(<SecuritySettingsView />);

    const disableButton = document.getByTestId(TestIds.TWOFA_DISABLE_BUTTON);

    await userEvent.click(disableButton);

    expect(navigateMock).toHaveBeenCalledWith(
      '/settings/security/two-factor-auth/disable'
    );
  });

  test('should open setup twofa modal', async () => {
    vi.mocked(useLoaderData).mockReturnValue([
      twoFactorAuthDisabled,
      sessions,
      tokens,
    ]);

    const document = render(<SecuritySettingsView />);

    const enableButton = document.getByTestId(TestIds.TWOFA_ENABLE_BUTTON);

    await userEvent.click(enableButton);

    expect(navigateMock).toHaveBeenCalledWith(
      '/settings/security/two-factor-auth/setup'
    );
  });
});

test.each(sessions)('should revoke session with id $id', async (session) => {
  const document = render(<SecuritySettingsView />);

  const revokeSessionButtons = document.getAllByTestId(
    TestIds.REVOKE_SESSION_BUTTON
  );
  expect(revokeSessionButtons).toHaveLength(sessions.length);

  await userEvent.click(revokeSessionButtons[sessions.indexOf(session)]);

  expect(revokeSessionMock).toBeCalledWith(session.id);
});

test.each(tokens)('should revoke token with id $id', async (token) => {
  const document = render(<SecuritySettingsView />);

  const revokeTokenButtons = document.getAllByTestId(
    TestIds.REVOKE_TOKEN_BUTTON
  );
  expect(revokeTokenButtons).toHaveLength(tokens.length);

  await userEvent.click(revokeTokenButtons[tokens.indexOf(token)]);

  expect(revokeTokenMock).toBeCalledWith(token.id);
});

test('should refresh twofa settings on state change', () => {
  vi.mocked(useLocation).mockReturnValue({
    hash: '',
    key: '',
    pathname: '',
    search: '',
    state: SecuritySettingsRefreshStates.TWOFA,
  });

  render(<SecuritySettingsView />);

  expect(revalidateMock).toHaveBeenCalled();
});

test('should refresh tokens on state change', () => {
  vi.mocked(useLocation).mockReturnValue({
    hash: '',
    key: '',
    pathname: '',
    search: '',
    state: SecuritySettingsRefreshStates.TOKEN,
  });

  render(<SecuritySettingsView />);

  expect(revalidateMock).toHaveBeenCalled();
});
