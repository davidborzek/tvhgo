import 'react-toastify/dist/ReactToastify.css';

import {
  Navigate,
  Outlet,
  Route,
  RouterProvider,
  createBrowserRouter,
  createRoutesFromElements,
  useRouteError,
} from 'react-router-dom';
import { useEffect, useMemo } from 'react';

import { ApiError } from './clients/api/api';
import AuthProvider from '@/providers/AuthProvider';
import ButtonLink from '@/components/common/button/ButtonLink';
import EmptyState from '@/components/common/emptyState/EmptyState';
import Error from './components/common/error/Error';
import LoadingProvider from '@/providers/LoadingProvider';
import { ThemeProvider } from '@/providers/ThemeProvider';
import { ToastContainer } from 'react-toastify';
import { useAuth } from '@/contexts/AuthContext';
import useLogout from '@/hooks/logout';
import { useTheme } from '@/contexts/ThemeContext';
import { useTranslation } from 'react-i18next';

type AuthenticationCheckerProps = {
  redirect?: string;
};

function Unauthenticated(props: AuthenticationCheckerProps) {
  const authContext = useAuth();
  const isAuthenticated = !!authContext.user;

  return isAuthenticated ? <Navigate to={props.redirect || '/'} /> : <Outlet />;
}

function Authenticated(props: AuthenticationCheckerProps) {
  const authContext = useAuth();
  const isAuthenticated = !!authContext.user;

  return isAuthenticated ? (
    <Outlet />
  ) : (
    <Navigate to={props.redirect || '/login'} />
  );
}

function Logout() {
  const { logout } = useLogout();

  useEffect(() => {
    logout();
  }, [logout]);

  return <Navigate to="/login" />;
}

function Notification() {
  const { theme } = useTheme();
  return (
    <ToastContainer
      position="top-right"
      autoClose={5000}
      limit={5}
      theme={theme}
      newestOnTop
    />
  );
}

function NotFound() {
  const { t } = useTranslation();

  return (
    <EmptyState title={t('page_not_found')} subtitle=" ">
      <ButtonLink quiet label={t('go_back')} href="/channels" />
    </EmptyState>
  );
}

function ErrorBoundary() {
  const { t } = useTranslation();
  const error = useRouteError();

  const [text, shouldLogout] = useMemo(() => {
    if (error instanceof ApiError) {
      switch (error.code) {
        case 401:
          return ['', true];
        case 403:
          return [t('forbidden'), false];
        case 404:
          return [t('not_found'), false];
      }
    }

    return [t('unexpected'), false];
  }, [error]);

  if (shouldLogout) {
    return <Navigate to="/logout" />;
  }

  return <Error message={text} />;
}

function App() {
  const router = createBrowserRouter(
    createRoutesFromElements(
      <Route element={<LoadingProvider />}>
        <Route element={<Unauthenticated />}>
          <Route path="/login" lazy={() => import('@/views/login/LoginView')} />
        </Route>

        <Route element={<Authenticated />}>
          <Route lazy={() => import('@/views/dashboard/DashboardView')}>
            <Route errorElement={<ErrorBoundary />}>
              <Route path="/" element={<Navigate to={'/channels'} replace />} />
              <Route
                path="/search"
                lazy={() => import('@/views/search/SearchView')}
              />
              <Route
                path="/channels"
                lazy={() => import('@/views/channels/list/ChannelListView')}
              />
              <Route
                path="/channels/:id"
                lazy={() => import('@/views/channels/detail/ChannelView')}
              />
              <Route
                path="/guide"
                lazy={() => import('@/views/epg/guide/GuideView')}
                shouldRevalidate={({ currentUrl, nextUrl }) => {
                  for (const [key, val] of currentUrl.searchParams) {
                    // We don't want to revalidate when only some query params changes.
                    if (key === 'offset' || key === 'search') {
                      continue;
                    }

                    if (val !== nextUrl.searchParams.get(key)) {
                      return true;
                    }
                  }

                  return false;
                }}
              />
              <Route
                path="/guide/events/:id"
                lazy={() => import('@/views/epg/event/EventView')}
              />

              <Route
                path="/dvr"
                element={<Navigate to={'/dvr/recordings'} replace />}
              />
              <Route
                path="/dvr/recordings"
                lazy={() =>
                  import('@/views/recordings/RecordingsView/RecordingsView')
                }
              />
              <Route
                path="/dvr/recordings/:id"
                lazy={() =>
                  import(
                    '@/views/recordings/RecordingDetailView/RecordingDetailView'
                  )
                }
              />
              <Route
                path="/dvr/config"
                lazy={() => import('@/views/dvr/config/list/DVRConfigListView')}
              />
              <Route
                path="/dvr/config/:id"
                lazy={() =>
                  import('@/views/dvr/config/detail/DVRConfigDetailView')
                }
              />

              <Route
                path="/settings"
                lazy={() => import('@/views/settings/SettingsView')}
              >
                <Route path="" element={<Navigate to={'general'} replace />} />
                <Route
                  path="general"
                  lazy={() => import('@/views/settings/GeneralSettingsView')}
                />
                <Route
                  path="security"
                  lazy={() => import('@/views/settings/SecuritySettingsView')}
                >
                  <Route
                    path="two-factor-auth/disable"
                    lazy={() =>
                      import(
                        '@/modals/twoFactorAuth/disable/TwoFactorAuthDisableModal'
                      )
                    }
                  />
                  <Route
                    path="two-factor-auth/setup"
                    lazy={() =>
                      import(
                        '@/modals/twoFactorAuth/setup/TwoFactorAuthSetupModal'
                      )
                    }
                  />
                  <Route
                    path="tokens/create"
                    lazy={() =>
                      import('@/modals/token/create/CreateTokenModal')
                    }
                  />
                </Route>
                <Route
                  path="users"
                  lazy={() =>
                    import('@/views/settings/users/list/UserListView')
                  }
                >
                  <Route
                    path="create"
                    lazy={() => import('@/modals/user/create/UserCreateModal')}
                  />
                </Route>
                <Route
                  path="users/:id"
                  lazy={() =>
                    import('@/views/settings/users/detail/UserDetailView')
                  }
                />
              </Route>
              <Route path="*" element={<NotFound />} />
            </Route>
          </Route>

          <Route path="/logout" element={<Logout />} />
        </Route>
      </Route>
    )
  );

  return (
    <ThemeProvider>
      <AuthProvider>
        <RouterProvider router={router} />
      </AuthProvider>
      <Notification />
    </ThemeProvider>
  );
}

export default App;
