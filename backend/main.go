// package main

// import (
// 	// "encoding/json"
// 	"math/rand"
// 	"net/http"

//     "github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// var guessingNumber int
// var attempts       int
// var token          string

// type User struct {
//     Username string
//     Password string
// }

// func loginHandler(c *gin.Context) {
// 	// Dummy login logic, just return a token
// 	token = "test123"
// 	c.Header("Authorization", token)
// 	c.JSON(http.StatusOK, gin.H{"token": token})
//     return
// }

// func authMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the token from the request header
// 		token := c.GetHeader("Authorization")
// 		// Dummy token validation, just check if it's not empty
//         expectedToken := "test123"
// 		if token != expectedToken {
// 			c.JSON(http.StatusUnauthorized, gin.H{
//                 "error":  "Unauthorized",
//                 "token": token,
//             })
// 			c.Abort()
// 			return
// 		}

// 		// Continue processing if the token is valid
// 		c.Next()
// 	}
// }

// func guessHandler(c *gin.Context) {
// 	// Parse the guessed number
// 	var requestBody map[string]int
// 	if err := c.ShouldBindJSON(&requestBody); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
//         return
//     }

//     // Get the guess from the JSON body
//     guess, ok := requestBody["guess"]
//     if !ok || guess < 1 || guess > 10 {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "The guess should be between 1-10"})
//         return
//     }

// 	// Check if the guess is correct
// 	if guess == guessingNumber {
// 		// Regenerate the number
// 		guessingNumber = rand.Intn(10)
// 		numAttempts := attempts
// 		attempts = 0 // Reset attempts on correct

//         // Return HTTP 201 Created with the JSON response
// 		c.JSON(http.StatusCreated, gin.H{
// 			"message":  "YES! You guessed it right!",
// 			"attempts": numAttempts,
// 		})
// 		return
// 	}

// 	// Increment the attempts
// 	attempts++

// 	// Return a hint if the guess is wrong
// 	hint := "Try again, the number is "
// 	if guess < guessingNumber {
// 		hint += "higher"
// 	} else {
// 		hint += "lower"
// 	}

// 	c.JSON(http.StatusOK, gin.H{"hint": hint, "attempts": attempts})
// }

// func main() {
// 	// Initialize Gin
// 	router := gin.Default()
//     router.Use(cors.Default())
// 	// Routes
// 	router.POST("/login", loginHandler)

// 	// Use middleware to protect the guess endpoint
// 	router.Use(authMiddleware())

// 	router.POST("/guess", guessHandler)

// 	// Initialize the guess number
// 	guessingNumber = rand.Intn(10)

// 	// Start the server
// 	router.Run(":8080")
// }

package main

import (
	"fmt"
	AuthController "github.com/natekrth/guessing-game/controllers/auth"
	GuessController "github.com/natekrth/guessing-game/controllers/guess"
	"github.com/natekrth/guessing-game/orm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	gorm.Model
	Username string
	Password string
}

var guessingNumber int

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	r.POST("/guess", GuessController.GuessHandler)
	r.Run("localhost:8080")
}
