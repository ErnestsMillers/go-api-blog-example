package entities

type Category struct {
	Id   int64  `json:"id"       gorm:"column:category_id; primaryKey; <-:create"`
	Name string `json:"name"     gorm:"column:category_name"`
}
