package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/darahayes/go-boom"
	"sample/models"
)
 
type JsonResponseCollection struct {
    Data    []*models.User `json:"data"`
    Message string 		`json:"message"`
}

type JsonResponse struct {
    Data    models.User `json:"data"`
    Message string 		`json:"message"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers(r)

	if err != nil {
		boom.NotFound(w, "No Data Found");
		return
	}

	var response = JsonResponseCollection{Message: "success", Data: users}

    json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := models.GetUser(r, id)
	
	if err != nil {
		boom.NotFound(w, "No Data Found");
		return
	}

    var response = JsonResponse{Message: "success", Data: user}

    json.NewEncoder(w).Encode(response)
}