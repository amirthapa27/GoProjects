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

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the invoice items"})
			return
		}
		var allInvoices []bson.M
		if err := result.All(ctx, &allInvoices); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allInvoices)

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()
		var invoice models.Invoice
		invoiceId := c.Param("invoice_id")
		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching invoice"})
			return
		}
		var invoiceView InvoiceViewFormat
		allOrderItems, err := ItemsByOrder(invoice.Order_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching invoiceView"})
		}
		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date
		invoiceView.Payment_method = "null"
		if invoice.Payment_method != nil {
			invoiceView.Payment_method = *invoice.Payment_method
		}

		invoiceView.Invoice_id = invoice.Invocie_id
		invoice.Payment_status = *&invoice.Payment_status
		invoiceView.Payment_due = allOrderItems[0]["payment_due"]
		invoiceView.Table_number = allOrderItems[0]["table number"]
		invoiceView.Order_details = allOrderItems[0]["order_items"]

		c.JSON(http.StatusOK, invoiceView)

	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()
		var invoice models.Invoice
		var order models.Order
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(invoice)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		err := orderCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_id}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Table was not found"})
			return

		}
		staus := "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &staus
		}

		invoice.Payment_due_date, _ = time.Parse(time.RFC3339, time.Now().AddDate(0, 0, 1).Format(time.RFC3339))
		invoice.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.ID = primitive.NewObjectID()
		invoice.Order_id = invoice.ID.Hex()

		result, err := invoiceCollection.InsertOne(ctx, invoice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invoice was not created"})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()
		var invoice models.Invoice
		var updateObj primitive.D
		invoiceId := c.Param("invoice_id")
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		upsert := true
		filter := bson.M{"invoice_id": invoiceId}
		if invoice.Payment_method != nil {
			updateObj = append(updateObj, bson.E{"payment_method", invoice.Payment_method})
		}

		if invoice.Payment_status != nil {
			updateObj = append(updateObj, bson.E{"payment_status", invoice.Payment_status})
		}
		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", invoice.Updated_at})

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		staus := "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &staus
		}

		result, err := invoiceCollection.UpdateOne(
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
