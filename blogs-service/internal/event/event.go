package event

type Publisher interface {
	PublishBlogCreatedEvent(BlogCreatedEvent)
	PublishBlogViewedEvent()
	PublishBlogDeletedEvent()
	PublishBlogLikeEvent()
	PublishBlogDislikeEvent()
	PublishBlogClappedEvent()
}

type BlogCreatedEvent struct {
	Title       string
	Description string
	UserId      string
	CreatedAt   string
}

type BlogViewedEvent struct {
	BlogID string
	UserID string
}

type BlogDeletedEvent struct {
	BlogId string
}

type BlogLikedEvent struct {
	BlogId string
	UserId string
}

type BlogDislikeEvent struct {
	BlogId string
	UserId string
}

type BlogClappedEvent struct {
	BlogId string
	UserId string
}
