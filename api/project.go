package api

import (
	"cmall/pkg/logging"
	"cmall/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateProject 创建众筹项目
func CreateProject(c *gin.Context) {
	service := service.CreateProjectService{}
	if err := c.ShouldBind(&service); err == nil {
		fmt.Println(service.Name, service.CategoryID, service.Info, service.TargetAmount, service.RemainingTime)
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ListProjects 众筹项目列表接口
func ListProjects(c *gin.Context) {
	service := service.ListProjectsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// ShowProject 商品详情接口
func ShowProject(c *gin.Context) {
	service := service.ShowProjectService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}
