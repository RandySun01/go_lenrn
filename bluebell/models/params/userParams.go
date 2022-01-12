package params

/*
@author RandySun
@create 2022-01-12-21:48
*/

// 定义请求的参数结构体
type UserParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
