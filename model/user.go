package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint  `gorm:"primaryKey,autoIncrement"`
	TelegramId int64 `gorm:"uniqueIndex"`
	FirstName  string
	LastName   string
	Username   string
}

// CreateUser creates a new user if it doesn't exist
func CreateUser(user *User) error {
	err := DB.Where(User{TelegramId: user.TelegramId}).FirstOrCreate(&user).Error
	return err
}

// GetUserByTelegramId returns a user by their telegram id
func GetUserByTelegramId(telegramId int64) (*User, error) {
	var user User
	err := DB.Where("telegram_id = ?", telegramId).First(&user).Error
	return &user, err
}

// GetUserById returns a user by their id
func GetUserById(id uint) (*User, error) {
	var user User
	err := DB.First(&user, id).Error
	return &user, err
}

// GetAllUsers returns all users
func GetAllUsers() ([]*User, error) {
	var users []*User
	err := DB.Find(&users).Error
	return users, err
}
