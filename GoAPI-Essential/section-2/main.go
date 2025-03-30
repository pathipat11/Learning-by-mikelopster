package main

import (
	"fmt"
)

// LoginError is a custom error type for login failures
type LoginError struct {
	Username string
	Message  string
}

// Implement the Error() method to satisfy the error interface
func (e *LoginError) Error() string {
	return fmt.Sprintf("Login error for user '%s': %s", e.Username, e.Message)
}

// Simulated function that attempts a user login
func login(username, password string) error {
	// Simulate checking username and password
	if username != "admin" || password != "password123" {
		return &LoginError{Username: username, Message: "invalid credentials"}
	}
	// Login successful
	return nil
}

func main() {
	// Attempt to login with wrong credentials
	err := login("user", "pass")
	if err != nil {
		switch e := err.(type) {
		case *LoginError:
			// Custom error handling
			fmt.Println("Custom error occurred:", e)
		default:
			// Other types of errors
			fmt.Println("Generic error occurred:", e)
		}
		return
	}

	// Continue with the rest of the program if login is successful
	fmt.Println("Login successful!")
}
