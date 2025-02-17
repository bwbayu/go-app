package service

import (
	"go-app/internal/models"
	"go-app/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.Delete(id)
}
