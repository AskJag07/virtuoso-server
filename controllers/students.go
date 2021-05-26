package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AskJag07/virtuoso-server/models"
)

func Students(client *mongo.Client) gin.HandlerFunc {

	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		var students []models.Student

		usersCollection := client.Database("App").Collection("users")

		projection := bson.D{
			{Key: "full_name", Value: 1},
			{Key: "email", Value: 1},
			{Key: "standard", Value: 1},
			{Key: "created_at", Value: 1},
			{Key: "updated_at", Value: 1},
		}
		cur, err := usersCollection.Find(ctx, bson.M{"admin": false}, options.Find().SetProjection(projection))
		defer cancel()
		if err != nil {
			c.JSON(
				http.StatusOK,
				gin.H{"error": "Students not found."},
			)
			return
		}

		if err = cur.All(ctx, &students); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err},
			)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"students": students,
		})

	}
}
