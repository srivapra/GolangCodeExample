package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	BookId     int    `json:"bookId"`
	BookTitle  string `json:"bookTitle"`
	BookAuthor string `json:"bookAuthor"`
}

var users []User

func CreateUser(c gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.AbortStatusWithJSON(http.StatusBadRequest, gin.H{"Error : ", err.Error()})
		return
	}
	users = append(users, user)
	c.JSON(http.StatusCreated, gin.H{"User Created Successfully "})
}

func main() {
	router := gin.Default()
	router.POST("/user", CreateUser)
	router.Run(":8080")
}
