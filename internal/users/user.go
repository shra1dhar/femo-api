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
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserService interface {
	GetUser(ID uint) (User, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetUser(ID uint) (User, error) {
	var user User
	if result := s.DB.First(&user, ID); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}