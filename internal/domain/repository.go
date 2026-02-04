package domain

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByTelegramID(ctx context.Context, telegramID int64) (*User, error)
	Exists(ctx context.Context, telegramID int64) (bool, error)
	UpdateType(ctx context.Context, telegramID int64, userType UserType) error
	List(ctx context.Context) ([]User, error)
}
