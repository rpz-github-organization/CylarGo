package server

import (
	"go-crud/api"
	"go-crud/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 以下为中间件, 顺序不能改
	// 使用 Session 插件
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	// 使用跨域管理插件
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		// API 生死检查
		v1.GET("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录的才能访问的
		v1.Use(middleware.AuthRequired())
		{
			// User Routing
			v1.GET("user/me", api.UserMe)
			v1.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
