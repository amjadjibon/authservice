package repository

import (
	"database/sql"

	_ "github.com/lib/pq"

	"authservice/database"
	"authservice/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(dsn string) IRepository {
	return &PostgresRepository{
		db: database.GetDB(dsn),
	}
}

func (p PostgresRepository) InsertUser(user *model.User) error {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetUserByUsername(username string) (*model.User, error) {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetUserByEmail(email string) (*model.User, error) {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetUserById(id int64) (*model.User, error) {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) DeleteUserById(id int64) error {
	// TODO implement me
	panic("implement me")
}
