package model

import (
	"github.com/jinzhu/gorm"
)

// Cart 订单模型
type Cart struct {
	gorm.Model
	UserID    uint
	ProjectID uint
	Num       uint
	MaxNum    uint
	Check     bool
}
