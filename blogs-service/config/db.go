package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection interface {
	Connect() (interface{}, error)
}

type ORMConnection struct{}

func (oc *ORMConnection) Connect() (*gorm.DB, error) {
	dsn := "host=blogs_db user=postgres password=postgres dbname=blogs port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	return db, nil
}

type SQLConnection struct{}

func (sc *SQLConnection) Connect() (*sql.DB, error) {
	dsn := "host=blogs_db user=postgres password=postgres dbname=blogs port=5432 sslmode=disable"
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
