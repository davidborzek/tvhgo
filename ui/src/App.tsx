import { Suspense, lazy, useEffect } from 'react';
import {
  BrowserRouter,
  Routes,
  Route,
  Navigate,
  Outlet,
} from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import { useAuth } from '@/contexts/AuthContext';
import useLogout from '@/hooks/logout';
import AuthProvider from '@/providers/AuthProvider';
import { ThemeProvider } from '@/providers/ThemeProvider';
import { useTheme } from '@/contexts/ThemeContext';

import 'react-toastify/dist/ReactToastify.css';
import LoadingProvider from '@/providers/LoadingProvider';
import EmptyState from '@/components/common/emptyState/EmptyState';
import ButtonLink from '@/components/common/button/ButtonLink';
import { useTranslation } from 'react-i18next';

const LoginView = lazy(() => import('@/views/login/LoginView'));
const DashboardView = lazy(() => import('@/views/dashboard/DashboardView'));

const ChannelListView = lazy(
  () => import('@/views/channels/list/ChannelListView')
);
const ChannelView = lazy(() => import('@/views/channels/detail/ChannelView'));

const GuideView = lazy(() => import('@/views/epg/guide/GuideView'));
const EventView = lazy(() => import('@/views/epg/event/EventView'));

const RecordingsView = lazy(
  () => import('@/views/recordings/RecordingsView/RecordingsView')
);
const RecordingDetailView = lazy(
  () => import('@/views/recordings/RecordingDetailView/RecordingDetailView')
);
const RecordingCreateView = lazy(
  () => import('@/views/recordings/create/RecordingCreateView')
);
const ChannelSelectModal = lazy(
  () => import('@/components/channels/selectModal/ChannelSelectModal')
);

const SettingsView = lazy(() => import('@/views/settings/SettingsView'));
const GeneralSettingsView = lazy(
  () => import('@/views/settings/GeneralSettingsView')
);
const SecuritySettingsView = lazy(
  () => import('@/views/settings/SecuritySettingsView')
);

const TwoFactorAuthDisableModal = lazy(
  () => import('@/modals/twoFactorAuth/disable/TwoFactorAuthDisableModal')
);
const TwoFactorAuthSetupModal = lazy(
  () => import('@/modals/twoFactorAuth/setup/TwoFactorAuthSetupModal')
);
const CreateTokenModal = lazy(
  () => import('@/modals/token/create/CreateTokenModal')
);

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
  }, []);

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
      <ButtonLink label={t('go_back')} href="/channels" />
    </EmptyState>
  );
}

function App() {
  return (
    <ThemeProvider>
      <LoadingProvider>
        <AuthProvider>
          <BrowserRouter>
            <Routes>
              <Route element={<Unauthenticated />}>
                <Route
                  path="/login"
                  element={
                    <Suspense>
                      <LoginView />
                    </Suspense>
                  }
                />
              </Route>

              <Route element={<Authenticated />}>
                <Route
                  element={
                    <Suspense>
                      <DashboardView />
                    </Suspense>
                  }
                >
                  <Route
                    path="/"
                    element={<Navigate to={'/channels'} replace />}
                  />
                  <Route
                    path="/channels"
                    element={
                      <Suspense>
                        <ChannelListView />
                      </Suspense>
                    }
                  />
                  <Route
                    path="/channels/:id"
                    element={
                      <Suspense>
                        <ChannelView />
                      </Suspense>
                    }
                  />
                  <Route
                    path="/guide"
                    element={
                      <Suspense>
                        <GuideView />
                      </Suspense>
                    }
                  />
                  <Route
                    path="/guide/events/:id"
                    element={
                      <Suspense>
                        <EventView />
                      </Suspense>
                    }
                  />
                  <Route
                    path="/recordings"
                    element={
                      <Suspense>
                        <RecordingsView />
                      </Suspense>
                    }
                  />
                  <Route
                    path="/recordings/:id"
                    element={
                      <Suspense>
                        <RecordingDetailView />
                      </Suspense>
                    }
                  />
                  <Route
                    path="/recordings/create"
                    element={
                      <Suspense>
                        <RecordingCreateView />
                      </Suspense>
                    }
                  >
                    <Route
                      path="select-channel"
                      element={
                        <Suspense>
                          <ChannelSelectModal />
                        </Suspense>
                      }
                    />
                  </Route>

                  <Route
                    path="/settings"
                    element={
                      <Suspense>
                        <SettingsView />
                      </Suspense>
                    }
                  >
                    <Route
                      path=""
                      element={<Navigate to={'general'} replace />}
                    />
                    <Route
                      path="general"
                      element={
                        <Suspense>
                          <GeneralSettingsView />
                        </Suspense>
                      }
                    />
                    <Route
                      path="security"
                      element={
                        <Suspense>
                          <SecuritySettingsView />
                        </Suspense>
                      }
                    >
                      <Route
                        path="two-factor-auth/disable"
                        element={
                          <Suspense>
                            <TwoFactorAuthDisableModal />
                          </Suspense>
                        }
                      />
                      <Route
                        path="two-factor-auth/setup"
                        element={
                          <Suspense>
                            <TwoFactorAuthSetupModal />
                          </Suspense>
                        }
                      />
                      <Route
                        path="tokens/create"
                        element={
                          <Suspense>
                            <CreateTokenModal />
                          </Suspense>
                        }
                      />
                    </Route>
                  </Route>
                  <Route path="*" element={<NotFound />} />
                </Route>

                <Route path="/logout" element={<Logout />} />
              </Route>
            </Routes>
          </BrowserRouter>
        </AuthProvider>
      </LoadingProvider>
      <Notification />
    </ThemeProvider>
  );
}

export default App;
