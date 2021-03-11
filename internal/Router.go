package internal

import (
	"BeatsPro/internal/controllers"
	"github.com/gorilla/mux"
)

type Router struct {
	muxRouter      *mux.Router
	userController *controllers.UserController
	tagController  *controllers.TagController
}

func NewRouter(
	router *mux.Router,
	userController *controllers.UserController,
	tagController *controllers.TagController) *Router {

	return &Router{
		muxRouter:      router,
		userController: userController,
		tagController:  tagController,
	}
}

func (router *Router) InitRouts() {
	router.muxRouter.HandleFunc("/user/register", router.userController.RegisterUser).Methods("POST")
	router.muxRouter.HandleFunc("/user/login", router.userController.LoginUser).Methods("POST")

	router.muxRouter.HandleFunc("/tag/create", router.tagController.CreateTag).Methods("POST")

}