package controllers

import (
	"io"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (userController *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "register")
}

func (userController *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "login")
}
