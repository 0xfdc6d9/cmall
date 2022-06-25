package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"

	"github.com/jinzhu/gorm"
)

// ShowOrderService 订单详情的服务
type ShowOrderService struct {
}

// Show 订单
func (service *ShowOrderService) Show(num string) serializer.Response {
	var order model.Order
	var project model.Project
	code := e.SUCCESS
	//根据id查找order
	if err := model.DB.Where("order_num=?", num).First(&order).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//根据order查找project
	if err := model.DB.Where("id=?", order.ProjectID).First(&project).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ERROR_NOT_EXIST_PRODUCT
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, project),
	}
}
