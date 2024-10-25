package repository

import (
	"errors"
	"fmt"

	"github.com/akmyrat/project1/internal/category/model"
	"github.com/jmoiron/sqlx"
)

const (
	Categories = "categories"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(DB *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: DB,
	}
}

func (r *CategoryRepository) Create(body *model.Category) (*model.Category, error) {
	if body.Name == "" {
		return nil, errors.New("name is required")
	}

	category := &model.Category{
		Name: body.Name,
		// CreatedAt: body.CreatedAt,
		// UpdatedAt: body.UpdatedAt,
	}

	query := fmt.Sprintf(`INSERT INTO %v (name) VALUES ($1) RETURNING id`, Categories)

	row := r.DB.QueryRow(query, category.Name)
	var insertedID int
	if err := row.Scan(&insertedID); err != nil {
		return nil, err
	}

	return &model.Category{
		ID:   insertedID,
		Name: body.Name,
	}, nil
}

func (r *CategoryRepository) GetAll() ([]*model.Category, error) {
	// SQL sorgusunu hazirliyoruz
	query := fmt.Sprintf(`SELECT id, name, created_at, updated_at FROM %v`, Categories)

	// Sorguyu veritabanina gonderiyoruz
	rows, err := r.DB.Query(query)

	// Hata olustuysa, hata ile fonsksiyondan cikiyoruz
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Kategorileri saklamak icin bir alan hazirliyoruz
	var categories []*model.Category

	// Her bir satir icin donguye giriyoruz
	for rows.Next() {
		//Yeni bir 'category' degiskeni tanimliyoruz
		var category model.Category

		// rows.Scan ile her satiri 'category' degiskenine aktariyoruz
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		// Hatayi dondurerek fonksiyondan cikiyoruz
		if err != nil {
			return nil, err
		}
		// 'categori ekliyoruz'
		categories = append(categories, &category)
	}

	// Satirlari islerken olusabilecek hatalari kontrol ediyoruz
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// Eger hata yoksa, katogiriler dilimini donduruyoruz
	return categories, nil
}

func (r *CategoryRepository) GetOne(id int) (*model.Category, error) {
	// SQL sorgusunu hazirliyoruz
	query := fmt.Sprintf(`SELECT id, name, created_at, updated_at FROM %v WHERE id = $1`, Categories)

	// Sorguyu veritabanina gonderiyoruz
	row := r.DB.QueryRow(query, id)

	//Kategorini saklamak icin yer sakliyoruz
	var category model.Category

	// row.Scan ile her satiri 'categori degiskenine aktariyoruz'
	err := row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	// Hata olustuysa kategoriyi nil yap ve hatayi goster
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Delete(id int) error {
	// SQL sorguyu hazirliyoruz
	query := fmt.Sprintf(`DELETE FROM %v WHERE id = $1`, Categories)

	//Sorguyu veritabanina gonderiyoruz
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) Update(id int, body *model.Category) error {
	if body.Name == "" {
		return errors.New("name is required")
	}

	//SQL sorguyu hazirliyoruz
	query := fmt.Sprintf(`UPDATE %v SET name = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`, Categories)

	//Sorguyu veritabanina gonderiyoruz
	_, err := r.DB.Exec(query, body.Name, id)
	if err != nil {
		return err
	}

	return nil
}
