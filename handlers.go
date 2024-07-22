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
