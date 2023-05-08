package model

import (
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model 嵌入结构体
	// ID        uint           `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	gorm.Model

	Account  string `json:"account" gorm:"column:account;size:150;comment:账号"`
	Password string `json:"password" gorm:"column:password;size:300;comment:密码"`

	Name   string `json:"name" gorm:"column:name;size:50;comment:名字"`
	Status string `json:"status" gorm:"column:status;size:20;comment:状态"`
}
