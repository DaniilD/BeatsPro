package internal

import (
	"BeatsPro/internal/controllers"
	"BeatsPro/internal/models/Tag"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/services"
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
		controllers.NewTagController(
			requests_validators.NewValidatorFactory(),
			Tag.NewTagFactory(),
			services.GetServiceLocator().GetTagService(),
		),
	)
}
