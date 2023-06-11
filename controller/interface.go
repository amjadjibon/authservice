package controller

import (
	"net/http"
)

type IAuthServiceHandler interface {
	Registration(w http.ResponseWriter, r *http.Request)
	Verification(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
