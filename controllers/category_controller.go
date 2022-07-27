package controllers

import (
	"go-api-blog-example/entities"
	"go-api-blog-example/services"

	"github.com/gofiber/fiber/v2"
)

func NewCategory(ctx *fiber.Ctx) error {
	var newCategory *entities.Category
	if err := ctx.BodyParser(&newCategory); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	Category, err := services.NewCategory(newCategory)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(Category)
}

func GetCategory(ctx *fiber.Ctx) error {
	Category, err := services.GetCategory(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Category)
}

func GetCategories(ctx *fiber.Ctx) error {
	Categories, err := services.GetCategories()
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Categories)
}

func UpdateCategory(ctx *fiber.Ctx) error {
	var Category *entities.Category
	if err := ctx.BodyParser(&Category); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	Category, err := services.UpdateCategory(Category, ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(Category)
}

func DeleteCategory(ctx *fiber.Ctx) error {
	Category, err := services.DeleteCategory(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Category)
}
