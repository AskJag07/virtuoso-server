package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/AskJag07/virtuoso-server/helpers"
	"github.com/AskJag07/virtuoso-server/models"
)

func Login(client *mongo.Client) gin.HandlerFunc {

	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		var user models.User
		var foundUser models.User

		usersCollection := client.Database("App").Collection("users")

		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
			return
		}

		err := usersCollection.FindOne(ctx, bson.M{"email": strings.ToLower(*user.Email)}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(
				http.StatusOK,
				gin.H{"error": "Email not found. Please register."},
			)
			return
		}

		passwordIsValid := helpers.VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(
				http.StatusOK,
				gin.H{"error": "Incorrect password."},
			)
			return
		}

		token, _ := helpers.GenerateAllTokens(*foundUser.Full_name, *foundUser.Session, foundUser.User_id, foundUser.Admin)

		helpers.UpdateAllTokens(token, foundUser.User_id, client)

		c.JSON(
			http.StatusOK,
			gin.H{
				"token":     foundUser.Token,
			},
		)

	}

}
