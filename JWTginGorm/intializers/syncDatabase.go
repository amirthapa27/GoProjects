package intializers

import "github.com/amirthapa27/JWTginGorm/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
