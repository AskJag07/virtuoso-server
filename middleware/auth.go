package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/AskJag07/virtuoso-server/helpers"
)

func Authentication(client *mongo.Client) gin.HandlerFunc {

	return func(c *gin.Context) {

		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{"error": "No Authorization header provided."},
			)
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("full_name", claims.Full_name)
		c.Set("session", claims.Session)
		c.Set("admin", claims.Admin)
		c.Set("uid", claims.Uid)

		c.Next()

	}

}
