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

func (r *PostRepository) Delete(id int) error {
	// SQL sorguyu hazirliyoruz
	query := fmt.Sprintf(`DELETE FROM %v WHERE id =$1`, Posts)

	// Sorguyu veritabanina iletiyoruz
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) GetAll() ([]*model.Post, error) {
	// SQL sorguyu hazirliyoruz
	query := `SELECT 
				p.id,
				p.category_id,
				p.user_id,
				p.title, 
				p.description, 
				p.image_path, 
				u.name AS user_name, 
				c.name AS category_name
				FROM posts AS p
				INNER JOIN categories AS c ON p.category_id = c.id
				INNER JOIN users AS u ON p.user_id = u.id 
				ORDER BY p.id ASC`

	var posts []*model.Post
	// Select ile sorguyu 'posts' degiskenine bagliyoruz
	err := r.DB.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) GetOne(id int) (*model.Post, error) {
	// SQL sorguyu hazirliyoruz
	query := `SELECT 
				p.id,
				p.category_id,
				p.user_id,
				p.title,
				p.description,
				p.image_path
				FROM posts AS p
				WHERE p.id= $1`

	var post model.Post
	err := r.DB.Get(&post, query, id)
	if err != nil {
		return nil, err
	}
	return &post, nil

}

func (r *PostRepository) Update(id int, body *model.Post) error {
	// SQL sorguyu hazirliyoruz
	query := fmt.Sprintf(`UPDATE %v SET title = $1, description = $2, user_id = $3, category_id = $4 WHERE id = $5`, Posts)

	_, err := r.DB.Exec(query, body.Title, body.Description, body.UserId, body.CategoryId, id)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil

}
