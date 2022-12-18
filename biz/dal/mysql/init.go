package mysql

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// Init DB 缓存单例
func Init(dsn string, maxConn int, maxOpen int) (DB *gorm.DB) {
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	//设置连接池
	sqlDB, _ := db.DB()
	//空闲
	sqlDB.SetMaxIdleConns(maxConn)
	//打开
	sqlDB.SetMaxOpenConns(maxOpen)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	hlog.Info("mysql init success")
	return
}
