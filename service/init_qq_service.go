package service

import (
	"cmall/pkg/e"
	"cmall/serializer"
	"fmt"
	"net/url"
	"os"
)

// InitQQService 商品详情的服务
type InitQQService struct {
}

// Init QQ初始化
func (service *InitQQService) Init() serializer.Response {
	code := e.SUCCESS
	responseType := "code"
	path := os.Getenv("QQ_Redirect_URI")
	redirectURI := url.QueryEscape(path)

	loginURL := fmt.Sprintf("https://graph.qq.com/oauth2.0/authorize?response_type=%s&client_id=%s&redirect_uri=%s&state=%s",
		responseType, os.Getenv("QQ_Client_ID"), redirectURI, os.Getenv("QQ_State"))

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   loginURL,
	}
}
