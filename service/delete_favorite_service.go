package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// DeleteFavoriteService 删除收藏的服务
type DeleteFavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProjectID uint `form:"project_id" json:"project_id"`
}

// Delete 删除收藏
func (service *DeleteFavoriteService) Delete() serializer.Response {
	var favorite model.Favorite
	code := e.SUCCESS

	err := model.DB.Where("user_id=? AND project_id=?", service.UserID, service.ProjectID).Find(&favorite).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&favorite).Error
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
