package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// DeleteCartService 项目预投删除的服务
type DeleteCartService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProjectID uint `form:"project_id" json:"project_id"`
}

// Delete 删除项目预投
func (service *DeleteCartService) Delete() serializer.Response {
	var cart model.Cart
	code := e.SUCCESS

	err := model.DB.Where("user_id=? AND project_id=?", service.UserID, service.ProjectID).Find(&cart).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&cart).Error
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
	}
}
