package model

func addSomeData() {
	img := Carousel{ImgPath: "https://hjm-cmall.oss-cn-hangzhou.aliyuncs.com/The_Mathematicians.jpg", ProductID: 1}
	DB.Create(&img)
}
