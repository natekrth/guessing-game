package guess

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_"github.com/natekrth/guessing-game/orm"
)

var guessingNumber int
var attempts int

var hmacSampleSecret []byte

func GuessHandler(c *gin.Context) {
	var requestBody map[string]int
	hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	header := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(header, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "forbidden", "message": err.Error()})
	}

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

	c.JSON(http.StatusOK, gin.H{"hint": hint, "attempts": attempts, "header": header, "t": tokenString})
}
