package models

import (
	"time"
)

// User 用户模型（超过20个字段），增加了性别字段
type User struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	Username          string    `gorm:"size:64;not null;index:idx_username"`       // 用户名
	Gender            string    `gorm:"size:10;not null;index:idx_gender"`         // 性别
	Age               int       `gorm:"not null"`                                  // 年龄
	Email             string    `gorm:"size:128;not null;uniqueIndex:idx_email"`   // 邮箱
	Phone             string    `gorm:"size:20;not null;uniqueIndex:idx_phone"`    // 电话
	Address           string    `gorm:"size:256;not null"`                         // 地址
	Nationality       string    `gorm:"size:64;not null"`                          // 国籍
	Occupation        string    `gorm:"size:64;not null"`                          // 职业
	MaritalStatus     string    `gorm:"size:16;not null;index:idx_marital_status"` // 婚姻状况
	Education         string    `gorm:"size:64;not null"`                          // 教育程度
	Hobby             string    `gorm:"size:128"`                                  // 爱好
	Income            float64   `gorm:"not null"`                                  // 收入
	RegistrationDate  time.Time `gorm:"not null;index:idx_registration_date"`      // 注册日期
	LastLogin         time.Time `gorm:"not null"`                                  // 最后登录时间
	LoyaltyPoints     int       `gorm:"not null"`                                  // 忠诚积分
	PreferredLanguage string    `gorm:"size:32;not null"`                          // 首选语言
	Currency          string    `gorm:"size:8;not null"`                           // 币种
	Timezone          string    `gorm:"size:32;not null"`                          // 时区
	Status            string    `gorm:"size:16;not null;index:idx_status"`         // 用户状态
	CreatedAt         time.Time // 创建时间
	UpdatedAt         time.Time // 更新时间
}

// TableName 指定数据库中的表名
func (User) TableName() string {
	return "users"
}
