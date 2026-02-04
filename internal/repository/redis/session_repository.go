package redis

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/domain"
	"github.com/redis/go-redis/v9"
)

type SessionRepository struct {
	r *redis.Client
}

func NewSessionRepository(r *redis.Client) *SessionRepository {
	return &SessionRepository{r: r}
}

func (r *SessionRepository) Get(ctx context.Context, telegramID int64) (*domain.Session, error) {
	key := strconv.FormatInt(telegramID, 10)
	data, err := r.r.Get(ctx, key).Result()
	if err == redis.Nil {
		// Default session
		return &domain.Session{TelegramID: telegramID, Language: "en"}, nil
	}
	if err != nil {
		return nil, err
	}

	var session domain.Session
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return nil, err
	}
	session.TelegramID = telegramID
	return &session, nil
}

func (r *SessionRepository) Save(ctx context.Context, session *domain.Session) error {
	key := strconv.FormatInt(session.TelegramID, 10)
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return r.r.Set(ctx, key, data, 24*time.Hour).Err()
}
