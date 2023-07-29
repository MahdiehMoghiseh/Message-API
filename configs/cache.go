package configs

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetFromCache(key string) ([]byte, error) {
	log.Println("using cacheeeee -> get")
	data, err := client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SetToCache(key string, data []byte, duration time.Duration) error {
	log.Println("using cacheee -> set")
	err := client.Set(key, data, duration).Err()
	if err != nil {
		return err
	}
	return nil
}
