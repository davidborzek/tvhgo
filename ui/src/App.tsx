import { useEffect } from 'react';
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
import LoginView from '@/views/login/LoginView';
import { useTheme } from '@/contexts/ThemeContext';

import 'react-toastify/dist/ReactToastify.css';
import DashboardView from '@/views/dashboard/DashboardView';
import GuideView from '@/views/epg/guide/GuideView';
import EventView from '@/views/epg/event/EventView';
import RecordingsView from '@/views/recordings/RecordingsView/RecordingsView';
import RecordingDetailView from '@/views/recordings/RecordingDetailView/RecordingDetailView';
import LoadingProvider from '@/providers/LoadingProvider';
import SettingsView from '@/views/settings/SettingsView';
import GeneralSettingsView from '@/views/settings/GeneralSettingsView';
import SecuritySettingsView from '@/views/settings/SecuritySettingsView';
import TwoFactorAuthDisableModal from '@/modals/twoFactorAuth/disable/TwoFactorAuthDisableModal';
import TwoFactorAuthSetupModal from '@/modals/twoFactorAuth/setup/TwoFactorAuthSetupModal';
import EmptyState from '@/components/common/emptyState/EmptyState';
import ButtonLink from '@/components/common/button/ButtonLink';
import { useTranslation } from 'react-i18next';
import CreateTokenModal from '@/modals/token/create/CreateTokenModal';
import ChannelListView from '@/views/channels/list/ChannelListView';
import ChannelView from '@/views/channels/detail/ChannelView';
import RecordingCreateView from './views/recordings/create/RecordingCreateView';
import ChannelSelectModal from './components/channels/selectModal/ChannelSelectModal';

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
                <Route path="/login" element={<LoginView />} />
              </Route>

              <Route element={<Authenticated />}>
                <Route element={<DashboardView />}>
                  <Route
                    path="/"
                    element={<Navigate to={'/channels'} replace />}
                  />
                  <Route path="/channels" element={<ChannelListView />} />
                  <Route path="/channels/:id" element={<ChannelView />} />
                  <Route path="/guide" element={<GuideView />} />
                  <Route path="/guide/events/:id" element={<EventView />} />
                  <Route path="/recordings" element={<RecordingsView />} />
                  <Route
                    path="/recordings/:id"
                    element={<RecordingDetailView />}
                  />
                  <Route
                    path="/recordings/create"
                    element={<RecordingCreateView />}
                  >
                    <Route
                      path="select-channel"
                      element={<ChannelSelectModal />}
                    />
                  </Route>

                  <Route path="/settings" element={<SettingsView />}>
                    <Route
                      path=""
                      element={<Navigate to={'general'} replace />}
                    />
                    <Route path="general" element={<GeneralSettingsView />} />
                    <Route path="security" element={<SecuritySettingsView />}>
                      <Route
                        path="two-factor-auth/disable"
                        element={<TwoFactorAuthDisableModal />}
                      />
                      <Route
                        path="two-factor-auth/setup"
                        element={<TwoFactorAuthSetupModal />}
                      />
                      <Route
                        path="tokens/create"
                        element={<CreateTokenModal />}
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
