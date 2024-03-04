package guess

import (
	_"fmt"
	"math/rand"
	"net/http"
	_"os"
	_"strings"

	"github.com/gin-gonic/gin"
	_"github.com/golang-jwt/jwt"
	_ "github.com/natekrth/guessing-game/orm"
)

var guessingNumber = rand.Intn(10)
var attempts int


func GuessHandler(c *gin.Context) {
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
	// Check if the guess is correct
	if guess == guessingNumber {
		// Regenerate the number
		guessingNumber = rand.Intn(10)
		numAttempts := attempts
		attempts = 0 // Reset attempts on correct

		// Return HTTP 201 Created with the JSON response
		c.JSON(http.StatusCreated, gin.H{
			"message":  "YES! You guessed it right!",
			"attempts": numAttempts,
		})
		return
	}

	// Increment the attempts
	attempts++

	// Return a hint if the guess is wrong
	hint := "Try again, the number is "
	if guess < guessingNumber {
		hint += "higher"
	} else {
		hint += "lower"
	}

	c.JSON(http.StatusOK, gin.H{"message": hint, "attempts": attempts})
}

func GuessAnswer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"answer": guessingNumber})
}