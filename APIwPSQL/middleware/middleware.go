package middleware

import (
	"apipsql/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getAllStocks(c *fiber.Ctx) error {
	stocks, err := models.GetAllStocks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error in getting all the links" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(stocks)
}

func getOneStock(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("ID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error while parsing" + err.Error(),
		})
	}
	stock, err := models.GetOneStock(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not retrive the stock" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(stock)

}

func createStock(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var stock models.Stocks
	err := c.BodyParser(&stock) //parse the jsn data into stock
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "parsing error" + err.Error(),
		})
	}
	err = models.CreateStock(stock)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create stock" + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(stock)
}

func deleteStock(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("ID"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "parsing error" + err.Error(),
		})
	}
	err = models.DeleteStock(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete stock" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted stock",
	})
}

func updateStock(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var stock models.Stocks
	err := c.BodyParser(&stock)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "parsing error" + err.Error(),
		})
	}
	err = models.UpdateStocks(stock)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update stocks" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(stock)
}

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.Get("/stocks/", getAllStocks)
	router.Get("/stock/:id", getOneStock)
	router.Post("/stock/", createStock)
	router.Delete("/stock/:id", deleteStock)
	router.Patch("/stock/", updateStock)

	router.Listen(":3000")
}
