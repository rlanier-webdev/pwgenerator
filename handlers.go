package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageHandler(c *gin.Context) {
	password := generatePassword()

	data := gin.H{
		"Title":    "Welcome",
		"Password": password,
	}
	c.HTML(http.StatusOK, "index.html", data)
}

// New API handler function
func PasswordAPIHandler(c *gin.Context) {
	password := generatePassword()

	// Return the generated password as JSON
	c.JSON(http.StatusOK, gin.H{
		"password": password,
	})
}
