package model

func addSomeData() {
	img := Carousel{ImgPath: "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/The_Mathematicians.jpg", ProductID: 1}
	DB.Create(&img)

	category := Category{CategoryID: 1, CategoryName: "粮油作物", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 2, CategoryName: "毛茶", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 3, CategoryName: "食用菌", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 4, CategoryName: "瓜果蔬菜", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 5, CategoryName: "花卉苗木", Num: 0}
	DB.Create(&category)
	category = Category{CategoryID: 6, CategoryName: "药材", Num: 0}
	DB.Create(&category)

}
