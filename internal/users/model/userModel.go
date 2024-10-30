package model

type User struct {
	ID       int    `json:"_" DB:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	// Role     string `json:"role" binding:"required"`
}
