package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"personal-library/backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if username != os.Getenv("USERNAME_SERVER") || password != os.Getenv("PASSWORD_SERVER") {
		log.Fatal("Invalid username or password")
	}

	routes.SetupRoutes()
	port := os.Getenv("PORT")
	server := &http.Server{Addr: ":" + port}

	stopServer := make(chan struct{})

	go func() {
		log.Printf("Server started on http://localhost:%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	initialToken := os.Getenv("TOKEN")
	go func() {
		for {
			time.Sleep(5 * time.Second) // Check every 5 seconds
			file, err := os.ReadFile(".env")
			if err != nil {
				log.Println("Error reading .env file:", err)
				continue
			}

			lines := strings.Split(string(file), "\n")
			currentToken := ""
			for _, line := range lines {
				if strings.HasPrefix(line, "TOKEN=") {
					currentToken = strings.TrimPrefix(line, "TOKEN=")
					currentToken = strings.TrimSpace(currentToken)
					break
				}
			}

			if currentToken != initialToken {
				log.Println("Token changed. Shutting down server...")
				close(stopServer)
				return
			}
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalChan:
		log.Println("\nReceived shutdown signal. Stopping server...")
	case <-stopServer:
		log.Println("\nToken changed. Stopping server...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped gracefully")
}
