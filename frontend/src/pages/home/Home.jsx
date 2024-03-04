import React, { useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { UserStateContext } from "../../App";

function Home() {
  const { userState } = useContext(UserStateContext);
  const [guess, setGuess] = useState("");
  const [message, setMessage] = useState("");
  const [attempt, setAttempt] = useState(0);
  const [answer, setAnswer] = useState("");
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
          Authorization: `Bearer ${userState}`, // Include the token in the Authorization header
        },
        body: JSON.stringify({ guess: guessedNumber }),
      });

      if (response.status === 401) {
        navigate("/"); // Navigate back to the login page if the token is invalid
        return;
      }

      // Read the response body only once and store it in a variable
      const responseData = await response.json();

      if ((response.status === 200) | (response.status === 201)) {
        setMessage(responseData.message);
        setAttempt(responseData.attempts);
        resetAnswerAndMessageAfterDelay()
      } else {
        setMessage(responseData.error);
      }
    } catch (error) {
      // console.error("Error:", error);
      setMessage("An error occurred while processing your guess.");
      resetAnswerAndMessageAfterDelay()
    }
  };

  const handleGetAnswerSubmit = async (event) => {
    event.preventDefault();
    try {
      const response = await fetch("http://localhost:8080/guess/ans", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${userState}`, // Include the token in the Authorization header
        }
      });
      // console.log(response);
      if (response.status === 401) {
        navigate("/"); // Navigate back to the login page if the token is invalid
        return;
      }
      const responseData = await response.json();

      if (response.status === 200) {
        setAnswer(responseData.answer);
        resetAnswerAndMessageAfterDelay()
      }
    } catch (error) {
      // console.error("Error:", error);
      setMessage("An error occurred while processing answer.");
      resetAnswerAndMessageAfterDelay()
    }
  }

  const resetAnswerAndMessageAfterDelay = () => {
    setTimeout(() => {
      setAnswer("");
      setMessage("");
    }, 5000);
  };

  return (
    <div class="d-flex justify-content-center align-items-center vh-100">
      <div class="col-xs-1 mt-20 align-items-center text-center">
        <h2 class="">Guessing Game</h2>
        <p>Guess a number between 1 to 10</p>
        <p>Attempts: {attempt}</p>
        <p>{message}</p>
        <form onSubmit={handleSubmit} class="d-flex flex-sm-column gap-3 mb-3">
          <input
            type="text"
            value={guess}
            onChange={handleChange}
            placeholder="Enter your guess"
            required
          />
          <button type="submit" class="btn btn-primary">
            Guess
          </button>
        </form>
        <form onSubmit={handleGetAnswerSubmit}>
          <button type="submit" class="btn btn-danger">
            Get Answer
          </button>
        </form>
        <p>Answer: {answer}</p>
      </div>
    </div>
  );
}

export default Home;
