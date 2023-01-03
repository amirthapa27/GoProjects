package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/amirthapa27/JWTginGorm/intializers"
	"github.com/amirthapa27/JWTginGorm/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//Get the email/pass of the req body
	var body struct {
		Email    string
		Password string
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the body"})
		return
	}
	//Hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash the password"})
		return
	}
	//Create the user
	user := models.User{Email: body.Email, Password: string(hashPassword)}
	result := intializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	//Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	//get the email and pass form the requested body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the body"})
		return
	}
	//Look up requested user
	var user models.User
	intializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}
	//Compare sent in pass with saved user hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password or email"})
		return
	}
	//Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	//Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": user})

}
