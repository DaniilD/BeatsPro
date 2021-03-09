package internal

import (
	"BeatsPro/internal/controllers"
	"github.com/gorilla/mux"
)

type RouterBuilder struct {
}

func NewRouterBuilder() *RouterBuilder {
	return &RouterBuilder{}
}

func (routerFactory *RouterBuilder) Build() *Router {
	return NewRouter(
		mux.NewRouter(),
		controllers.NewUserController(),
		controllers.NewTagController(),
	)
}
