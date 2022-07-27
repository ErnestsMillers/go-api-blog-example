package services

import (
	"errors"

	"go-api-blog-example/database"
	"go-api-blog-example/entities"
	"go-api-blog-example/utils"
)

func NewPost(post *entities.Post) (*entities.Post, error) {
	var err error
	tx := database.Conn.Create(&post)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}
	return post, err
}

func GetPost(postId string) (*entities.Post, error) {
	var err error
	var post *entities.Post
	tx := database.Conn.Find(&post, postId)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if post.Id == 0 {
		err = errors.New("post not found")
	}
	return post, err
}

func GetPosts() (*[]entities.Post, error) {
	var err error
	var posts *[]entities.Post
	tx := database.Conn.Order("post_id").Find(&posts)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if *posts == nil {
		err = errors.New("post not found")
	}
	return posts, err
}

func UpdatePost(newPost *entities.Post, postId string) (*entities.Post, error) {
	post, err := GetPost(postId)
	if err == nil {
		tx := database.Conn.Model(post).Updates(newPost)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return post, err
}

func DeletePost(postId string) (*entities.Post, error) {
	post, err := GetPost(postId)
	if err == nil {
		tx := database.Conn.Delete(post)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return post, err
}
