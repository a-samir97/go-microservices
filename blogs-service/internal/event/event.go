package event

import "time"

type Publisher interface {
	PublishBlogViewed(blogEvent BlogViewedEvent) 
	PublishBlogCreated(blogCreated BlogCreatedEvent)
	PublishBlogLiked(blogliked BlogLikedEvent)
	PublishBlogDisliked(blogDisliked BlogDislikeEvent)
	PublishBlogClaped(blogClapped BlogClappedEvent)
}

type BlogCreatedEvent struct {
	Title       string
	Description string
	UserId      int
	CreatedAt   time.Time
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
