package config

import (
	"database/sql"
	"log"
	"users/internal/domain"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection interface {
	Connect() (interface{}, error)
}

type ORMConnection struct{}

func (oc *ORMConnection) Connect() (*gorm.DB, error) {
	dsn := "host=users_db user=postgres password=postgres dbname=users port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.User{})
	return db, nil
}

type SQLConnection struct{}

func (sc *SQLConnection) Connect() (*sql.DB, error) {
	dsn := "host=users_db user=postgres password=postgres dbname=users port=5432 sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalln("SqlConnection Error", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("SqlConnection Error", err.Error())
	}

	return db, nil
}
