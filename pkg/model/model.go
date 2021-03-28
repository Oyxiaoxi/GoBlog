package model

import (
	"goblog/pkg/logger"

	"gorm.io/gorm"
	// gormlogger "gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	config := mysql.New(mysql.Config{
		DSN: "homestead:secret@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{})
	// DB, err = gorm.Open(config, &gorm.Config{
    //     Logger: gormlogger.Default.LogMode(gormlogger.Info),
    // })
	logger.LogError(err)

	return DB
}
