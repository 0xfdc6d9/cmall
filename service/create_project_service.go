package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

type CreateProjectService struct {
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Raiser        uint   `form:"raiser" json:"raiser"`
	RaisedAmount  int    `form:"raiser_amount" json:"raiser_amount"`
	TargetAmount  int    `form:"target_amount" json:"target_amount"`
	RemainingTime int    `form:"remaining_time" json:"remaining_time"`
}

// Create 创建众筹项目
func (service *CreateProjectService) Create() serializer.Response {
	project := model.Project{
		Name:          service.Name,
		CategoryID:    service.CategoryID,
		Info:          service.Info,
		ImgPath:       service.ImgPath,
		Raiser:        service.Raiser,
		RaisedAmount:  service.RaisedAmount,
		TargetAmount:  service.TargetAmount,
		RemainingTime: service.RemainingTime,
	}

	code := e.SUCCESS
	err := model.DB.Create(&project).Error
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
		Data:   serializer.BuildProject(project),
	}
}
