package server

import (
	"log"
	"net/http"
)

func Init() {

	corsWrapper := NewCorsWrapper()
	r := NewRouter()

	if err := http.ListenAndServe(":8080", corsWrapper.Handler(r)); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}
