package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/AskJag07/virtuoso-server/helpers"
	"github.com/AskJag07/virtuoso-server/models"
)

var validate = validator.New()

func Register(client *mongo.Client) gin.HandlerFunc {

	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		usersCollection := client.Database("App").Collection("users")

		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error":  err.Error(),
					"status": http.StatusBadRequest,
				},
			)
			return
		}

		validationErr := validate.Struct(user)
		defer cancel()
		if validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error":  validationErr.Error(),
					"status": http.StatusBadRequest,
				},
			)
			return
		}

		count, err := usersCollection.CountDocuments(ctx, bson.M{"email": *user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error":  "Unable to check for user email.",
					"status": http.StatusInternalServerError,
				},
			)
			return
		}

		password := helpers.HashPassword(*user.Password)
		user.Password = &password

		if count > 0 {
			c.JSON(
				http.StatusConflict,
				gin.H{
					"error":  "this email already exists",
					"status": http.StatusConflict,
				},
			)
			return
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		user.Admin = false
		token, _ := helpers.GenerateAllTokens(*user.Full_name, *user.Session, user.User_id, user.Admin)
		user.Token = &token

		_, insertErr := usersCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error":  "Unable to register new user.",
					"status": http.StatusInternalServerError,
				},
			)
			return
		}
		defer cancel()

		c.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"token":  user.Token,
			},
		)

	}

}
