package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"webSurfer/db"
	"webSurfer/schemas"
)

func GetUsers(c *gin.Context) {
	client := db.GetClient()
	collection := client.Database("main").Collection("users")

	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	defer cursor.Close(context.Background())

	var users []schemas.User
	for cursor.Next(context.Background()) {
		var person schemas.User
		if err := cursor.Decode(&person); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
			return
		}
		users = append(users, person)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Success",
		"data":    users,
	})
}

func GetUser(c *gin.Context) {
	c.String(http.StatusOK, "Get a specific user")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "Create a new user")
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "Delete a user")
}
