package router

import (
	"GoMicroservice/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	//router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	//router.HandleFunc("/api/users/{id}", controllers.GetUsersById).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	//router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	//router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")

}
