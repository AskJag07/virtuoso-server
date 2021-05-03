package controllers

import (
	"net/http"
	"strings"

	"github.com/AskJag07/virtuoso-server/services"
)

var TestController = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	token := authHeaderParts[1]

	hasScope := services.CheckScope("read:messages", token)

	if !hasScope {
		message := "Insufficient scope."
		services.ResponseJSON(message, w, http.StatusForbidden)
		return
	}

	message := "Authentication is required."
	services.ResponseJSON(message, w, http.StatusOK)

})
