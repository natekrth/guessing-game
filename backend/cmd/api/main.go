package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []user{
	{Username: "test", Password: "test1234"},
    {Username: "nate", Password: "nate1234"},
}

// getAlbums responds with the list of all albums as JSON.
func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:8080")
}
