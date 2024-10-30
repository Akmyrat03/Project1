package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/akmyrat/project1/internal/users/model"
	"github.com/akmyrat/project1/internal/users/repository"
	"github.com/dgrijalva/jwt-go"
)

var (
	salt       = "###1231231@#!$rf23423$!!@#342"
	signingKey = []byte("###%5645646566")
)

type UserService struct {
	repo *repository.UserRepository
}

type RoleBasedClaims struct {
	jwt.StandardClaims     //It IssuedAt , ExpiresAt informations hold
	UserId             int `json:"user_id"`
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user model.User) (int, error) {
	user.Password = generateHashedPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHashedPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &RoleBasedClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.ID,
	})
	return token.SignedString(signingKey)
}

func generateHashedPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *UserService) DeleteUser(UserId int) error {
	return s.repo.DeleteUser(UserId)
}
