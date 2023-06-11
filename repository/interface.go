package repository

import (
	"authservice/model"
)

type IRepository interface {
	InsertUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(id int64) (*model.User, error)
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
