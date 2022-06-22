package model

import (
	"github.com/jinzhu/gorm"
)

// Favorite 收藏夹模型
type Favorite struct {
	gorm.Model
	UserID    uint
	ProductID uint
}
