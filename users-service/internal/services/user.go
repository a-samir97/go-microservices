package services

import (
	"users/internal/domain"
	"users/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser(user domain.User) (*domain.User, error) {

	return us.repo.Save(&user)
}

func (us *UserService) FindByID(id string) (*domain.User, error) {
	return us.repo.FindByID(id)
}

func (us *UserService) UpdateUser(updatedUser domain.User, id string) (*domain.User, error) {

	return us.repo.Update(updatedUser, id)
}

func (us *UserService) DeleteUser(id string) error {

	return us.repo.Delete(id)
}
