package controllers

import "net/http"

var StatusController = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Working!"))

})
