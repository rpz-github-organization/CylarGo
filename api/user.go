package api

import (
	"go-crud/serializer"
	"go-crud/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegisterService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		// gin 上下文绑定 UserRegisterService 没出错，则：
		if user, err := userRegisterService.Register(); err != nil {
			// 若注册动作执行抛出错误，则返回错误信息
			c.JSON(200, err)
		} else {
			// 未报错则返回 新建用户 的序列化信息
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var userLoginService service.UserLoginService
	if err := c.ShouldBind(&userLoginService); err == nil {
		if user, err := userLoginService.Login(); err != nil {
			c.JSON(200, err)
		} else {
			// 设置Session
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.ID)
			_ = s.Save()

			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "登出成功",
	})
}
