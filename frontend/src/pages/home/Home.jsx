import React, { useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { UserStateContext } from "../../App";
import Button from 'react-bootstrap/Button';

function Home() {
  const { userState } = useContext(UserStateContext);
  const [guess, setGuess] = useState("");
  const [message, setMessage] = useState("");
  const [attempt, setAttempt] = useState(0);
  const navigate = useNavigate();

  const handleChange = (event) => {
    setGuess(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    const guessedNumber = parseInt(guess);
    if (isNaN(guessedNumber)) {
      setMessage("Please enter a valid number.");
      return;
    }

    try {
      const response = await fetch("http://localhost:8080/guess", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${userState}`, // Include the token in the Authorization header
        },
        body: JSON.stringify({ guess: guessedNumber }),
      });

      if (response.status === 401) {
        navigate("/"); // Navigate back to the login page if the token is invalid
        return;
      }

      // Read the response body only once and store it in a variable
      const responseData = await response.json();

      if (response.status === 200 | response.status === 201) {
        setMessage(responseData.message);
        setAttempt(responseData.attempt);
      } else {
        setMessage(responseData.error);
      }
    } catch (error) {
      console.error("Error:", error);
      setMessage("An error occurred while processing your guess.");
    }
  };

  return (
    <div>
      <h2>Guess the Number</h2>
      <p>Guess a number between 1 to 10</p>
      <p>Attemps: {attempt}</p>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={guess}
          onChange={handleChange}
          placeholder="Enter your guess"
          required
        />
        <button type="submit" class="btn btn-primary">Guess</button>
      </form>
      <p>{message}</p>
    </div>
  );
}

export default Home;
