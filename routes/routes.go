package routes

import (
	"log"
	"net/http"

	"github.com/lopeslyra/go-api-rest/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}