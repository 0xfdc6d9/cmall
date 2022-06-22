package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// SearchProductsService 搜索商品的服务
type SearchProductsService struct {
	Search string `form:"search" json:"search"`
}

// Show 搜索商品
func (service *SearchProductsService) Show() serializer.Response {
	products := []model.Product{}
	code := e.SUCCESS

	err := model.DB.Where("name LIKE ?", "%"+service.Search+"%").Find(&products).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products1 := []model.Product{}
	err = model.DB.Where("info LIKE ?", "%"+service.Search+"%").Find(&products1).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products = append(products, products1...)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
