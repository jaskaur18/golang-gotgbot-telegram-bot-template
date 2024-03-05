package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/constant"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Session struct {
	r          *redis.Client
	TelegramID int64  `json:"-"`
	Language   string `json:"language"`
}

func GetSession(r *redis.Client, telegramID int64) (*Session, error) {
	var session Session
	ctx, cancel := context.WithTimeout(context.Background(), constant.GetRedisTimeOut())
	defer cancel()

	key := strconv.FormatInt(telegramID, 10)

	data, err := r.Get(ctx, key).Result()
	if err != nil {
		// Check if the error is due to key not found in Redis.
		if errors.Is(err, redis.Nil) {
			// Create an empty session and save it to Redis.
			emptySession := &Session{
				r:          r,
				TelegramID: telegramID,
				Language:   "en"}
			err := emptySession.Save() // Save empty session to Redis
			if err != nil {
				return nil, err
			}
			return emptySession, nil
		}
		// Return the error for other cases.
		return &session, err
	}

	err = json.Unmarshal([]byte(data), &session)
	if err != nil {
		// If unmarshal fails, it's assumed the data format has changed.
		log.Printf("Error unmarshalling session. Possibly due to data format change: %v", err)

		// Delete the key from Redis
		errDelete := r.Del(ctx, key).Err()
		if errDelete != nil {
			log.Printf("Error deleting session from Redis: %v", errDelete)
		}

		// Return a fresh session with default language
		return &Session{
			r:          r,
			TelegramID: telegramID,
			Language:   "en",
		}, nil
	}

	session.r = r
	session.TelegramID = telegramID

	return &session, nil
}

func (s *Session) Save() error {
	ctx, cancel := context.WithTimeout(context.Background(), constant.GetRedisTimeOut())
	defer cancel()

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	var Time24Hours = 24
	err = s.r.Set(ctx, strconv.FormatInt(s.TelegramID, 10), data, time.Duration(Time24Hours)*time.Hour).Err()
	if err != nil {
		msg := fmt.Sprintf("error saving session to redis telegramId: %d", s.TelegramID)
		return errors.New(msg)
	}

	return nil
}

func (s *Session) Clear() error {
	ctx, cancel := context.WithTimeout(context.Background(), constant.GetRedisTimeOut())
	defer cancel()

	err := s.r.Del(ctx, strconv.FormatInt(s.TelegramID, 10)).Err()
	if err != nil {
		msg := fmt.Sprintf("error deleting session from redis telegramId: %d", s.TelegramID)
		return errors.New(msg)
	}

	return nil
}
