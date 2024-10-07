package domain

import "time"

type Blog struct {
	ID          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Claps       int32     `json:"claps"`    // like medium
	Likes       int32     `json:"likes"`    // TODO: handle this in the future (separated model)
	Dislikes    int32     `json:"dislikes"` // TODO: handle this in the future (separated model)
	Views       int32     `json:"views"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
