package repository

import (
	"context"
	"fmt"

	"github.com/akmyrat/project1/internal/users/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	Users = "users"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(DB *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %v (name, username, password) VALUES ($1, $2, $3) RETURNING id`, Users)
	err := r.DB.QueryRow(context.Background(), query, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, name, username, password FROM %v WHERE username = $1 AND password = $2", Users)
	err := r.DB.QueryRow(context.Background(), query, username, password).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(UserId int) error {
	query := fmt.Sprintf(`DELETE FROM %v WHERE id = $1`, Users)
	_, err := r.DB.Exec(context.Background(), query, UserId)
	return err
}
