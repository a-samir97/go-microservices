package repository

import "blogs/internal/domain"

type BlogRepository interface {
	Create(domain.Blog) (domain.Blog, error)
	Update(domain.Blog, string) (domain.Blog, error)
	Delete(string) error
	GetByID(string) *domain.Blog
	List() []domain.Blog
}
