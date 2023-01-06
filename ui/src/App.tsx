import { useEffect } from 'react';
import {
  BrowserRouter,
  Routes,
  Route,
  Navigate,
  Outlet,
} from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import { useAuth } from './contexts/AuthContext';
import useLogout from './hooks/logout';
import AuthProvider from './providers/AuthProvider';
import { ThemeProvider } from './providers/ThemeProvider';
import LoginView from './views/LoginView/LoginView';
import { useTheme } from './contexts/ThemeContext';

import 'react-toastify/dist/ReactToastify.css';
import ChannelListView from './views/ChannelListView/ChannelListView';
import DashboardView from './views/DashboardView/DashboardView';
import GuideView from './views/GuideView/GuideView';
import EventView from './views/EventView/EventView';
import RecordingsView from './views/RecordingsView/RecordingsView';
import RecordingDetailView from './views/RecordingDetailView/RecordingDetailView';

type AuthenticationCheckerProps = {
  redirect?: string;
};

function Unauthenticated(props: AuthenticationCheckerProps) {
  const authContext = useAuth();
  const isAuthenticated = !!authContext.username;

  return isAuthenticated ? <Navigate to={props.redirect || '/'} /> : <Outlet />;
}

function Authenticated(props: AuthenticationCheckerProps) {
  const authContext = useAuth();
  const isAuthenticated = !!authContext.username;

  return isAuthenticated ? (
    <Outlet />
  ) : (
    <Navigate to={props.redirect || '/login'} />
  );
}

function Logout() {
  const { logout, loading } = useLogout();

  useEffect(() => {
    logout();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

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

function App() {
  return (
    <ThemeProvider>
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route element={<Unauthenticated />}>
              <Route path="/login" element={<LoginView />} />
            </Route>

            <Route element={<Authenticated />}>
              <Route element={<DashboardView />}>
                <Route path="/" element={<ChannelListView />} />
                <Route path="/guide" element={<GuideView />} />
                <Route path="/guide/events/:id" element={<EventView />} />
                <Route path="/recordings" element={<RecordingsView />} />
                <Route
                  path="/recordings/:id"
                  element={<RecordingDetailView />}
                />
                <Route path="/settings" element={<></>} />
              </Route>

              <Route path="/logout" element={<Logout />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </AuthProvider>
      <Notification />
    </ThemeProvider>
  );
}

export default App;
