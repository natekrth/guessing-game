import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

function Login({ handleLogin }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [notify, setNotify] = useState("");
  const navigate = useNavigate();

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
      });

      if (response.status === 409 || response.status === 401 || response.status === 404) {
        setNotify("Login Failed");
        resetNotifyAfterDelay();
      } else if (response.status === 200) {
        const data = await response.json();
        handleLogin(data.token); // Update userState with token
        navigate("/home"); // Redirect to Home page
      }
    } catch (error) {
      setNotify("Failed to login. Please try again.");
      resetNotifyAfterDelay();
    }
  };

  const resetNotifyAfterDelay = () => {
    setTimeout(() => {
      setNotify("");
    }, 5000);
  };
  
  return (
    <div>
      <h2>Login</h2>
      <p>{notify}</p>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Username:</label>
          <input
            type="text"
            value={username}
            onChange={handleUsernameChange}
            placeholder="username"
            required
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={handlePasswordChange}
            placeholder="password"
            required
          />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}

export default Login;
