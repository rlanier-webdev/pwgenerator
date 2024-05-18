package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MinLength        = 15 // Minimum length of the password
)

// getSeed initializes a new random number generator with a seed based on the current time
func getSeed() *rand.Rand {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	generator := rand.New(source)
	return generator
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
	encounteredLetter := false           // Flag to track if a letter has been encountered yet

	for i := range password {
		index := generator.Intn(len(temp))
		char := rune(temp[index])

		// Check if the selected character is a letter
		if isLetter(char) && !encounteredLetter{
			// Convert the letter to uppercase
			char = toUpper(char)
			encounteredLetter = true
		}

		password[i] = char
	}
	return string(password)
}

func main() {
	// Code
	var (
		alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers  string = "0123456789"
		symbols  string = "!@#$%^&*()_+-=[]{}|;:,.<>?/~"
	)

	password := getPassword(alphabet, numbers, symbols)
	fmt.Println(password)
}
