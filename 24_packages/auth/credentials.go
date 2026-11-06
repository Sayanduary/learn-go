package auth

import (
	"fmt"
)

func LoginWithCredentials(username string, password string) {
	fmt.Println("Logged in using: ", username, password)
}

func extractSession() string {
	return "logged in"
}
