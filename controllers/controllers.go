package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/lopeslyra/go-api-rest/database"
    "github.com/lopeslyra/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    var p []models.Personalidade
    database.DB.Find(&p)
    json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var personalidade models.Personalidade
    database.DB.First(&personalidade, id)
    json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
    var personalidade models.Personalidade
    json.NewDecoder(r.Body).Decode(&personalidade)
    database.DB.Create(&personalidade)
    json.NewEncoder(w).Encode(personalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    var personalidade models.Personalidade
    database.DB.Delete(&personalidade, id)
    json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    var personalidade models.Personalidade
    json.NewDecoder(r.Body).Decode(&personalidade)
    database.DB.First(&personalidade, id)
    json.NewEncoder(w).Encode(personalidade)
}