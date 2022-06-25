package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// SearchProjectsService 搜索商品的服务
type SearchProjectsService struct {
	Search string `form:"search" json:"search"`
}

// Show 搜索商品
func (service *SearchProjectsService) Show() serializer.Response {
	projects := []model.Project{}
	code := e.SUCCESS

	err := model.DB.Where("name LIKE ?", "%"+service.Search+"%").Find(&projects).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	projects1 := []model.Project{}
	err = model.DB.Where("info LIKE ?", "%"+service.Search+"%").Find(&projects1).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	projects = append(projects, projects1...)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProjects(projects),
	}
}
