package redis

import (
	"log"
)

func Initialize() {
	_, err := Client.FlushAll(Ctx).Result()
	if err != nil {
		log.Println("Failed to initialize Redis")
	}
	log.Println("Successfully initialized Redis")
}
