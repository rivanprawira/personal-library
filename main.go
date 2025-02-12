package main

import (
	"fmt"
	"log"
	"net/http"
	"personal-library/routes"
)

func main() {
	routes.SetupRoutes()

	fmt.Println("Server berjalan pada URL: http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
