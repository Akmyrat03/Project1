package repository

import (
	"fmt"

	"github.com/akmyrat/project1/internal/users/model"
	"github.com/jmoiron/sqlx"
)

const (
	Users = "users"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(DB *sqlx.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %v (name, username, password) VALUES ($1, $2, $3) RETURNING id`, Users)
	rows := r.DB.QueryRow(query, user.Name, user.Username, user.Password)
	if err := rows.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}

func (r *UserRepository) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf(`SELECT * FROM %v WHERE username = $1 and password = $2`, Users)
	err := r.DB.Get(&user, query, username, password)
	if err != nil {
		return model.User{}, err
		fmt.Errorf("error getting user: %v", err)
	}
	return user, nil
}
