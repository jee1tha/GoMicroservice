package controllers

import (
	"GoMicroservice/entities"
	"GoMicroservice/service"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	service.CreateUserService(user)
	json.NewEncoder(w).Encode(user)
}
