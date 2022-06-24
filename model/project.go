package model

import (
	"cmall/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Project struct {
	gorm.Model

	// 项目名称
	Name string

	// 类别标号
	CategoryID int

	// 项目简介
	Info string `gorm:"size:1000"`

	// 图片路径
	ImgPath string

	// 项目发起者
	Raiser uint

	// 已筹集金额
	RaisedAmount int

	// 目标金额
	TargetAmount int

	// 剩余天数
	RemainingTime int
}

// View 获取点击数
func (project *Project) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProjectViewKey(project.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 增加众筹项目点击数
func (project *Project) AddView() {
	// 增加众筹项目点击数
	cache.RedisClient.Incr(cache.ProjectViewKey(project.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(project.ID)))
}
