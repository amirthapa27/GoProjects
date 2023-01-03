package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/amirthapa27/JWTginGorm/intializers"
	"github.com/amirthapa27/JWTginGorm/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequirAuth(c *gin.Context) {
	//Get the cookie from request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

	}
	//Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

		}
		//Find the user with token sub
		var user models.User
		intializers.DB.Find(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//Attach to req
		c.Set("user", user)

		//continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		// fmt.Println(err)
	}

}
