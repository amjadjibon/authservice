package repository

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

type IRepository interface {
	InsertUser(user User) error
	GetUserByUsername(username string) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserById(id int64) (User, error)
	DeleteUserById(id int64) error
}

func FactoryRepository(store, dsn string) IRepository {
	switch store {
	case "postgres":
		return NewPostgresRepository(dsn)
	default:
		panic("unknown store")
	}
}
