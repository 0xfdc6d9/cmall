package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// CreateAddressService 收货地址创建的服务
type CreateAddressService struct {
	UserID  uint   `form:"user_id" json:"user_id"`
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

// Create 创建收货地址信息
func (service *CreateAddressService) Create() serializer.Response {
	var address model.Address
	code := e.SUCCESS
	address = model.Address{
		UserID:  service.UserID,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}

	err := model.DB.Create(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var addresses []model.Address
	err = model.DB.Where("user_id=?", service.UserID).Order("created_at desc").Find(&addresses).Error
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
		Data:   serializer.BuildAddresses(addresses),
	}
}
