import { cleanup, render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { afterEach, expect, vi, test, describe } from 'vitest';
import LoginView from './LoginView';
import useLogin from '@/hooks/login';
import userEvent from '@testing-library/user-event';

vi.mock('@/hooks/login');

afterEach(() => {
  vi.resetAllMocks();
  cleanup();
});

describe('login', () => {
  test('should render', () => {
    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: () => {},
      twoFactorRequired: false,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    expect(document.asFragment()).toMatchSnapshot();
  });

  test('should render loading button', () => {
    vi.mocked(useLogin).mockReturnValue({
      loading: true,
      login: () => {},
      twoFactorRequired: false,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    expect(document.getByText('login_pending')).toBeInTheDocument();
  });

  test('should show errors when no input is present', async () => {
    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: () => {},
      twoFactorRequired: false,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    await userEvent.click(document.getByText('login'));
    await document.findAllByText('username_required');

    expect(document.getByText('username_required')).toBeInTheDocument();
    expect(document.getByText('password_required')).toBeInTheDocument();
  });

  test('should show username_required when username is missing', async () => {
    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: () => {},
      twoFactorRequired: false,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    const passwordInput = document.container.querySelector(
      'input[name="password"]'
    ) as Element;

    await userEvent.type(passwordInput, 'password');

    await userEvent.click(document.getByText('login'));
    await document.findAllByText('username_required');

    expect(document.getByText('username_required')).toBeInTheDocument();
    expect(document.queryByText('password_required')).not.toBeInTheDocument();
  });

  test('should show password_required when password is missing', async () => {
    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: () => {},
      twoFactorRequired: false,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    const usernameInput = document.container.querySelector(
      'input[name="username"]'
    ) as Element;

    await userEvent.type(usernameInput, 'username');

    await userEvent.click(document.getByText('login'));
    await document.findAllByText('password_required');

    expect(document.getByText('password_required')).toBeInTheDocument();
    expect(document.queryByText('username_required')).not.toBeInTheDocument();
  });

  test('should login', async () => {
    const loginMock = vi.fn();
    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: loginMock,
      twoFactorRequired: false,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    const usernameInput = document.container.querySelector(
      'input[name="username"]'
    ) as Element;
    const passwordInput = document.container.querySelector(
      'input[name="password"]'
    ) as Element;

    const username = 'myUsername';
    const password = 'myPassword';

    await userEvent.type(usernameInput, username);
    await userEvent.type(passwordInput, password);

    expect(usernameInput).toHaveValue(username);
    expect(passwordInput).toHaveValue(password);

    await userEvent.click(document.getByText('login'));

    expect(loginMock).toHaveBeenNthCalledWith(1, username, password);
  });
});

describe('login with totp', () => {
  test('should render two factor auth', () => {
    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: () => {},
      twoFactorRequired: true,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    expect(document.asFragment()).toMatchSnapshot();
  });

  test('should render loading button', () => {
    vi.mocked(useLogin).mockReturnValue({
      loading: true,
      login: () => {},
      twoFactorRequired: true,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    expect(document.getByText('login_pending')).toBeInTheDocument();
  });

  test('should show two_factor_code_required when code is missing', async () => {
    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: () => {},
      twoFactorRequired: true,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    await userEvent.click(document.getByText('login'));
    await document.findAllByText('two_factor_code_required');

    expect(document.getByText('two_factor_code_required')).toBeInTheDocument();
  });

  test('should login with totp', async () => {
    const loginMock = vi.fn();
    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: loginMock,
      twoFactorRequired: false,
    });

    const document = render(<LoginView />, { wrapper: MemoryRouter });

    const usernameInput = document.container.querySelector(
      'input[name="username"]'
    ) as Element;
    const passwordInput = document.container.querySelector(
      'input[name="password"]'
    ) as Element;

    const username = 'myUsername';
    const password = 'myPassword';
    const code = '123456';

    await userEvent.type(usernameInput, username);
    await userEvent.type(passwordInput, password);

    expect(usernameInput).toHaveValue(username);
    expect(passwordInput).toHaveValue(password);

    vi.mocked(useLogin).mockReturnValue({
      loading: false,
      login: loginMock,
      twoFactorRequired: true,
    });

    document.rerender(<LoginView />);

    const codeInput = document.container.querySelector(
      'input[name="code"]'
    ) as Element;

    await userEvent.type(codeInput, code);

    await userEvent.click(document.getByText('login'));

    expect(loginMock).toHaveBeenNthCalledWith(1, username, password, code);
  });
});
