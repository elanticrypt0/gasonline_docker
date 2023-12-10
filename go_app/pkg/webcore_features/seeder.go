package webcore_features

import (
	"log"

	"github.com/elanticrypt0/gasonline/api/models"
	"github.com/elanticrypt0/gasonline/pkg/webcore"
	"github.com/elanticrypt0/go4it"
	"github.com/gofiber/fiber/v2"
)

const seedDir = "./seeds/"

func Seed(c *fiber.Ctx, gas *webcore.GasonlineApp) error {
	seedCategories(gas)
	return c.JSON("OK")
}

func seedCategories(gas *webcore.GasonlineApp) {
	cat_list := []models.Category{}
	go4it.ReadAndParseJson(seedDir+"categories", &cat_list)
	// for _, category := range cat_list {
	// 	models.CreateCategory(gas, category.Name)
	// }
	gas.App.DB.Primary.Save(&cat_list)
	log.Println("Categories seeded")
}
