package serializer

import "go-crud/model"

// User 用户序列化器
type User struct {
	ID        uint   `form:"id" json:"id"`
	UserPhone string `form:"phone" json:"user_phone"`
	Nickname  string `form:"name" json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserPhone: user.UserPhone,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
