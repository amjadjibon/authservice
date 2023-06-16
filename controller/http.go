package controller

import (
	"fmt"
	"net/http"
	"time"

	"authservice/model"
	"authservice/repository"
	"authservice/utils"
)

type AuthServiceHandler struct {
	repo repository.IRepository
}

func NewAuthServiceHandler(repo repository.IRepository) *AuthServiceHandler {
	return &AuthServiceHandler{repo: repo}
}

func (a AuthServiceHandler) Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	serializer := RegistrationSerializer{}
	if err := Serialize(r, &serializer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &model.User{
		Username:  serializer.Username,
		Password:  serializer.Password,
		Email:     serializer.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := a.repo.InsertUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	verification := &model.Verification{
		UserId: user.Id,
		Code:   utils.RandomString(6),
	}

	err = a.repo.InsertVerification(verification)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	subject := "Verification Code"
	body := fmt.Sprintf("Your verification code is: %s", verification.Code)

	err = utils.SendEmail(user.Email, subject, body)
	if err != nil {
		panic(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a AuthServiceHandler) Verification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	serializer := VerificationSerializer{}
	if err := Serialize(r, &serializer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO implement me
	panic("implement me")
}

func (a AuthServiceHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	serializer := LoginSerializer{}
	if err := Serialize(r, &serializer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(serializer.Username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(serializer.Password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := a.repo.GetUserByUsername(serializer.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if user.Password != serializer.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
