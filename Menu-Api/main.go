package main

import (
	"github.com/amirthapa27/menu-api/middleware"
	"github.com/amirthapa27/menu-api/models"
)

func main() {
	models.Setup()
	middleware.SetupAndListen()
}
