package conf

import (
	"github.com/joho/godotenv"
	"key-management-service/biz/dal/mysql"
	"key-management-service/biz/dal/redis"
	"os"
	"strconv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// mysql 连接初始化
	dsn := os.Getenv("MYSQL_DSN")
	// 最大空闲连接数
	var maxConn uint64 = 0
	maxConnStr := os.Getenv("MYSQL_MAX_CONN")
	if maxConnStr != "" {
		maxConn, _ = strconv.ParseUint(maxConnStr, 10, 64)
	}

	// 最大连接数（需要大于最大空闲连接数）
	var maxOpen uint64 = 0
	maxOpenStr := os.Getenv("MYSQL_MAX_CONN")
	if maxOpenStr != "" {
		maxOpen, _ = strconv.ParseUint(maxOpenStr, 10, 64)
	}

	mysql.Init(dsn, int(maxConn), int(maxOpen))
	// redis 连接初始化
	var db uint64 = 0
	dbStr := os.Getenv("REDIS_DB")
	if dbStr != "" {
		db, _ = strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	}

	addr := os.Getenv("REDIS_ADDR")
	pw := os.Getenv("REDIS_PW")
	redis.Init(int(db), addr, pw)
}
