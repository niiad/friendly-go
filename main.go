package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
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

func addFriend(context *gin.Context) {
	var newFriend friend

	if err := context.BindJSON(&newFriend); err != nil {
		return
	}

	friends = append(friends, newFriend)

	context.IndentedJSON(http.StatusCreated, newFriend)
}

func friendById(context *gin.Context) {
	id := context.Param("id")
	friend, err := getFriendById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "friend not found"})

		return
	}

	context.IndentedJSON(http.StatusOK, friend)
}

func getFriendById(id string) (*friend, error) {
	for i, friend := range friends {
		if friend.ID == id {
			return &friends[i], nil
		}
	}

	return nil, errors.New("friend not found")
}

func main() {
	router := gin.Default()

	router.GET("/friends", getFriends)
	router.GET("/friends/:id", friendById)
	router.POST("/friends", addFriend)

	router.Run("localhost:8080")
}
