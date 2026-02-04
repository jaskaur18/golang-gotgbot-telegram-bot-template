package service

import (
	"context"

	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/domain"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/repository/postgres"
)

type UserService struct {
	repo *postgres.UserRepository
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, telegramID int64, firstName string, lastName, username *string) error {
	// Check if exists
	exists, err := s.repo.Exists(ctx, telegramID)
	if err != nil {
		return err
	}
	if exists {
		return nil // idempotent
	}

	// Create
	user := &domain.User{
		TelegramID: telegramID,
		FirstName:  firstName,
		LastName:   lastName,
		Username:   username,
		UserType:   domain.UserTypeUser,
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.List(ctx)
}

func (s *UserService) GetUser(ctx context.Context, telegramID int64) (*domain.User, error) {
	return s.repo.GetByTelegramID(ctx, telegramID)
}

func (s *UserService) UpdateUserType(ctx context.Context, telegramID int64, isAdmin bool) error {
	userType := domain.UserTypeUser
	if isAdmin {
		userType = domain.UserTypeAdmin
	}
	return s.repo.UpdateType(ctx, telegramID, userType)
}
