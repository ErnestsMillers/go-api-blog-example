package services

import (
	"errors"

	"go-api-blog-example/database"
	"go-api-blog-example/entities"
	"go-api-blog-example/utils"
)

func NewCategory(category *entities.Category) (*entities.Category, error) {
	var err error
	tx := database.Conn.Create(&category)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}
	return category, err
}

func GetCategory(categoryId string) (*entities.Category, error) {
	var err error
	var category *entities.Category
	tx := database.Conn.Find(&category, categoryId)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if category.Id == 0 {
		err = errors.New("category not found")
	}
	return category, err
}

func GetCategories() (*[]entities.Category, error) {
	var err error
	var categories *[]entities.Category
	tx := database.Conn.Order("category_id").Find(&categories)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if *categories == nil {
		err = errors.New("categories not found")
	}
	return categories, err
}

func UpdateCategory(newCategory *entities.Category, categoryId string) (*entities.Category, error) {
	category, err := GetCategory(categoryId)
	if err == nil {
		tx := database.Conn.Model(category).Updates(newCategory)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return category, err
}

func DeleteCategory(categoryId string) (*entities.Category, error) {
	category, err := GetCategory(categoryId)
	if err == nil {
		tx := database.Conn.Delete(category)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return category, err
}
