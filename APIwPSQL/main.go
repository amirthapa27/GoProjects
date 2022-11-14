package main

import (
	"apipsql/middleware"
	"apipsql/models"
)

func main() {
	models.Setup()
	middleware.SetupAndListen()

}
