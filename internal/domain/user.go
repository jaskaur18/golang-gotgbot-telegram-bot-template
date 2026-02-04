package domain

import (
	"time"
)

type UserType string

const (
	UserTypeUser  UserType = "USER"
	UserTypeAdmin UserType = "ADMIN"
)

type User struct {
	ID         string
	TelegramID int64
	FirstName  string
	LastName   *string
	Username   *string
	Language   *string
	UserType   UserType
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
