package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"github.com/amirthapa27/golangJWT/database"
	"github.com/amirthapa27/golangJWT/helpers"
	"github.com/amirthapa27/golangJWT/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollectiion(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedPassword)
}

func VerifyPassword(userPassword, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		check = false
		msg = "email or password is incorrect"
	}
	return check, msg
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		validationError := validate.Struct(user)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": validationError.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occured while checking for the email"})
		}
		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occured while checking for the phonr"})
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Email or Phone number Already exists"})
			return
		}
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID() //will generate a new object id
		user.User_id = user.ID.Hex()      //will return object id as string

		token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_Name, *user.Last_Name, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		//inserting the new user in the database
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInsertionNumber)

	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		}
		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name, *foundUser.User_type, *&foundUser.User_id)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, foundUser)

	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.CheckUserType(c, "ADMIN") //check the user type
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var recordPerPage, err1 = strconv.Atoi(c.Query("recordPerPage"))
		if err1 != nil || recordPerPage < 1 {
			recordPerPage = 10 //10 records on the first page by default
		}
		var page, err2 = strconv.Atoi(c.Query("page"))
		if err2 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{ //defines the datapoint to go the user
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}},
		}
		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": " error occured while listing user items"})
		}
		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil { //retrives all the documents in allUsers
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allUsers[0])

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		if err := helpers.MatchUserTypeToUid(c, userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user) //give us the user with the id specified
	}

}
