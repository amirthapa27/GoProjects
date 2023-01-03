package main

import (
	"fmt"

	"github.com/amirthapa27/JWTginGorm/controllers"
	"github.com/amirthapa27/JWTginGorm/intializers"
	"github.com/amirthapa27/JWTginGorm/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectDB()
	intializers.SyncDatabase()
} //

func main() {

	router := gin.Default()

	router.POST("/signup", controllers.Signup)

	router.POST("/login", controllers.Login)

	router.GET("/validate", middleware.RequirAuth, controllers.Validate)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.Run()
	fmt.Println("Hello2")

}
