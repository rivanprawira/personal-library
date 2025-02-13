package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"personal-library/routes"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	expectedUsername := os.Getenv("USERNAME_SERVER")
	expectedPassword := os.Getenv("PASSWORD_SERVER")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	inputUsername, _ := reader.ReadString('\n')
	inputUsername = strings.TrimSpace(inputUsername)
	fmt.Print("Enter password: ")
	inputPassword, _ := reader.ReadString('\n')
	inputPassword = strings.TrimSpace(inputPassword)

	if inputUsername != expectedUsername || inputPassword != expectedPassword {
		log.Fatal("Invalid username or password")
	}

	routes.SetupRoutes()
	port := os.Getenv("PORT")
	log.Printf("Login successful. Starting server on URL: http://localhost:%s...", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
