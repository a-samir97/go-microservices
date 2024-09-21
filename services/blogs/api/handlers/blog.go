package handlers

import "blogs/internal/services"

type BlogHandler struct {
	service *services.BlogService
}

func NewBlogHandler(service services.BlogService) *BlogHandler {
	return &BlogHandler{service: &service}
}
