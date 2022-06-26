package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

type FalsePayService struct {
	ID        int `json:"id"`
	Num       int `json:"num"`
	ProjectID int `json:"project_id"`
}

func (f *FalsePayService) Upd() serializer.Response {
	var order model.Order
	code := e.SUCCESS
	//根据id查找order
	if err := model.DB.Where("id=?", f.ID).First(&order).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 更新订单类型
	order.Type++
	err := model.DB.Save(&order).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	var project model.Project
	err = model.DB.First(&project, f.ProjectID).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 更新众筹项目已筹集金额
	project.RaisedAmount += f.Num
	err = model.DB.Save(&project).Error
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
