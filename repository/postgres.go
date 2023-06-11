package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func getDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func NewPostgresRepository(dsn string) IRepository {
	return &PostgresRepository{
		db: getDB(dsn),
	}
}

func (p PostgresRepository) InsertUser(user User) error {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetUserByUsername(username string) (User, error) {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetUserByEmail(email string) (User, error) {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetUserById(id int64) (User, error) {
	// TODO implement me
	panic("implement me")
}

func (p PostgresRepository) DeleteUserById(id int64) error {
	// TODO implement me
	panic("implement me")
}
