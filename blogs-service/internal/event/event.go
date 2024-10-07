package event

type Publisher interface {
	PublishEvent(blogEvent interface{}) error
	PublishBlogLiked(blogliked BlogLikedEvent)
	PublishBlogDisliked(blogDisliked BlogDislikeEvent)
	PublishBlogClaped(blogClapped BlogClappedEvent)
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
	Likes  int
}

type BlogDislikeEvent struct {
	BlogId   int
	Dislikes int
}

type BlogClappedEvent struct {
	BlogId int
	Claps  int
}
