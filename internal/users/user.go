package users

import (
	"gorm.io/gorm"
	"time"
)

type Service struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	Id    string
	Fullname    string
	Email  string
	CreatedAt time.Time
}

type UserService interface {
	GetUser(ID uint) (User, error)
}

// NewService - returns a new comments service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
