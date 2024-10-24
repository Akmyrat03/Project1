package service

import (
	"github.com/akmyrat/project1/internal/post/model"
	"github.com/akmyrat/project1/internal/post/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) CreatePost(post *model.Post) (*model.Post, error) {
	return s.repo.Create(post)
}
