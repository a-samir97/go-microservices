package repository

import "users/internal/domain"

type UserRepository interface {
	Save(*domain.User) error
	FindByID(id string) (*domain.User, error)
	Update(id string) (*domain.User, error)
	Delete(id string) error
}

// implement the interface implicitly
type UserRepo struct {
	// DB Connection
}

func (up *UserRepo) Save(user domain.User) error {
	// insert user to the database
	return nil
}

func (up *UserRepo) FindByID(id string) (*domain.User, error) {

	return nil, nil
}

func (up *UserRepo) Update(id string) (*domain.User, error) {

	return nil, nil
}

func (up *UserRepo) Delete(id string) error {

	return nil
}
