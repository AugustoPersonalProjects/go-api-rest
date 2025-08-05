package main

import (
	"fmt"

	"github.com/lopeslyra/go-api-rest/models"
	"github.com/lopeslyra/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Personalidade 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Personalidade 2", Historia: "Historia 2"},
	}

	fmt.Println("Server starting on port 8000...")
	routes.HandleRequests()
}