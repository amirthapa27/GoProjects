package middleware

import (
	"strconv"

	"github.com/amirthapa27/menu-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getMenu(c *fiber.Ctx) error {
	menu, err := models.GetMenu()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not retrive the menu data" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(menu)
}

func createCategory(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var menu models.Menu
	err := c.BodyParser(&menu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not parse the json data for menu" + err.Error(),
		})
	}
	err = models.CreateCategory(menu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create a new category" + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(menu)
}

func deleteCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("category_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not parse ID for deleting category" + err.Error(),
		})
	}
	err = models.DeleteCategory(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete the category",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted category",
	})
}

func updateCategory(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var menu models.Menu
	err := c.BodyParser(&menu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not parse the json data for updating menu" + err.Error(),
		})
	}
	err = models.UpdateMenu(menu)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update",
		})
	}
	return c.Status(fiber.StatusOK).JSON(menu)
}

func getAllItems(c *fiber.Ctx) error {
	items, err := models.GetAllItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not retrive all the items" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(items)
}

func createItem(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var item models.Items
	err := c.BodyParser(&item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not parse the create data for item",
		})
	}
	err = models.CreateItem(item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create the item" + err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

func updateItem(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var item models.Items
	err := c.BodyParser(&item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not parse the update data for item" + err.Error(),
		})
	}
	err = models.UpdateItems(item)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update item" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(item)
}

func deleteItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("item_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Parsing error in deleting item" + err.Error(),
		})
	}
	err = models.DeleteItem(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not delete the item" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Item deleted",
	})

}

func getCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("menu_category_id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Parsing error in retriving category" + err.Error(),
		})
	}
	items, err := models.GetCategory(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not retrive data",
		})
	}
	return c.Status(fiber.StatusOK).JSON(items)
}

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/menu/", getMenu)
	router.Post("/menu/", createCategory)
	router.Delete("menu/:category_id", deleteCategory)
	router.Patch("/menu/", updateCategory)

	router.Get("/items/", getAllItems)
	router.Post("/items/", createItem)
	router.Delete("/items/:item_id", deleteItem)
	router.Patch("/items/", updateItem)

	router.Get("/category/:menu_category_id", getCategory)

	router.Listen(":8000")
}
