package main

import (
	"os"

	"github.com/amir27/restBKD/middleware"
	"github.com/amir27/restBKD/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//create a db for food
	// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

	port := os.Getenv("PORT") //getting the port

	//if port is nil then by default port will be 8000
	if port == "" {
		port = "8000"
	}

	//create a gin router
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItem(router)

	//start the router
	router.Run(":" + port)

}
