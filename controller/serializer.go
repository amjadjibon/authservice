package controller

import (
	"encoding/json"
	"net/http"
)

type RegistrationSerializer struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type VerificationSerializer struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type LoginSerializer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Serialize(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
