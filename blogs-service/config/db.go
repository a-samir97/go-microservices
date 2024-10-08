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

	_, err = db.Exec(`
	
		CREATE TABLE IF NOT EXISTS blogs (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0 ,
			claps INTEGER DEFAULT 0,
			views INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT current_timestamp,
			updated_at TIMESTAMP DEFAULT current_timestamp
		);
	`)

	if err != nil {
		return nil, err
	}

	return db, nil
}
