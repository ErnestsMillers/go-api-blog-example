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

	category, err := services.NewCategory(newCategory)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(category)
}

func GetCategory(ctx *fiber.Ctx) error {
	category, err := services.GetCategory(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(category)
}

func GetCategories(ctx *fiber.Ctx) error {
	categories, err := services.GetCategories()
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(categories)
}

func UpdateCategory(ctx *fiber.Ctx) error {
	var category *entities.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	category, err := services.UpdateCategory(category, ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(category)
}

func DeleteCategory(ctx *fiber.Ctx) error {
	Category, err := services.DeleteCategory(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Category)
}
