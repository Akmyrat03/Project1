package model

type Post struct {
	ID           int     `json:"id" db:"id"`
	CategoryId   int     `json:"category_id" form:"category_id" db:"category_id" binding:"required"`
	UserId       int     `json:"user_id" form:"user_id" db:"user_id" binding:"required"`
	CategoryName string  `db:"category_name"`
	UserName     string  `db:"user_name"`
	Title        string  `json:"title" form:"title" binding:"required"`
	Description  string  `json:"description" form:"description" binding:"required"`
	ImagePath    *string `json:"image_path" form:"image_path" db:"image_path"`
}
