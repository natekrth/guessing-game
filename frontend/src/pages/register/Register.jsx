import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

function Register({ handleRegister }) {
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
      const response = await fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });
    //   console.log(response);
      if (
        response.status === 409 ||
        response.status === 401 ||
        response.status === 404
      ) {
        setNotify("Register Failed");
        resetNotifyAfterDelay();
      } else if (response.status === 200) {
        setNotify("Register Successful");
        navigate("/home"); // Redirect to Home page
      }
    } catch (error) {
      setNotify("Failed to Register. Please try again.");
      resetNotifyAfterDelay();
    }
  };

  const resetNotifyAfterDelay = () => {
    setTimeout(() => {
      setNotify("");
    }, 5000);
  };

  return (
    <div class="d-flex justify-content-center align-items-center vh-100">
      <div class="col-xs-1 text-center">
        <h2>Register</h2>
        <p>Guessing Game</p>
        <p style={{ color: "red" }}>{notify}</p>
        <form onSubmit={handleSubmit} class='grid gap-3"'>
          <div class="p-2 g-col-6">
            <label>Username:</label>
            <input
              type="text"
              value={username}
              onChange={handleUsernameChange}
              placeholder="username"
              required
            />
          </div>
          <div class="p-2 g-col-6">
            <label>Password:</label>
            <input
              type="password"
              value={password}
              onChange={handlePasswordChange}
              placeholder="password"
              required
            />
          </div>
          <button type="submit" class="btn btn-primary">
            Register
          </button>
        </form>
      </div>
    </div>
  );
}

export default Register;
