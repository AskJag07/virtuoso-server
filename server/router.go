package server

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/AskJag07/virtuoso-server/config"
	"github.com/AskJag07/virtuoso-server/controllers"
	"github.com/AskJag07/virtuoso-server/middleware"
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

	router.GET("/", controllers.Status())

	router.POST("/auth/register", controllers.Register(client))
	router.POST("/auth/login", controllers.Login(client))

	router.Use(middleware.Authentication(client))

	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	router.GET("/api-2", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-2"})

	})

	return router

}
