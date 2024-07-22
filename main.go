package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	MinLength = 15 // Minimum length of the password
)

// getSeed initializes a new random number generator with a seed based on the current time
func getSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Function to check if a character is a letter
func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

// Function to convert a letter to uppercase
func toUpper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		// Convert lowercase letter to uppercase
		return char - 'a' + 'A'
	}
	return char
}

func getPassword(alphabet, numbers, symbols string) string {
	temp := alphabet + numbers + symbols

	generator := getSeed()

	password := make([]rune, MinLength) // Create rune array for password
	encounteredLetter := false          // Flag to track if a letter has been encountered yet

	for i := range password {
		index := generator.Intn(len(temp))
		char := rune(temp[index])

		// Check if the selected character is a letter
		if isLetter(char) && !encounteredLetter {
			// Convert the letter to uppercase
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
	// Route to serve the initial page
	r.GET("/", PageHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
