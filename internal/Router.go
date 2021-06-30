package internal

import (
	"BeatsPro/internal/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	muxRouter        *mux.Router
	userController   *controllers.UserController
	tagController    *controllers.TagController
	healthController *controllers.HealthController
}

func NewRouter(
	router *mux.Router,
	userController *controllers.UserController,
	tagController *controllers.TagController,
	healthController *controllers.HealthController) *Router {

	return &Router{
		muxRouter:        router,
		userController:   userController,
		tagController:    tagController,
		healthController: healthController,
	}
}

func (router *Router) InitRouts() {
	router.muxRouter.HandleFunc("/user/register", router.userController.RegisterUser).Methods("POST")
	router.muxRouter.HandleFunc("/user/login", router.userController.LoginUser).Methods("POST")

	router.muxRouter.HandleFunc("/tag/create", router.tagController.CreateTag).Methods(http.MethodPost)
	router.muxRouter.HandleFunc("/tag/update/{id}", router.tagController.UpdateTag).Methods(http.MethodPut)
	router.muxRouter.HandleFunc("/tag/delete/{id}", router.tagController.DeleteTag).Methods(http.MethodDelete)
	router.muxRouter.HandleFunc("/tag/{id}", router.tagController.GetTagById).Methods(http.MethodGet)

	router.muxRouter.HandleFunc("/healthz/ping", router.healthController.Ping).Methods(http.MethodGet)
}
