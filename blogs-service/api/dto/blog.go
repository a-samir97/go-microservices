package dto

import "time"

type BlogRequest struct {
	UserId      string
	Title       string
	Description string
}

type BlogResponse struct {
	Id          string
	UserID      string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
