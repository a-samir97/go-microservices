package repository

import (
	"log"
	"users/internal/config"
	"users/internal/domain"

	"gorm.io/gorm"
)

type ORMRepository struct {
	Db gorm.DB
}

func NewORMRepository() *ORMRepository {
	ormConnection := config.ORMConnection{}
	db, err := ormConnection.Connect()
	if err != nil {
		log.Fatal("ORM Connection Error", err.Error())
		return nil
	}

	return &ORMRepository{Db: *db}
}

func (up *ORMRepository) Save(user *domain.User) (*domain.User, error) {
	result := up.Db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (up *ORMRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	result := up.Db.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (up *ORMRepository) Update(updatedUser domain.User, id string) (*domain.User, error) {
	result := up.Db.Save(domain.User{
		ID: id, FirstName: updatedUser.FirstName,
		LastName: updatedUser.LastName,
		Email:    updatedUser.Email,
		Avatar:   updatedUser.Avatar,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &updatedUser, nil
}

func (up *ORMRepository) Delete(id string) error {
	result := up.Db.Delete(domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
