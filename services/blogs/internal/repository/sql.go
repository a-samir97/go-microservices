package repository

import (
	"blogs/config"
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

