package params

/*
@author RandySun
@create 2022-01-12-21:48
*/

// 定义请求的参数结构体

// UserParamSignUp 注册
type UserParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// UserParamLogin 登录
type UserParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
