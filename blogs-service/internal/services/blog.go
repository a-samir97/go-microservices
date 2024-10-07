package services

import (
	"blogs/internal/domain"
	"blogs/internal/event"
	"blogs/internal/repository"
	"log"
)

type BlogService struct {
	repository     repository.BlogRepository
	eventPublisher event.Publisher
}

func NewBlogService(repo repository.BlogRepository, eventPublisher event.Publisher) BlogService {

	return BlogService{repository: repo, eventPublisher: eventPublisher}
}

func (bs *BlogService) Create(user domain.Blog) (*domain.Blog, error) {

	return bs.repository.Create(user)
}

func (bs *BlogService) Update(updatedUser domain.Blog, id string) (*domain.Blog, error) {
	return bs.repository.Update(updatedUser, id)
}

func (bs *BlogService) Delete(id int) error {
	return bs.repository.Delete(id)
}

func (bs *BlogService) GetById(id int) (*domain.Blog, error) {
	blog, err := bs.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	bs.eventPublisher.PublishEvent(blog)
	return blog, err
}

func (bs *BlogService) List() ([]domain.Blog, error) {
	return bs.repository.List()
}

func (bs *BlogService) LikeBlog(id int) (*domain.Blog, error) {
	blog, err := bs.repository.Like(id)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	blogLikedEvent := event.BlogLikedEvent{BlogId: blog.ID, Likes: int(blog.Likes)}
	bs.eventPublisher.PublishBlogLiked(blogLikedEvent)
	return blog, err
}

func (bs *BlogService) DislikeBlog(id int) (*domain.Blog, error) {
	blog, err := bs.repository.Dislike(id)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	blogDislikedEvent := event.BlogDislikeEvent{BlogId: blog.ID, Dislikes: int(blog.Dislikes)}
	bs.eventPublisher.PublishBlogDisliked(blogDislikedEvent)
	return blog, err
	// publish event to kafka (BlogDisliked)
}

func (bs *BlogService) ClapBlog(id int) (*domain.Blog, error) {
	blog, err := bs.repository.Clap(id)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	blogClappedEvent := event.BlogClappedEvent{BlogId: blog.ID, Claps: int(blog.Claps)}
	bs.eventPublisher.PublishBlogClaped(blogClappedEvent)
	return blog, err
}
