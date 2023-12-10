package features

import (
	"strconv"

	"github.com/elanticrypt0/gasonline/api/models"
	"github.com/elanticrypt0/gasonline/pkg/webcore"
	"github.com/gofiber/fiber/v2"
)

func FindOneCategory(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	cat := models.NewCategory()
	return c.JSON(cat.FindOne(gas.App.DB.Primary, id))
}

func FindAllCategories(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	cat := models.NewCategory()
	categories := cat.FindAll(gas.App.DB.Primary)
	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	// name comes from json in body
	// cat := new(models.Category)
	cat := models.NewCategory()
	c.BodyParser(&cat)
	category := cat.Create(gas.App.DB.Primary, cat.Name)
	return c.JSON(category)
}

func UpdateCategory(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	cat := models.NewCategory()
	category := cat.FindOne(gas.App.DB.Primary, id)
	c.BodyParser(&cat)
	category = cat.Update(gas.App.DB.Primary, *category)
	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	id, _ := strconv.Atoi(c.Params("id", "0"))
	cat := models.NewCategory()
	category := cat.Delete(gas.App.DB.Primary, id)
	return c.JSON(category)
}
