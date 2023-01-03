package routes

import (
	"github.com/amir27/restBKD/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItem(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByORder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItems_id", controllers.UpdateOrderItem())

}
