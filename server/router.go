package server

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/AskJag07/virtuoso-server/controllers"
	"github.com/AskJag07/virtuoso-server/middlewares"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.Handle("/", controllers.StatusController).Methods("GET")

	jwtMiddleware := middlewares.AuthMiddleware()
	r.Handle("/restricted", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(controllers.TestController))).Methods("GET")

	return r

}
