package middleware

import (
	"fmt"
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
				http.StatusInternalServerError,
				gin.H{"error": fmt.Sprintf("No Authorization header provided")},
			)
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.Full_name)
		c.Set("last_name", claims.School)
		c.Set("last_name", claims.Standard)
		c.Set("uid", claims.Uid)

		c.Next()

	}

}
