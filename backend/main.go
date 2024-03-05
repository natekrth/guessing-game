package main

import (
	"fmt"
	AuthController "github.com/natekrth/guessing-game/controllers/auth"
	GuessController "github.com/natekrth/guessing-game/controllers/guess"
	"github.com/natekrth/guessing-game/orm"
	"github.com/natekrth/guessing-game/middleware"
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
	// CORS configuration
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"} // Allow all origins
    config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
    config.AllowHeaders = []string{"Authorization", "Content-Type"} // Allow Authorization header
    r.Use(cors.New(config))

	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	r.DELETE("/user/delete", middleware.JWTAuthen(), AuthController.DeleteUser)
	r.POST("/guess", middleware.JWTAuthen(), GuessController.GuessHandler)
	r.GET("/guess/ans", middleware.JWTAuthen(), GuessController.GuessAnswer)
	r.PATCH("/guess/update", GuessController.UpdateAnswer)
	r.Run("localhost:8080")
}
