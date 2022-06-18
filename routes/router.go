package router

import (
	"github.com/gorilla/mux"
	"sample/controllers"
)

func Routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()

	//User Module routes
	userRoute := router.PathPrefix("/users").Subrouter()
	userRoute.HandleFunc("/", controllers.GetUsers).Methods("GET")
	userRoute.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	userRoute.HandleFunc("/store", controllers.GetUsers).Methods("POST")
	userRoute.HandleFunc("/update", controllers.GetUsers).Methods("PUT")
	userRoute.HandleFunc("/delete", controllers.GetUsers).Methods("DELETE")
	//End User routes

	return router;
}