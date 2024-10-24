package model

import "time"

type Category struct {
	ID        int       `json:"id" DB:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
