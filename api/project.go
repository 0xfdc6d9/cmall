package api

import (
	"cmall/pkg/logging"
	"cmall/service"
	"github.com/gin-gonic/gin"
)

// CreateProject 创建众筹项目
func CreateProject(c *gin.Context) {
	service := service.CreateProjectService{}
	if err := c.ShouldBind(&service); err == nil {
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

// ShowProject 众筹项目详情接口
func ShowProject(c *gin.Context) {
	service := service.ShowProjectService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// DeleteProject 删除众筹项目的接口
func DeleteProject(c *gin.Context) {
	service := service.DeleteProjectService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

// UpdateProject 更新众筹项目的接口
func UpdateProject(c *gin.Context) {
	service := service.UpdateProjectService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// SearchProjects 搜索众筹项目的接口
func SearchProjects(c *gin.Context) {
	service := service.SearchProjectsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
