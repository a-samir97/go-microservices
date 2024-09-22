package repository

import (
	"database/sql"
	"log"
	"users/internal/config"
	"users/internal/domain"
)

type SqlRepository struct {
	Db *sql.DB
}

func NewSqlRepository() *SqlRepository {
	sqlConnection := config.SQLConnection{}
	db, err := sqlConnection.Connect()
	if err != nil {
		log.Fatal("Sql Connection Error", err.Error())
		return nil
	}
	return &SqlRepository{Db: db}
}
func (up *SqlRepository) Save(user *domain.User) (*domain.User, error) {
	sqlStatment, err := up.Db.Prepare(`
		INSERT INTO users (first_name, last_name, username, password, email)
		VALUES ($1, $2, $3, $4, $5, $6)
	`)
	if err != nil {
		return nil, err
	}

	_, err = sqlStatment.Query(user.FirstName, user.LastName, user.Password, user.Email)

	if err != nil {
		return nil, err
	}

	defer sqlStatment.Close()
	return user, nil
}

func (up *SqlRepository) FindByID(id string) (*domain.User, error) {

	var user domain.User

	row := up.Db.QueryRow(`
		SELECT first_name, last_name, username, email FROM users WHERE user_id = ?
	`, id)

	if err := row.Scan(user.FirstName, user.LastName, user.Username, user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (up *SqlRepository) Update(updatedUser domain.User, id string) (*domain.User, error) {
	statment, err := up.Db.Prepare(`
	UPDATE users set first_name = ?, last_name = ?, email = ? WHERE user_id = ? 
	`)

	if err != nil {
		return nil, err
	}

	_, err = statment.Exec(updatedUser.FirstName, updatedUser.LastName, updatedUser.Email)

	if err != nil {
		return nil, err
	}
	defer statment.Close()
	return &updatedUser, nil
}

func (up *SqlRepository) Delete(id string) error {
	stmt, err := up.Db.Prepare(`DELETE FROM users WHERE user_id = ?`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	defer stmt.Close()
	return nil
}
