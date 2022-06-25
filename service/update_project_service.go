package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
	"time"
)

// UpdateProjectService 众筹项目更新的服务
type UpdateProjectService struct {
	ID            uint      `form:"id" json:"id"`
	Name          string    `form:"name" json:"name"`
	CategoryID    int       `form:"category_id" json:"category_id"`
	Info          string    `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string    `form:"img_path" json:"img_path"`
	Raiser        uint      `form:"raiser" json:"raiser"`
	TargetAmount  int       `form:"target_amount" json:"target_amount"`
	RemainingTime time.Time `form:"remaining_time" json:"remaining_time"`
}

// Update 更新众筹项目
func (service *UpdateProjectService) Update() serializer.Response {
	// TODO 更新众筹项目的发起人
	project := model.Project{
		Name:          service.Name,
		CategoryID:    service.CategoryID,
		Info:          service.Info,
		ImgPath:       service.ImgPath,
		Raiser:        service.Raiser,
		RaisedAmount:  0,
		TargetAmount:  service.TargetAmount,
		RemainingTime: service.RemainingTime,
	}
	project.ID = service.ID
	code := e.SUCCESS
	err := model.DB.Save(&project).Error
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
