package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ShowProjectService 商品详情的服务
type ShowProjectService struct {
}

// Show 商品
func (service *ShowProjectService) Show(id string) serializer.Response {
	var project model.Project
	code := e.SUCCESS
	err := model.DB.First(&project, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//增加点击数
	project.AddView()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProject(project),
	}
}
