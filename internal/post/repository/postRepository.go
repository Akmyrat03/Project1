package repository

import (
	"errors"
	"fmt"

	"github.com/akmyrat/project1/internal/post/model"
	"github.com/jmoiron/sqlx"
)

const (
	Posts = "posts"
)

type PostRepository struct {
	DB *sqlx.DB
}

func NewPostRepository(DB *sqlx.DB) *PostRepository {
	return &PostRepository{
		DB: DB,
	}
}

func (r *PostRepository) Create(body *model.Post) (*model.Post, error) {
	// title bos ise hata dondurur
	if body.Title == "" {
		return nil, errors.New("title is required")
	}

	// SQL sorgusunu hazirlama
	query := fmt.Sprintf(`INSERT INTO %v (title, description, category_id, user_id, image_path) 
				VALUES ($1, $2, $3, $4, $5) 
					RETURNING id`, Posts)

	// Sorguyu calistirip veriyi veritabanina ekliyoruz
	row := r.DB.QueryRow(query, body.Title, body.Description, body.CategoryId, body.UserId, body.ImagePath)
	var InsertedId int

	// Sorgu sirasinda dondirilen id degeri InsertedId degerine aktariliyor
	if err := row.Scan(&InsertedId); err != nil {
		return nil, err
	}

	body.ID = InsertedId
	return body, nil

}
