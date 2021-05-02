package server

import (
	"github.com/gorilla/mux"

	"github.com/AskJag07/virtuoso-server/controllers"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.Handle("/", controllers.StatusController).Methods("GET")

	return r

}
