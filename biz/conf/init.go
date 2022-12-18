package conf

import (
	_ "github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/logger/logrus"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
	"key-management-service/biz/core"
	"key-management-service/biz/dal/mysql"
	"key-management-service/biz/dal/redis"
	"os"
	"path"
	"strconv"
	"time"
)

// osEnvConfig 环境变量中读取的参数，用于配置初始化
type osEnvConfig struct {
	mysqlDsn     string
	mysqlMaxConn int
	mysqlMaxOpen int

	redisDb   int
	redisAddr string
	redisPw   string

	host string

	logFilePath string
	logFileName string
	// A file can be up to *M
	logMaxSize int
	// Save up to * files at the same time.
	logMaxBackups int
	// A file can exist for a maximum of * days.
	logMaxAge int
}

func envInit() (envConfig *osEnvConfig) {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	envConfig = &osEnvConfig{}

	// logger 配置
	logFilePath := os.Getenv("LOG_FILE_PATH")
	if logFilePath == "" {
		dir := "."
		logFilePath = dir + "/logs/"
	}
	envConfig.logFilePath = logFilePath

	logFileName := os.Getenv("LOG_FILE_NAME")
	if logFileName == "" {
		// Set filename to date
		logFileName = time.Now().Format("2006-01-02")
	}
	envConfig.logFileName = logFileName

	var logMaxSize uint64 = 20
	logMaxSizeStr := os.Getenv("LOG_MAX_SIZE")
	if logMaxSizeStr != "" {
		logMaxSize, _ = strconv.ParseUint(logMaxSizeStr, 10, 64)
	}
	envConfig.logMaxSize = int(logMaxSize)

	var logMaxBackups uint64 = 5
	logMaxBackupsStr := os.Getenv("LOG_MAX_BACKUPS")
	if logMaxBackupsStr != "" {
		logMaxBackups, _ = strconv.ParseUint(logMaxBackupsStr, 10, 64)
	}
	envConfig.logMaxBackups = int(logMaxBackups)

	var logMaxAge uint64 = 5
	logMaxAgeStr := os.Getenv("LOG_MAX_AGE")
	if logMaxAgeStr != "" {
		logMaxAge, _ = strconv.ParseUint(logMaxAgeStr, 10, 64)
	}
	envConfig.logMaxAge = int(logMaxAge)

	// mysql 配置
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		panic("MYSQL_DSN is empty")
	}
	envConfig.mysqlDsn = dsn
	// 最大空闲连接数
	var maxConn uint64 = 0
	maxConnStr := os.Getenv("MYSQL_MAX_CONN")
	if maxConnStr != "" {
		maxConn, _ = strconv.ParseUint(maxConnStr, 10, 64)
	}
	envConfig.mysqlMaxConn = int(maxConn)

	// 最大连接数（需要大于最大空闲连接数）
	var maxOpen uint64 = 0
	maxOpenStr := os.Getenv("MYSQL_MAX_OPEN")
	if maxOpenStr != "" {
		maxOpen, _ = strconv.ParseUint(maxOpenStr, 10, 64)
	}
	envConfig.mysqlMaxOpen = int(maxOpen)

	// redis 配置
	var db uint64 = 0
	dbStr := os.Getenv("REDIS_DB")
	if dbStr != "" {
		db, _ = strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	}
	envConfig.redisDb = int(db)
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		panic("REDIS_ADDR is empty")
	}
	pw := os.Getenv("REDIS_PW")
	if addr == "" {
		panic("REDIS_PW is empty")
	}
	envConfig.redisAddr = addr
	envConfig.redisPw = pw

	return
}

// Init 初始化配置项
func Init() {

	// 环境变量初始化
	envConfig := envInit()

	// logger init
	logInit(envConfig)
	// mysql 连接初始化
	core.GlobalConfig.DbClient = mysql.Init(envConfig.mysqlDsn, envConfig.mysqlMaxConn, envConfig.mysqlMaxOpen)
	// redis 连接初始化
	core.GlobalConfig.RedisClient = redis.Init(envConfig.redisDb, envConfig.redisAddr, envConfig.redisPw)
}

func logInit(envConfig *osEnvConfig) {
	if err := os.MkdirAll(envConfig.logFilePath, 0o777); err != nil {
		panic(err.Error())
		return
	}

	logFileName := envConfig.logFileName + ".log"
	fileName := path.Join(envConfig.logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err.Error())
			return
		}
	}

	// For logrus detailed settings, please refer to https://github.com/hertz-contrib/logger/tree/main/logrus and https://github.com/sirupsen/logrus
	hertzLog := logrus.NewLogger()
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    envConfig.logMaxSize,
		MaxBackups: envConfig.logMaxBackups,
		MaxAge:     envConfig.logMaxAge,
		Compress:   false, // Compress with gzip.
	}
	hertzLog.SetOutput(lumberjackLogger)
	hertzLog.SetLevel(hlog.LevelDebug)
	hlog.SetLogger(hertzLog)
}
