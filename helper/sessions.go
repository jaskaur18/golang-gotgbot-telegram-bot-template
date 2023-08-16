package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	TelegramID    int64  `json:"-"`
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	DepositAmount int    `json:"depositAmount"`

	ApiUrlsAndStatus []struct {
		APIUrl    string
		StatusUrl string
	} `json:"apiUrlsAndStatus"`

	Public    bool   `json:"public"`
	ServiceID string `json:"serviceId"`
}

func GetSession(telegramId int64) (*Session, error) {
	var session Session
	ctx, cancel := context.WithTimeout(context.Background(), RedisTimeOut)
	defer cancel()

	key := strconv.FormatInt(telegramId, 10)

	data, err := Redis.Get(ctx, key).Result()
	if err != nil {
		// Check if the error is due to key not found in Redis.
		if err == redis.Nil {
			// Create an empty session and save it to Redis.
			emptySession := &Session{TelegramID: telegramId}
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
		// Handle unmarshal error here.
		err = Redis.Del(ctx, key).Err()
		if err != nil {
			log.Printf("Error deleting session: %v", err)
		}
		return &session, err
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
