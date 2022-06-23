package model

import (
	"cmall/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Project struct {
	gorm.Model
	Name       string
	CategoryID int //类别标号
	Title      string
	Info       string `gorm:"size:1000"`
	ImgPath    string
	Raiser     string // 项目发起者
	Progress   uint   // 项目进度
}

// View 获取点击数
func (project *Project) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProjectViewKey(project.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}
