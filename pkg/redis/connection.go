package redis

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func Connect() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Connected to Redis")
}

// func Initialize() {
// 	_, err := Client.FlushAll(Ctx).Result()
// 	if err != nil {
// 		slog.Error("Failed to initialize Redis")
// 	}
// 	slog.Info("Successfully initialized Redis")
// }
