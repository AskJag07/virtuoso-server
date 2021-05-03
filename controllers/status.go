package controllers

import (
	"net/http"

	"github.com/AskJag07/virtuoso-server/services"
)

var StatusController = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	message := "Everything's all right!"
	services.ResponseJSON(message, w, http.StatusOK)

})
