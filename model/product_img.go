package model

import (
	"github.com/jinzhu/gorm"
)

// ProductImg 商品图片模型
type ProductImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
