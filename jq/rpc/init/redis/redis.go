package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Rdb *redis.Client

func RedisInit() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := Rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	log.Println("redis init success")
}
