package conf

import (
	"go-crud/cache"
	"go-crud/model"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	// 连接数据库 MySQL和Redis
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
