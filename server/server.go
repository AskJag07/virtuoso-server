package server

import (
	"log"
	"net/http"
)

func Init() {

	r := NewRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}
