import { ApiError, updateUser } from '@/clients/api/api';
import { afterEach, beforeEach, describe, expect, it, test, vi } from 'vitest';

import { renderHook } from '@testing-library/react';
import { useAuth } from '@/contexts/AuthContext';
import { useNotification } from './notification';
import { useUpdateUser } from './user';

const updateUserOpts = {
  email: 'someEmail',
  displayName: 'someName',
  username: 'someUsername',
};

const user = {
  id: 1,
  username: 'someUsername',
  email: 'someEmail',
  displayName: 'someName',
  createdAt: 0,
  updatedAt: 0,
  twoFactor: false,
};

vi.mock('@/clients/api/api', async () => {
  const actual = await vi.importActual('@/clients/api/api');
  return {
    ...actual,
    updateUser: vi.fn(),
  };
});
vi.mock('@/contexts/AuthContext');
vi.mock('@/hooks/notification');

describe('useUpdateUser', () => {
  const setUserMock = vi.fn();
  const notifyErrorMock = vi.fn();
  const notifySuccessMock = vi.fn();
  const dismissNotificationMock = vi.fn();

  beforeEach(() => {
    vi.mocked(useAuth).mockReturnValue({ setUser: setUserMock, user: null });

    vi.mocked(useNotification).mockReturnValue({
      dismissNotification: dismissNotificationMock,
      notifyError: notifyErrorMock,
      notifySuccess: notifySuccessMock,
    });
  });

  afterEach(() => {
    vi.resetAllMocks();
  });

  it('should update a user', async () => {
    vi.mocked(updateUser).mockResolvedValue(user);

    const { result } = renderHook(() => useUpdateUser());

    await result.current.update(updateUserOpts);

    expect(useNotification).toHaveBeenNthCalledWith(1, 'manageUser');

    expect(dismissNotificationMock).toHaveBeenCalledOnce();

    expect(updateUser).toHaveBeenCalledOnce();
    expect(updateUser).toHaveBeenCalledWith(updateUserOpts);

    expect(notifySuccessMock).toHaveBeenCalledOnce();
    expect(notifySuccessMock).toHaveBeenCalledWith('user_updated_successfully');

    expect(setUserMock).toHaveBeenCalledOnce();
    expect(setUserMock).toHaveBeenCalledWith(user);

    expect(notifyErrorMock).toHaveBeenCalledTimes(0);
  });

  test.each([
    [500, 'Internal Server Error', 'unexpected'],
    [400, 'username already exists', 'username_already_exists'],
    [400, 'email already exists', 'email_already_exists'],
  ])(
    'should fail with %i, %s -> %s',
    async (statusCode, message, expectedNotification) => {
      vi.mocked(updateUser).mockRejectedValue(
        new ApiError(statusCode, message)
      );

      const { result } = renderHook(() => useUpdateUser());

      await result.current.update(updateUserOpts);

      expect(useNotification).toHaveBeenNthCalledWith(1, 'manageUser');

      expect(dismissNotificationMock).toHaveBeenCalledOnce();

      expect(updateUser).toHaveBeenCalledOnce();
      expect(updateUser).toHaveBeenCalledWith(updateUserOpts);

      expect(notifyErrorMock).toHaveBeenCalledOnce();
      expect(notifyErrorMock).toHaveBeenCalledWith(expectedNotification);

      expect(setUserMock).toHaveBeenCalledTimes(0);
      expect(notifySuccessMock).toHaveBeenCalledTimes(0);
    }
  );
});
