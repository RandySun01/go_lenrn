package service

import (
	"bluebell/dao/mysql"
	"bluebell/models/params"
	"bluebell/models/user"
	"bluebell/pkg/snowflake"
	"errors"
)

/*
@author RandySun
@create 2022-01-12-8:50
*/

// 存放业务逻辑的代码

func SignUp(p *params.UserParamSignUp) (err error) {
	// 判断用户存不住
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		// 数据库查询出错
		return err
	}
	if exist {
		//  用户已经存在
		return errors.New("用户已经存在")
	}
	// 生成uid
	userId := snowflake.GenId()

	// 构建user实例
	userInfo := &user.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进入数据库
	return mysql.InsertUser(userInfo)
}

func Login(p *params.UserParamLogin) (err error) {
	// 判断用户存不住
	userInfo := &user.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 数据库中查询是否存在
	return mysql.Login(userInfo)
}
