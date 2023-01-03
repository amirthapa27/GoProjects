package routes

import (
	"github.com/amirthapa27/golangJWT/controllers"
	"github.com/gin-gonic/gin"
)

// create a func AuthRoutes
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup()) //will handle the request for signing up a user
	incomingRoutes.POST("users/login", controllers.Login())   //will handle the request for logging in
}
