package internal

import (
	"BeatsPro/internal/controllers"
	"BeatsPro/internal/operations"
	"BeatsPro/internal/response/responseFactories"
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
			operations.GetOperationLocator().GetCreateTagOperation(),
			operations.GetOperationLocator().GetUpdateTagOperation(),
			operations.GetOperationLocator().GetDeleteTagOperation(),
			operations.GetOperationLocator().GetGetByIdTagOperation(),
			responseFactories.GetResponseFactoryLocator().GetConsumerErrorResponseFactory(),
		),
		new(controllers.HealthController),
	)
}
