package server

import (
	"github.com/rs/cors"
)

func NewCorsWrapper() *cors.Cors {

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
		AllowedOrigins: []string{"http://localhost:8000"},
	})

	return corsWrapper

}
