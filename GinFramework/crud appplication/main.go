package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

// Get all users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// Get a single user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func GetUserByQuery(c *gin.Context) {
	username := c.Query("username")

	var filttereduser []User

	for _, user := range users {
		if user.Username != "" && user.Username == username {
			filttereduser = append(filttereduser, user)
		}
	}
	c.JSON(http.StatusOK, filttereduser)

}

// Create a new user
func CreateUser(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = len(users) + 1
	users = append(users, user)

	c.JSON(http.StatusCreated, user)
}

// Update an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			var updatedUser User

			if err := c.BindJSON(&updatedUser); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			updatedUser.ID = user.ID
			users[i] = updatedUser

			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

// Delete a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func main() {
	r := gin.Default()

	// Initialize the users variable
	users = []User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}

	// Define routes for user API
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUserByID)
	r.GET("/usersByQuery", GetUserByQuery)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.Run(":8080")
}
