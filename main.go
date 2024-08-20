package main

import (
	"log"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
)

const (
	MinLength = 15 // Minimum length of the password
)

func getSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func toUpper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}

func getPassword(alphabet, numbers, symbols string) string {
	temp := alphabet + numbers + symbols

	generator := getSeed()

	password := make([]rune, MinLength)
	encounteredLetter := false

	for i := range password {
		index := generator.Intn(len(temp))
		char := rune(temp[index])

		if isLetter(char) && !encounteredLetter {
			char = toUpper(char)
			encounteredLetter = true
		}

		password[i] = char
	}
	return string(password)
}

func generatePassword() string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()_+-=[]{}|;:,.<>?/~"
	return getPassword(alphabet, numbers, symbols)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")
	r.GET("/", PageHandler)

	// New API route
	r.GET("/api/password", PasswordAPIHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
