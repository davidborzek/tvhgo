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
import Login from './views/Login/Login';
import { useTheme } from './contexts/ThemeContext';

import 'react-toastify/dist/ReactToastify.css';
import ChannelList from './views/ChannelList/ChannelList';
import Dashboard from './views/Dashboard/Dashboard';
import Guide from './views/Guide/Guide';
import Event from './views/Event/Event';
import Recordings from './views/Recordings/Recordings';

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
              <Route path="/login" element={<Login />} />
            </Route>

            <Route element={<Authenticated />}>
              <Route element={<Dashboard />}>
                <Route path="/" element={<ChannelList />} />
                <Route path="/guide" element={<Guide />} />
                <Route path="/guide/events/:id" element={<Event />} />
                <Route path="/recordings" element={<Recordings />} />
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
