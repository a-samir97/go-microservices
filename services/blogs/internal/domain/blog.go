package domain

import "time"

type Blog struct {
	ID          string
	User_ID     string
	Title       string
	Description string
	Claps       int32 // like medium 
	Likes       int32 // TODO: handle this in the future (separated model)
	Dislikes    int32 // TODO: handle this in the future (separated model)
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
