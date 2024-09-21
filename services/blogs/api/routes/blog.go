package routes

import "blogs/api/handlers"

type BlogRoutes struct {
	handler handlers.BlogHandler
}

func NewBlogRoutes(handler handlers.BlogHandler) *BlogRoutes {
	return &BlogRoutes{handler: handler}
}
