package server

import (
	"github.com/AskJag07/virtuoso-server/config"
	"github.com/AskJag07/virtuoso-server/db"
)

func Init() {

	port := config.GetVar("PORT")

	if port == "" {
		port = "8888"
	}

	client := db.Init()
	router := NewRouter(client)

	router.Run(":" + port)

}
