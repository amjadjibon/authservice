package model

import (
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Verification struct {
	Id     int64
	UserId int64
	Code   string
}
