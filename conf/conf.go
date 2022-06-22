package conf

import (
	"cmall/cache"
	"cmall/model"
	"cmall/pkg/logging"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load("my.env")

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logging.Info(err)
		panic(err)
	}

	fmt.Println(os.Getenv("MYSQL_DSN"))
	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()

}
