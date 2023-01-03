package middleware

import (
	"fmt"
	"net/http"

	"github.com/amirthapa27/golangJWT/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token") //get the first value asscoited with the key("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization provided")})
			c.Abort() //Abort prevents pending handlers from being called.
			return
		}
		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}
