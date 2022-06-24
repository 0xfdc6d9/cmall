package serializer

import (
	"cmall/model"
	"time"
)

type Project struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    int    `json:"category_id"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Raiser        uint   `json:"raiser"`
	RaisedAmount  int    `json:"raiser_amount"`
	TargetAmount  int    `json:"target_amount"`
	RemainingTime int    `json:"remaining_time"` // 剩余天数
}

// BuildProject 序列化众筹项目
func BuildProject(item model.Project) Project {
	return Project{
		ID:            item.ID,
		Name:          item.Name,
		CategoryID:    item.CategoryID,
		Info:          item.Info,
		ImgPath:       item.ImgPath,
		Raiser:        item.Raiser,
		RaisedAmount:  item.RaisedAmount,
		TargetAmount:  item.TargetAmount,
		RemainingTime: int(item.RemainingTime.Sub(time.Now()).Hours() / 24.0),
	}
}

// BuildProjects 序列化商品列表
func BuildProjects(items []model.Project) (projects []Project) {
	for _, item := range items {
		project := BuildProject(item)
		projects = append(projects, project)
	}
	return projects
}
