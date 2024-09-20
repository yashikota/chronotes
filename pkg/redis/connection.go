package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Redis")
}
