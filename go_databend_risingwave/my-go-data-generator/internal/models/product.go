package models

import (
	"time"
)

// Product 产品模型，包含超过20个字段
type Product struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	ProductName     string    `gorm:"size:128;not null;index:idx_productname"` // 产品名称
	Category        string    `gorm:"size:64;not null;index:idx_category"`     // 分类
	Description     string    `gorm:"type:text"`                               // 产品描述
	Price           float64   `gorm:"not null"`                                // 价格
	Stock           int       `gorm:"not null"`                                // 库存数量
	SKU             string    `gorm:"size:64;not null;uniqueIndex:idx_sku"`    // 库存单位编码
	Manufacturer    string    `gorm:"size:128;not null"`                       // 制造商
	Weight          float64   `gorm:"not null"`                                // 重量
	Dimensions      string    `gorm:"size:64;not null"`                        // 规格尺寸
	Color           string    `gorm:"size:32;not null"`                        // 颜色
	Material        string    `gorm:"size:64;not null"`                        // 材质
	ReleaseDate     time.Time `gorm:"not null"`                                // 发布日期
	WarrantyPeriod  string    `gorm:"size:32"`                                 // 保修期
	CountryOfOrigin string    `gorm:"size:64;not null"`                        // 产地
	Rating          float64   `gorm:"not null;index:idx_rating"`               // 评分
	NumberOfReviews int       `gorm:"not null"`                                // 评论数
	Discount        float64   `gorm:"not null"`                                // 折扣信息
	StockStatus     string    `gorm:"size:32;not null;index:idx_stock_status"` // 库存状态
	Supplier        string    `gorm:"size:128;not null"`                       // 供应商
	CreatedAt       time.Time // 创建时间
	UpdatedAt       time.Time // 更新时间
}

func (Product) TableName() string {
	return "products"
}
