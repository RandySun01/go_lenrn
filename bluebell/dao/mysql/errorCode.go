package mysql

import "errors"

/*
@author RandySun
@create 2022-01-16-17:05
*/
var (
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorUserNotExist    = errors.New("用户不存在, 请注册")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorInvalidId       = errors.New("无效的Id")
)
