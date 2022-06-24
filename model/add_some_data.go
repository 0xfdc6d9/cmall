package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

func addSomeData() {
	// 测试图床链接
	//img := Carousel{ImgPath: "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/The_Mathematicians.jpg", ProductID: 1}
	//DB.Create(&img)

	// 添加类别
	category := Category{CategoryID: 1, CategoryName: "食用油", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 2, CategoryName: "粮食", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 3, CategoryName: "水果", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 4, CategoryName: "菌类", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 5, CategoryName: "蔬菜", Num: 0}
	DB.Create(&category)

	// 添加图片
	// EdibleOil（食用油类）
	project := Project{
		Model:         gorm.Model{},
		Name:          "玉米油",
		CategoryID:    1,
		Info:          "This is corn oil",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/EdibleOil/corn_oil.jpg",
		Raiser:        1,
		RaisedAmount:  1000,
		TargetAmount:  10000,
		RemainingTime: time.Now().AddDate(0, 0, 10),
	}
	DB.Create(&project)
	project = Project{
		Model:         gorm.Model{},
		Name:          "橄榄油",
		CategoryID:    1,
		Info:          "This is olive oil",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/EdibleOil/olive_oil.jpg",
		Raiser:        1,
		RaisedAmount:  1001,
		TargetAmount:  10001,
		RemainingTime: time.Now().AddDate(0, 0, 11),
	}
	DB.Create(&project)

	// Foodstuff（粮食）
	project = Project{
		Model:         gorm.Model{},
		Name:          "玉米",
		CategoryID:    2,
		Info:          "This is corn",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/Foodstuff/corn.jpg",
		Raiser:        1,
		RaisedAmount:  2000,
		TargetAmount:  20000,
		RemainingTime: time.Now().AddDate(0, 0, 20),
	}
	DB.Create(&project)
	project = Project{
		Model:         gorm.Model{},
		Name:          "水稻",
		CategoryID:    2,
		Info:          "This is paddy",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/Foodstuff/paddy.jpg",
		Raiser:        1,
		RaisedAmount:  2001,
		TargetAmount:  20001,
		RemainingTime: time.Now().AddDate(0, 0, 21),
	}
	DB.Create(&project)

	// Fruit（水果）
	project = Project{
		Model:         gorm.Model{},
		Name:          "苹果",
		CategoryID:    3,
		Info:          "This is an apple",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/Fruits/apple.jpg",
		Raiser:        1,
		RaisedAmount:  3000,
		TargetAmount:  30000,
		RemainingTime: time.Now().AddDate(0, 0, 30),
	}
	DB.Create(&project)
	project = Project{
		Model:         gorm.Model{},
		Name:          "香蕉",
		CategoryID:    3,
		Info:          "This is a banana",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/Fruits/banana.jpg",
		Raiser:        1,
		RaisedAmount:  3001,
		TargetAmount:  30001,
		RemainingTime: time.Now().AddDate(0, 0, 31),
	}
	DB.Create(&project)

	// Fungi
	project = Project{
		Model:         gorm.Model{},
		Name:          "木耳",
		CategoryID:    4,
		Info:          "This is edible fungus",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/Fungi/edible_fungus.jpg",
		Raiser:        1,
		RaisedAmount:  4000,
		TargetAmount:  40000,
		RemainingTime: time.Now().AddDate(0, 0, 40),
	}
	DB.Create(&project)
	project = Project{
		Model:         gorm.Model{},
		Name:          "蘑菇",
		CategoryID:    4,
		Info:          "These are mushrooms",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/Fungi/mushrooms.jpg",
		Raiser:        1,
		RaisedAmount:  4001,
		TargetAmount:  40001,
		RemainingTime: time.Now().AddDate(0, 0, 41),
	}
	DB.Create(&project)

	// Vegetable（蔬菜）
	project = Project{
		Model:         gorm.Model{},
		Name:          "娃娃菜",
		CategoryID:    5,
		Info:          "These are baby cabbages",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/Vegetable/baby_cabbage.jpg",
		Raiser:        1,
		RaisedAmount:  5000,
		TargetAmount:  50000,
		RemainingTime: time.Now().AddDate(0, 0, 50),
	}
	DB.Create(&project)
	project = Project{
		Model:         gorm.Model{},
		Name:          "芹菜",
		CategoryID:    5,
		Info:          "This is a celery",
		ImgPath:       "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/Vegetable/celery.jpg",
		Raiser:        1,
		RaisedAmount:  5001,
		TargetAmount:  50001,
		RemainingTime: time.Now().AddDate(0, 0, 51),
	}
	DB.Create(&project)

}
