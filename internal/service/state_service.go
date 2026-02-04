package service

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type StateService struct {
	r *redis.Client
}

func NewStateService(r *redis.Client) *StateService {
	return &StateService{r: r}
}

func (s *StateService) GetState(ctx context.Context, userID int64) (string, error) {
	key := fmt.Sprintf("fsm:%d:state", userID)
	val, err := s.r.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (s *StateService) SetState(ctx context.Context, userID int64, state string) error {
	key := fmt.Sprintf("fsm:%d:state", userID)
	// TTL of 1 hour for state to avoid locks? Or indefinite?
	// Usually wizards expire. Let's do 24h.
	return s.r.Set(ctx, key, state, 24*time.Hour).Err()
}

func (s *StateService) ClearState(ctx context.Context, userID int64) error {
	key := fmt.Sprintf("fsm:%d:state", userID)
	return s.r.Del(ctx, key).Err()
}

// SetData stores temporary data for the wizard
func (s *StateService) SetData(ctx context.Context, userID int64, key string, value string) error {
	rKey := fmt.Sprintf("fsm:%d:data:%s", userID, key)
	return s.r.Set(ctx, rKey, value, 24*time.Hour).Err()
}

func (s *StateService) GetData(ctx context.Context, userID int64, key string) (string, error) {
	rKey := fmt.Sprintf("fsm:%d:data:%s", userID, key)
	return s.r.Get(ctx, rKey).Result()
}
