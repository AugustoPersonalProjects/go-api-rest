package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lopeslyra/go-api-rest/controllers"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}