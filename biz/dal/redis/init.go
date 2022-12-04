package redis

import (
	"github.com/go-redis/redis"
)

// REDIS 缓存客户端单例
var REDIS *redis.Client

func Init(db int, addr string, password string) {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	REDIS = client
}
