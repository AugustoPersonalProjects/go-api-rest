package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lopeslyra/go-api-rest/controllers"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", r))
}