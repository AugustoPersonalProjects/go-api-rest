package main

import (
	"fmt"

	"github.com/lopeslyra/go-api-rest/routes"
)

func main() {
	fmt.Println("Server starting on port 8000...")
	routes.HandleRequests()
}