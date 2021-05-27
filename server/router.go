package server

import (
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
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

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://govirtuoso.org"},
		AllowMethods: []string{"OPTIONS", "POST", "GET"},
		AllowHeaders: []string{"Content-Type", "token"},
	}))

	router := r.Group("/api")

	router.GET("/", controllers.Status())

	router.POST("/auth/register", controllers.Register(client))
	router.POST("/auth/login", controllers.Login(client))

	router.Use(middleware.Authentication(client))

	router.GET("/students", controllers.Students(client))

	return r

}
