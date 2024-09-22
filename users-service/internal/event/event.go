package event

import "time"

type Publisher interface {
	PublishUserCreatedEvent(UserCreatedEvent)
	PublishUserUpdatedEvent()
	PublishUserDeletedEvent()
}

type Consumer interface {
}

type UserCreatedEvent struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	CreatedAt time.Time
}

type UserUpdatedEvent struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	UpdatedAt time.Time
}

type UserDeletedEvent struct {
	ID string
}
