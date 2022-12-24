import { useState } from "react";
import useLogin from "../../hooks/login";

export default function Login() {
  const { login, error, loading } = useLogin();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  return (
    <div className="Login">
      <p>Login</p>
      <p>
        <input
          type="text"
          value={username}
          onChange={(evt) => {
            setUsername(evt.target.value);
          }}
        />
      </p>
      <p>
        <input
          type="password"
          onChange={(evt) => {
            setPassword(evt.target.value);
          }}
        />
      </p>
      <p>
        <button onClick={() => login(username, password)}>Login</button>
      </p>
    </div>
  );
}
