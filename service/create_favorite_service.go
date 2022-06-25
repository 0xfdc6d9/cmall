package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// CreateFavoriteService 收藏创建的服务
type CreateFavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProjectID uint `form:"project_id" json:"project_id"`
}

// Create 创建收藏夹
func (service *CreateFavoriteService) Create() serializer.Response {
	var favorite model.Favorite
	code := e.SUCCESS
	model.DB.Where("user_id=? AND project_id=?", service.UserID, service.ProjectID).Find(&favorite)
	if favorite == (model.Favorite{}) {
		favorite = model.Favorite{
			UserID:    service.UserID,
			ProjectID: service.ProjectID,
		}
		if err := model.DB.Create(&favorite).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		code = e.ERROR_EXIST_FAVORITE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
