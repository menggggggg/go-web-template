package user

import "time"

type User struct {
	ID        uint64    `json:"id"`         // 唯一标识
	UserName  string    `json:"user_name"`  // 用户名
	RealName  string    `json:"real_name"`  // 真实姓名
	Password  string    `json:"password"`   // 密码
	Phone     string    `json:"phone"`      // 手机号
	Email     string    `json:"email"`      // 邮箱
	Status    int       `json:"status"`     // 用户状态(1:启用 2:停用)
	Creator   uint64    `json:"creator"`    // 创建者
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

type CreateRequest struct {
	UserName string `json:"user_name" binding:"required"` // 用户名
	RealName string `json:"real_name"`                    // 真实姓名
	Password string `json:"password"  binding:"required"` // 密码
	Phone    string `json:"phone"  binding:"required"`    // 手机号
	Email    string `json:"email" `                       // 邮箱
}

type GetByNameRequest struct {
	Name string `form:"name" binding:"required"`
}

type GetRequest struct {
	UserId string `uri:"userId" binding:"required"`
}

type UpdateRequest struct {
	UserName string `json:"user_name"` // 用户名
	RealName string `json:"real_name"` // 真实姓名
	Password string `json:"password"`  // 密码
	Phone    string `json:"phone"`     // 手机号
	Email    string `json:"email" `    // 邮箱
}

type DeleteRequest struct {
	UserId string `uri:"userId" binding:"required"`
}
