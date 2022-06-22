package model

import (
	"github.com/jinzhu/gorm"
)

// ProductParamImg 商品图片模型
type ProductParamImg struct {
	gorm.Model
	ProductID uint
	ImgPath   string
}
