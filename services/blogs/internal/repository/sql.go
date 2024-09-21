package repository

import (
	"blogs/config"
	"blogs/internal/domain"
	"database/sql"
	"log"
)

type SqlRepository struct {
	// DB connection
	Db *sql.DB
}

func NewSqlRepository() *SqlRepository {
	sqlConnection := config.SQLConnection{}
	db, err := sqlConnection.Connect()
	if err != nil {
		log.Fatal("ORM Connection Error", err.Error())
		return nil
	}

	return &SqlRepository{Db: db}
}

func (bs *SqlRepository) Create(user domain.Blog) (domain.Blog, error) {

	return user, nil
}

func (bs *SqlRepository) Update(updatedUser domain.Blog, id string) (domain.Blog, error) {

	return updatedUser, nil
}

func (bs *SqlRepository) Delete(id string) error {
	return nil
}

func (bs *SqlRepository) GetByID(id string) *domain.Blog {
	return nil
}

func (bs *SqlRepository) List() []domain.Blog {
	return nil
}
