package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

type Session struct {
	TelegramID int64  `json:"-"`
	Language   string `json:"language"`
}

func GetSession(telegramId int64) (*Session, error) {
	var session Session
	ctx, cancel := context.WithTimeout(context.Background(), RedisTimeOut)
	defer cancel()

	key := fmt.Sprintf("%d", telegramId)

	data, err := Redis.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error getting session: %v", err)
		// Check if the error is due to key not found in Redis.
		if err == redis.Nil {
			// Create an empty session and save it to Redis.
			emptySession := &Session{TelegramID: telegramId, Language: "en"}
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
		errDelete := Redis.Del(ctx, key).Err()
		if errDelete != nil {
			log.Printf("Error deleting session from Redis: %v", errDelete)
		}

		// Return a fresh session with default language
		return &Session{TelegramID: telegramId, Language: "en"}, nil
	}

	session.TelegramID = telegramId

	return &session, nil
}

func (s *Session) Save() error {
	ctx, cancel := context.WithTimeout(context.Background(), RedisTimeOut)
	defer cancel()

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = Redis.Set(ctx, strconv.FormatInt(s.TelegramID, 10), data, 0).Err()
	if err != nil {
		msg := fmt.Sprintf("error saving session to redis telegramId: %d", s.TelegramID)
		return errors.New(msg)
	}

	return nil
}

func (s *Session) Clear() error {
	ctx, cancel := context.WithTimeout(context.Background(), RedisTimeOut)
	defer cancel()

	err := Redis.Del(ctx, strconv.FormatInt(s.TelegramID, 10)).Err()
	if err != nil {
		msg := fmt.Sprintf("error deleting session from redis telegramId: %d", s.TelegramID)
		return errors.New(msg)
	}

	return nil
}
