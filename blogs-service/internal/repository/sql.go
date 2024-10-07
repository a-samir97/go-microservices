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

func (bs *SqlRepository) Create(blog domain.Blog) (*domain.Blog, error) {

	stmt, err := bs.Db.Prepare(`
		INSERT INTO blogs (title, description, user_id)
		VALUES($1, $2, $3)
	`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(blog.Title, blog.Description, blog.UserId)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return &blog, nil
}

func (bs *SqlRepository) Update(updatedBlog domain.Blog, id string) (*domain.Blog, error) {
	stmt, err := bs.Db.Prepare(`
		UPDATE blogs set title=$1, description=$2
	`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(updatedBlog.Title, updatedBlog.Description)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return &updatedBlog, nil
}

func (bs *SqlRepository) Delete(id int) error {
	stmt, err := bs.Db.Prepare(`DELETE FROM blogs WHERE id = $1`)

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

func (bs *SqlRepository) GetByID(id int) (*domain.Blog, error) {
	var blog domain.Blog

	row := bs.Db.QueryRow(`SELECT id, title, description, user_id, likes, dislikes, claps FROM blogs WHERE id = $1`, id)

	if err := row.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.UserId, &blog.Likes, &blog.Dislikes, &blog.Claps); err != nil {
		return nil, err
	}
	return &blog, nil
}

func (bs *SqlRepository) List() ([]domain.Blog, error) {
	rows, err := bs.Db.Query(`SELECT id, title, description, user_id, likes, dislikes, claps FROM blogs`)

	if err != nil {
		return nil, err
	}
	var blogs []domain.Blog
	for rows.Next() {
		var blog domain.Blog
		rows.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.UserId, &blog.Likes, &blog.Dislikes, &blog.Claps)
		blogs = append(blogs, blog)
	}
	rows.Close()
	return blogs, nil
}

func (bs *SqlRepository) Like(id int) (*domain.Blog, error) {
	stmt, err := bs.Db.Prepare(`UPDATE blogs set likes = likes + 1 WHERE id = $1`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return bs.GetByID(id)
}

func (bs *SqlRepository) Dislike(id int) (*domain.Blog, error) {
	stmt, err := bs.Db.Prepare(`UPDATE blogs set dislikes = dislikes + 1 WHERE id = $1`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return bs.GetByID(id)
}

func (bs *SqlRepository) Clap(id int) (*domain.Blog, error) {

	stmt, err := bs.Db.Prepare(`UPDATE blogs set claps = claps + 1 WHERE id = $1`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return bs.GetByID(id)
}
