package server

import (
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/AskJag07/virtuoso-server/config"
	"github.com/AskJag07/virtuoso-server/controllers"
)

func NewRouter(client *mongo.Client) *gin.Engine {

	Production := config.GetVar("PRODUCTION")

	production, err := strconv.ParseBool(Production)
	if err != nil {
		log.Panic(err)
	}

	if production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8000"},
		AllowMethods: []string{"OPTIONS", "POST", "GET"},
		AllowHeaders: []string{"Content-Type"},
	}))

	router.GET("/", controllers.Status())

	router.POST("/auth/register", controllers.Register(client))
	router.POST("/auth/login", controllers.Login(client))

	return router

}
