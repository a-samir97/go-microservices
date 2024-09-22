package services

import (
	"blogs/internal/domain"
	"blogs/internal/repository"
)

type BlogService struct {
	repository repository.BlogRepository
}

func NewBlogService(repo repository.BlogRepository) BlogService {

	return BlogService{repository: repo}
}

func (bs *BlogService) Create(user domain.Blog) (domain.Blog, error) {

	return bs.repository.Create(user)
}

func (bs *BlogService) Update(updatedUser domain.Blog, id string) (domain.Blog, error) {
	return bs.repository.Update(updatedUser, id)
}

func (bs *BlogService) Delete(id string) error {
	return bs.repository.Delete(id)
}

func (bs *BlogService) GetById(id string) *domain.Blog {
	return bs.repository.GetByID(id)
}

func (bs *BlogService) List() []domain.Blog {
	return bs.repository.List()
}
