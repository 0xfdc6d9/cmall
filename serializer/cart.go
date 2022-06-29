package serializer

import (
	"cmall/model"
	"time"
)

// Cart 项目预投序列化器
type Cart struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProjectID     uint   `json:"project_id"`
	CreatedAt     int64  `json:"created_at"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	Check         bool   `json:"check"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	Raiser        uint   `json:"raiser"`
	RaisedAmount  int    `json:"raiser_amount"`
	TargetAmount  int    `json:"target_amount"`
	RemainingTime int    `json:"remaining_time"` // 剩余天数
}

// BuildCart 序列化项目预投
func BuildCart(item1 model.Cart, item2 model.Project) Cart {
	return Cart{
		ID:            item1.ID,
		UserID:        item1.UserID,
		ProjectID:     item1.ProjectID,
		CreatedAt:     item1.CreatedAt.Unix(),
		Num:           item1.Num,
		MaxNum:        item1.MaxNum,
		Check:         item1.Check,
		Name:          item2.Name,
		ImgPath:       item2.ImgPath,
		Raiser:        item2.Raiser,
		RaisedAmount:  item2.RaisedAmount,
		TargetAmount:  item2.TargetAmount,
		RemainingTime: int(item2.RemainingTime.Sub(time.Now()).Hours() / 24.0),
	}
}

// BuildCarts 序列化项目预投列表
func BuildCarts(items []model.Cart) (carts []Cart) {
	for _, item1 := range items {
		item2 := model.Project{}
		err := model.DB.First(&item2, item1.ProjectID).Error
		if err != nil {
			continue
		}
		cart := BuildCart(item1, item2)
		carts = append(carts, cart)
	}
	return carts
}
