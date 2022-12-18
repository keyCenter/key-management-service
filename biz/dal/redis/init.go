package redis

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

// Init Redis 缓存客户端单例
func Init(db int, addr string, password string) (redisClient *redis.Client) {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		hlog.Info("redis init error %s", err)
		panic(err)
	}

	redisClient = client
	hlog.Info("redis init success")
	return
}
