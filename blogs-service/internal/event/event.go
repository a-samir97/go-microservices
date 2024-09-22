package event

type Publisher interface {
	PublishEvent(blogEvent interface{}, topic string) error
}

type BlogCreatedEvent struct {
	Title       string
	Description string
	UserId      string
	CreatedAt   string
}

type BlogViewedEvent struct {
	BlogID int
}

type BlogDeletedEvent struct {
	BlogId int
}

type BlogLikedEvent struct {
	BlogId int
	UserId int
}

type BlogDislikeEvent struct {
	BlogId int
	UserId int
}

type BlogClappedEvent struct {
	BlogId int
	UserId int
}
