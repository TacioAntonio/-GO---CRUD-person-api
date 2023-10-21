package routes

import (
	"github.com/gorilla/mux"

	personsController "go-api/app/controllers"
)

func SetupRoutes() *mux.Router {
	route := mux.NewRouter().StrictSlash(true)

	// # personsController
	route.HandleFunc("/persons", personsController.GetAll).Methods("GET")
	route.HandleFunc("/person/{id}", personsController.GetPersonByID).Methods("GET")
	route.HandleFunc("/person/create", personsController.Create).Methods("POST")
	route.HandleFunc("/person/update/{id}", personsController.Update).Methods("PUT")
	route.HandleFunc("/person/delete/{id}", personsController.Delete).Methods("DELETE")

	return route
}
