package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type friend struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Workplace string `json:"workplace"`
}

var friends = []friend{
	{ID: "1", FullName: "John Doe", Email: "johndoe@gmail.com", Workplace: "WWW"},
	{ID: "2", FullName: "Mavin Picks", Email: "mavpicks@gmail.com", Workplace: "Dell"},
	{ID: "3", FullName: "Jonah Paul", Email: "jpaul@gmail.com", Workplace: "YouTube"},
}

func getFriends(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, friends)
}

func main() {
	router := gin.Default()

	router.GET("/friends", getFriends)
	router.Run("localhost:8080")
}
