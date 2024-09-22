package domain

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
