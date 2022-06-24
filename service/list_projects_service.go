package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// ListProjectsService 众筹项目列表服务
type ListProjectsService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
	CategoryID uint `form:"category_id" json:"category_id"`
}

// List 视频列表
func (service *ListProjectsService) List() serializer.Response {
	projects := []model.Project{}

	total := 0
	code := e.SUCCESS

	// 设置默认的一页展示项目数量
	if service.Limit == 0 {
		service.Limit = 15
	}

	if service.CategoryID == 0 { // 默认搜索全部
		if err := model.DB.Model(model.Project{}).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&projects).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else { // 搜索具体类别下的项目
		if err := model.DB.Model(model.Project{}).Where("category_id=?", service.CategoryID).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("category_id=?", service.CategoryID).Limit(service.Limit).Offset(service.Start).Find(&projects).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.BuildListResponse(serializer.BuildProjects(projects), uint(total))
}
