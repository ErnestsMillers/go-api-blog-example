package routes

import (
	"go-api-blog-example/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(api fiber.Router) {
	apiCategory := api.Group("/category")
	apiCategory.Post("/", controllers.NewCategory)
	apiCategory.Get("/", controllers.GetCategories)
	apiCategory.Get("/:id", controllers.GetCategory)
	apiCategory.Put("/:id", controllers.UpdateCategory)
	apiCategory.Delete("/:id", controllers.DeleteCategory)
}
