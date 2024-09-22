package repository

import "blogs/internal/domain"

type BlogRepository interface {
	Create(domain.Blog) (*domain.Blog, error)
	Update(domain.Blog, string) (*domain.Blog, error)
	Delete(int) error
	GetByID(int) (*domain.Blog, error)
	List() ([]domain.Blog, error)
}
