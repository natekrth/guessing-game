package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var guessingNumber int
var attempts       int


func guessHandler(c *gin.Context) {
	var requestBody map[string]int
	if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // Get the guess from the JSON body
    guess, ok := requestBody["guess"]
    if !ok || guess < 1 || guess > 10 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "The guess should be between 1-10"})
        return
    }

	// Check if the user guess is correct
	if guess == guessingNumber {
		// Regenerate the guess number
		guessingNumber = rand.Intn(10)
		numAttempts := attempts
		attempts = 0 // Reset attempts on correct guess
		c.JSON(http.StatusCreated, gin.H{
			"message":  "Congratulations! You guessed it right!",
			"attempts": numAttempts,
		})
		return
	}

	// Increment the attempts
	attempts++

	hint := "Try again, the number is "
	if guess < guessingNumber {
		hint += "higher"
	} else {
		hint += "lower"
	}

	c.JSON(http.StatusOK, gin.H{"hint": hint, "attempts": attempts})
}

func main() {
	// Initialize Gin
	router := gin.Default()

	router.POST("/guess", guessHandler)

	// Initialize the guess number
	guessingNumber = rand.Intn(10)

	// Start the server
	router.Run(":8080")
}
