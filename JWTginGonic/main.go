package main

import (
	"log"
	"os"

	"github.com/amirthapa27/golangJWT/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env") //Load will read your env file(s) and load them into ENV for this process.
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
	port := os.Getenv("PORT") //use port defeined
	//if port in env is  nill use default port 8000
	if port == "" {
		port = "8000"
	}
	//define a router using gin-gonice
	router := gin.New()
	//use the gin logger
	router.Use(gin.Logger())

	// from routes import  the routes function
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	//create api
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to api-1"}) //status code and message
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to api-2"})
	})
	//start the api
	router.Run(":" + port)
}
