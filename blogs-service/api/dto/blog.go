package dto

import "time"

type BlogRequest struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BlogResponse struct {
	Id          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Likes       int       `json:"likes"`
	Dislikes    int       `json:"dislikes"`
	Claps       int       `json:"claps"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
