package helper

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis() {
	opt, err := redis.ParseURL(Env.RedisUri)
	if err != nil {
		log.Fatal("Error parsing redis url: ", err)
	}
	Redis = redis.NewClient(opt)

	err = VerifyRedisConnection()
	if err != nil {
		log.Fatal("Error connecting to Redis: ", err)
	}

	log.Printf("🔗 Redis Connected\n")
}

func VerifyRedisConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), RedisTimeOut)
	defer cancel()
	_, err := Redis.Ping(ctx).Result()

	if err != nil {
		return err
	}

	return nil
}
