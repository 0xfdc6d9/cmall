package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// CreateCartService 购物车创建的服务
type CreateCartService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProjectID uint `form:"project_id" json:"project_id"`
}

// Create 创建购物车
func (service *CreateCartService) Create() serializer.Response {
	var project model.Project
	code := e.SUCCESS
	err := model.DB.First(&project, service.ProjectID).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 待添加的商品在数据库中不存在
	if project == (model.Project{}) {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var cart model.Cart
	model.DB.Where("user_id=? AND project_id=?", service.UserID, service.ProjectID).Find(&cart)
	//如果不存在该购物车则创建
	if cart == (model.Cart{}) {
		cart = model.Cart{
			UserID:    service.UserID,
			ProjectID: service.ProjectID,
			Num:       1,
			MaxNum:    1000000,
			Check:     false,
		}

		err = model.DB.Create(&cart).Error
		if err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   serializer.BuildCart(cart, project),
		}
	} else if cart.Num < cart.MaxNum { //如果存在该购物车且num小于maxnum
		cart.Num++
		err = model.DB.Save(&cart).Error
		if err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		return serializer.Response{
			Status: 201,
			Msg:    "商品已在购物车，数量+1",
			Data:   serializer.BuildCart(cart, project),
		}
	} else {
		return serializer.Response{
			Status: 202,
			Msg:    "超过最大上限",
		}
	}
}
