package models

import (
	"time"
)

// Order 订单模型（超过20个字段）
// 注意：UserID、ProductID 为逻辑依赖，不启用真正的外键约束
type Order struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	OrderNumber     string    `gorm:"size:64;not null;uniqueIndex:idx_ordernumber"` // 订单编号
	UserID          uint      `gorm:"not null;index:idx_userid"`                     // 用户ID（逻辑关系）
	ProductID       uint      `gorm:"not null;index:idx_productid"`                  // 产品ID（逻辑关系）
	OrderDate       time.Time `gorm:"not null;index:idx_order_date"`               // 订单日期
	Quantity        int       `gorm:"not null"`                                    // 数量
	TotalAmount     float64   `gorm:"not null"`                                    // 总金额
	PaymentMethod   string    `gorm:"size:32;not null"`                            // 支付方式
	ShippingAddress string    `gorm:"size:256;not null"`                           // 收货地址
	BillingAddress  string    `gorm:"size:256;not null"`                           // 账单地址
	OrderStatus     string    `gorm:"size:32;not null;index:idx_order_status"`     // 订单状态
	DiscountAmount  float64   `gorm:"not null"`                                    // 折扣金额
	TaxAmount       float64   `gorm:"not null"`                                    // 税费
	ShippingCost    float64   `gorm:"not null"`                                    // 运费
	TrackingNumber  string    `gorm:"size:64;index:idx_tracking_number"`           // 物流单号
	DeliveryDate    time.Time `gorm:"index:idx_delivery_date"`                     // 预计送达日期
	ReturnStatus    string    `gorm:"size:32;index:idx_return_status"`             // 退货状态
	CustomerNote    string    `gorm:"type:text"`                                   // 客户备注
	InternalNote    string    `gorm:"type:text"`                                   // 内部备注
	IsGift          bool      `gorm:"not null"`                                    // 是否礼物
	GiftMessage     string    `gorm:"type:text"`                                   // 礼物留言
	ExtraInfo       string    `gorm:"type:text"`                                   // 额外信息
	CreatedAt       time.Time // 创建时间
	UpdatedAt       time.Time // 更新时间
}

// TableName 指定数据库中的表名
func (Order) TableName() string {
	return "orders"
}