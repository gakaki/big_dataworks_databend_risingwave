package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect 建立与 MySQL 数据库的连接，dsn 可从配置或环境变量传入
func Connect(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	return db
}