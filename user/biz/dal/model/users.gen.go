// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	UserID                string         `gorm:"column:user_id;primaryKey;type:char(32)" json:"user_id"` // 用户ID（UUID）
	UserName              string         `gorm:"column:user_name;not null;type:varchar(64)" json:"user_name"` // 用户名
	UserPasswordHash      string         `gorm:"column:user_password_hash;not null;type:varchar(255)" json:"user_password_hash"` // 密码哈希
	Email                 string         `gorm:"column:email;not null;type:varchar(128);unique" json:"email"` // 邮箱（唯一）
	PersonalizedSignature string         `gorm:"column:personalized_signature;type:varchar(255);default:null" json:"personalized_signature"` // 个性描述（允许为空）
	Picture               string         `gorm:"column:picture;type:varchar(255);default:null" json:"picture"` // 头像（允许为空）
	CreatedAt             time.Time      `gorm:"column:created_at;type:timestamp(6);default:CURRENT_TIMESTAMP(6)" json:"created_at"` // 创建时间
	UpdatedAt             time.Time      `gorm:"column:updated_at;type:timestamp(6);default:CURRENT_TIMESTAMP(6);autoUpdateTime" json:"updated_at"` // 更新时间
	DeletedAt             gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"` // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
