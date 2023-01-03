package routes

import (
	"github.com/amirthapa27/golangJWT/controllers"
	"github.com/amirthapa27/golangJWT/middleware"
	"github.com/gin-gonic/gin"
)

// create a func UserRoutes
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())              //USE will make usre that this route is authenticated everytime before it is called with the help of the middleware
	incomingRoutes.GET("/users", controllers.GetUsers())         //will provide us with all the users
	incomingRoutes.GET("/users/:user_id", controllers.GetUser()) //will provide us with the user whose if is mentioned

}
