package model

import (
	"github.com/jinzhu/gorm"
)

// ProductInfoImg 商品图片模型
type ProductInfoImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
