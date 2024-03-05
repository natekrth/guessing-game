package auth

import (
	"fmt"
	"github.com/natekrth/guessing-game/orm"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check User Exists
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "User Already Exists"})
		return
	}
	// Create User
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := orm.User{Username: json.Username, Password: string(encryptedPassword)}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Create Success", "userId": user.ID})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User Create Failed"})
	}
}

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check User Exists
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User Does Not Exists"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExist.ID,
			"exp":    time.Now().Add(time.Minute * 1).Unix(), // expiration time for token
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Login Success", "token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Login Failed"})
	}
}

type DeleteBody struct {
	Username string `json:"username" binding:"required"`
}

func DeleteUser(c *gin.Context) {
    var requestBody DeleteBody
	
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if user exists
    var user orm.User
    if err := orm.Db.Where("username = ?", requestBody.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
        return
    }

    if err := orm.Db.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete user"})
        fmt.Println("Error deleting user:", err)
        return
    }

	orm.Db.Where("username = ?", requestBody.Username).Delete(&user)
    c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User deleted successfully"})
}
