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
