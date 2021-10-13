package util

import (
	redis "gopkg.in/redis.v5"
)

var client *redis.Client

func InitializeRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default db
	})
}
