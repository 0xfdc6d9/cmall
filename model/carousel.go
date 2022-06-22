package model

import (
	"github.com/jinzhu/gorm"
)

// Carousel 轮播图模型
type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductID uint
}
