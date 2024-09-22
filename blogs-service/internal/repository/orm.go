package repository

import (
	"blogs/config"
	"log"

	"gorm.io/gorm"
)

type ORMRepository struct {
	// DB connection
	Db *gorm.DB
}

func NewORMRepository() *ORMRepository {
	sqlConnection := config.ORMConnection{}
	db, err := sqlConnection.Connect()
	if err != nil {
		log.Fatal("ORM Connection Error", err.Error())
		return nil
	}

	return &ORMRepository{Db: db}
}
