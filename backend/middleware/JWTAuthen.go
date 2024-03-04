package middleware

import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
)

func JWTAuthen() gin.HandlerFunc {
    return func(c *gin.Context) {
        hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
        header := c.Request.Header.Get("Authorization")
        if header == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            return
        }
        tokenString := strings.Replace(header, "Bearer ", "", 1)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
            }
            return hmacSampleSecret, nil
        })
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }

        if token == nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            return
        }

        userID, ok := claims["userId"].(string)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
            return
        }

        // Set the user ID in the context
        c.Set("userId", userID)

        c.Next()
    }
}
