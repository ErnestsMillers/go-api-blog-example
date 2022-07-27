package controllers

import (
	"go-api-blog-example/entities"
	"go-api-blog-example/services"

	"github.com/gofiber/fiber/v2"
)

func NewPost(ctx *fiber.Ctx) error {
	var newPost *entities.Post
	if err := ctx.BodyParser(&newPost); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// If category doesn't exist can't create it
	_, err := services.GetCategory(ctx.Params("category"))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	Post, err := services.NewPost(newPost)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(Post)
}

func GetPost(ctx *fiber.Ctx) error {
	Post, err := services.GetPost(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Post)
}

func GetPosts(ctx *fiber.Ctx) error {
	Posts, err := services.GetPosts()
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Posts)
}

func UpdatePost(ctx *fiber.Ctx) error {
	var Post *entities.Post
	if err := ctx.BodyParser(&Post); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// If category doesn't exist can't update it
	// _, err := services.GetCategory(ctx.Params("category"))
	// if err != nil {
	// return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	// }

	Post, err := services.UpdatePost(Post, ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(Post)
}

func DeletePost(ctx *fiber.Ctx) error {
	Post, err := services.DeletePost(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Post)
}
