package repository

import (
	"errors"

	"github.com/akmyrat/project1/internal/post/model"
	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	DB *sqlx.DB
}

const (
	Posts = "posts"
)

func NewRepository(DB *sqlx.DB) *PostRepository {
	return &PostRepository{
		DB: DB,
	}
}

func (r *PostRepository) Create(post *model.Post) (*model.Post, error) {
	if post.Title == "" {
		return nil, errors.New("Title is required")
	}

	query := `INSERT INTO %s (title, description, category_id, user_id, image_path) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	row := r.DB.QueryRow(query, post.Title, post.Description, post.CategoryId, post.UserId, post.ImagePath)
	var InsertedId int
	if err := row.Scan(&InsertedId); err != nil {
		return nil, err
	}
	post.ID = InsertedId
	return post, nil
}
