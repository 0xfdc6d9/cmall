package serializer

import (
	"cmall/model"
	"time"
)

// Favorite 收藏序列化器
type Favorite struct {
	UserID        uint   `json:"user_id"`
	ProjectID     uint   `json:"id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	CategoryID    int    `json:"category_id"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Raiser        uint   `json:"raiser"`
	RaisedAmount  int    `json:"raiser_amount"`
	TargetAmount  int    `json:"target_amount"`
	RemainingTime int    `json:"remaining_time"` // 剩余天数
}

// BuildFavorite 序列化收藏夹
func BuildFavorite(item1 model.Favorite, item2 model.Project) Favorite {
	return Favorite{
		UserID:        item1.UserID,
		ProjectID:     item1.ProjectID,
		CreatedAt:     item1.CreatedAt.Unix(),
		Name:          item2.Name,
		CategoryID:    item2.CategoryID,
		Info:          item2.Info,
		ImgPath:       item2.ImgPath,
		Raiser:        item2.Raiser,
		RaisedAmount:  item2.RaisedAmount,
		TargetAmount:  item2.TargetAmount,
		RemainingTime: int(item2.RemainingTime.Sub(time.Now()).Hours() / 24.0),
	}
}

// BuildFavorites 序列化收藏夹列表
func BuildFavorites(items []model.Favorite) (favorites []Favorite) {
	for _, item1 := range items {
		item2 := model.Project{}
		err := model.DB.First(&item2, item1.ProjectID).Error
		if err != nil {
			continue
		}
		favorite := BuildFavorite(item1, item2)
		favorites = append(favorites, favorite)
	}
	return favorites
}
