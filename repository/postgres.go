package repository

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"authservice/model"
)

type PostgresRepository struct {
	db *gorm.DB
}

func GetDB(dsn string) *gorm.DB {
	config := &gorm.Config{}
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{}, &Verification{})
	if err != nil {
		panic(err)
	}

	return db
}

func NewPostgresRepository(dsn string) IRepository {
	return &PostgresRepository{
		db: GetDB(dsn),
	}
}

type User struct {
	gorm.Model
	Id        int64  `gorm:"index"`
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Verification struct {
	gorm.Model
	Id        int64 `gorm:"index"`
	UserId    int64 `gorm:"index"`
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p PostgresRepository) InsertUser(user *model.User) error {
	gormUser := &User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}

	err := p.db.Create(gormUser).Error
	if err != nil {
		return err
	}

	return nil
}

func (p PostgresRepository) GetUserByUsername(username string) (*model.User, error) {
	gormUser := &User{}
	err := p.db.Where("username = ?", username).First(gormUser).Error
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Id:        gormUser.Id,
		Username:  gormUser.Username,
		Password:  gormUser.Password,
		Email:     gormUser.Email,
		CreatedAt: gormUser.CreatedAt,
		UpdatedAt: gormUser.UpdatedAt,
	}

	return user, nil
}

func (p PostgresRepository) GetUserByEmail(email string) (*model.User, error) {
	gormUser := &User{}
	err := p.db.Where("email = ?", email).First(gormUser).Error
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Id:        gormUser.Id,
		Username:  gormUser.Username,
		Password:  gormUser.Password,
		Email:     gormUser.Email,
		CreatedAt: gormUser.CreatedAt,
		UpdatedAt: gormUser.UpdatedAt,
	}

	return user, nil
}

func (p PostgresRepository) GetUserById(id int64) (*model.User, error) {
	gormUser := &User{}
	err := p.db.Where("id = ?", id).First(gormUser).Error
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Id:        gormUser.Id,
		Username:  gormUser.Username,
		Password:  gormUser.Password,
		Email:     gormUser.Email,
		CreatedAt: gormUser.CreatedAt,
		UpdatedAt: gormUser.UpdatedAt,
	}

	return user, nil
}

func (p PostgresRepository) DeleteUserById(id int64) error {
	err := p.db.Delete(&User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (p PostgresRepository) InsertVerification(verification *model.Verification) error {
	gormVerification := &Verification{
		UserId: verification.UserId,
		Code:   verification.Code,
	}

	err := p.db.Create(gormVerification).Error
	if err != nil {
		return err
	}

	return nil
}
