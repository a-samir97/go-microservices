package services

import (
	"blogs/internal/domain"
	"blogs/internal/event"
	"blogs/internal/repository"
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

	bs.eventPublisher.PublishEvent(blog, "blog_viewed")
	return blog, err
}

func (bs *BlogService) List() ([]domain.Blog, error) {
	return bs.repository.List()
}
