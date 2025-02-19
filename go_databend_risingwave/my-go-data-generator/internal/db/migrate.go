package db

import (
	"log"

	"gorm.io/gorm"
	"my-go-data-generator/internal/models"
)

// Migrate 执行数据库迁移，自动创建或更新表结构
func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	log.Println("数据库迁移成功")
	// 模型中已定义了一些索引，如需定制可以在此继续追加
}