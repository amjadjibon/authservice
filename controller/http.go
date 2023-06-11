package controller

import (
	"net/http"

	"authservice/repository"
)

type AuthServiceHandler struct {
	repo repository.IRepository
}

func NewAuthServiceHandler(repo repository.IRepository) *AuthServiceHandler {
	return &AuthServiceHandler{repo: repo}
}

func (a AuthServiceHandler) Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (a AuthServiceHandler) Verification(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (a AuthServiceHandler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}
