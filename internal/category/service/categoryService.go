package service

import (
	"github.com/akmyrat/project1/internal/category/model"
	"github.com/akmyrat/project1/internal/category/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) CreateCategory(body *model.Category) (*model.Category, error) {
	return s.repo.Create(body)
}

func (s *CategoryService) GetAllCategories() ([]*model.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetOneCategory(id int) (*model.Category, error) {
	return s.repo.GetOne(id)
}

func (s *CategoryService) DeleteCategoryByID(id int) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) UpdateCategoryByiD(id int, category *model.Category) error {
	err := s.repo.Update(id, category)
	if err != nil {
		return err
	}
	return nil
}
