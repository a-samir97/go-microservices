package repository

import (
	"users/internal/domain"
)

type UserRepository interface {
	Save(*domain.User) (*domain.User, error)
	FindByID(id int) (*domain.User, error)
	Update(user domain.User, id int) (*domain.User, error)
	Delete(id string) error
	FindByEmail(email string) (*domain.User, error)
}
