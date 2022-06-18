package controllers

import (
	"net/http"
	"encoding/json"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("profile")
}
