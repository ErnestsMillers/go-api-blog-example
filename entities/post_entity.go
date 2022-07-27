package entities

type Post struct {
	Id       int64  `json:"id"       gorm:"column:post_id; primaryKey; <-:create"`
	Name     string `json:"name"     gorm:"column:post_name"`
	Text     string `json:"text"     gorm:"column:post_text"`
	Category int32  `json:"category" gorm:"column:category_id"`
}
