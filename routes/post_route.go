package routes

import (
	"go-api-blog-example/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(api fiber.Router) {
	apiPost := api.Group("/post")
	apiPost.Post("/", controllers.NewPost)
	apiPost.Get("/", controllers.GetPosts)
	apiPost.Get("/:id", controllers.GetPost)
	apiPost.Put("/:id", controllers.UpdatePost)
	apiPost.Delete("/:id", controllers.DeletePost)
}
