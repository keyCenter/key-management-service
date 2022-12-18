package core

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// Config 初始化后全局持有的配置
type Config struct {
	RedisClient *redis.Client
	DbClient    *gorm.DB
}

// GlobalConfig 全局持有配置
var GlobalConfig = &Config{}
