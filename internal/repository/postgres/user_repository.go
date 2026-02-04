package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/db"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/domain"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(conn *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		q: db.New(conn),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	params := db.CreateUserParams{
		TelegramID: &user.TelegramID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Username:   user.Username,
		Language:   user.Language,
		UserType:   db.Usertype(user.UserType),
	}

	id, err := r.q.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (r *UserRepository) GetByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	u, err := r.q.GetUserByTelegramID(ctx, &telegramID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:         u.ID,
		TelegramID: *u.TelegramID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Username:   u.Username,
		Language:   u.Language,
		UserType:   domain.UserType(u.UserType),
		CreatedAt:  u.CreatedAt.Time,
		UpdatedAt:  u.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) Exists(ctx context.Context, telegramID int64) (bool, error) {
	return r.q.CheckUserExist(ctx, &telegramID)
}

func (r *UserRepository) UpdateType(ctx context.Context, telegramID int64, userType domain.UserType) error {
	return r.q.UpdateUserType(ctx, db.UpdateUserTypeParams{
		TelegramID: &telegramID,
		UserType:   db.Usertype(userType),
	})
}

func (r *UserRepository) List(ctx context.Context) ([]domain.User, error) {
	users, err := r.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.User, len(users))
	for i, u := range users {
		result[i] = domain.User{
			ID:         u.ID,
			TelegramID: *u.TelegramID,
			FirstName:  u.FirstName,
			LastName:   u.LastName,
			Username:   u.Username,
			Language:   u.Language,
			UserType:   domain.UserType(u.UserType),
			CreatedAt:  u.CreatedAt.Time,
			UpdatedAt:  u.UpdatedAt.Time,
		}
	}
	return result, nil
}
