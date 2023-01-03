package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/amir27/restBKD/database"
	"github.com/amir27/restBKD/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()
		result, err := orderCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the order items"})
			return
		}
		var allOrder []bson.M
		if err := result.All(ctx, &allOrder); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrder)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		orderId := c.Param("order_id")
		var order models.Order
		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching orders"})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		var table models.Table
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(order)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Table was not found"})
				return
			}
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, err := orderCollection.InsertOne(ctx, order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Order was not created"})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var order models.Order
		var table models.Table
		var updateObj primitive.D
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if order.Table_id != nil {
			err := orderCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Table was not found"})
				return
			}
			updateObj = append(updateObj, bson.E{"menu", order.Table_id})
		}
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

		upsert := true
		orderId := c.Param("order_id")
		filter := bson.M{"order_id": orderId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Order update failed"})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}

func OrderItemOrderCreator(order models.Order) string {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx, order)
	defer cancel()
	return order.Order_id
}
