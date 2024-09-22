package repository

import (
	"users/internal/domain"
)

type UserRepository interface {
	Save(*domain.User) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	Update(user domain.User, id string) (*domain.User, error)
	Delete(id string) error
}
