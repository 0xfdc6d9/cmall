package model

import (
	"github.com/jinzhu/gorm"
)

// Address 收货地址模型
type Address struct {
	gorm.Model
	UserID  uint
	Name    string
	Phone   string
	Address string
}
