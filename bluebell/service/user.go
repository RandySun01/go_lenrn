package service

import (
	"bluebell/dao/mysql"
	"bluebell/models/modelParams"
	"bluebell/models/modelUser"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

/*
@author RandySun
@create 2022-01-12-8:50
*/

// 存放业务逻辑的代码

func SignUp(p *modelParams.UserParamSignUp) (err error) {
	// 判断用户存不住
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		// 数据库查询出错
		return err
	}

	// 生成uid
	userId := snowflake.GenId()

	// 构建user实例
	userInfo := &modelUser.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进入数据库
	return mysql.InsertUser(userInfo)
}

func Login(p *modelParams.UserParamLogin) (userInfo *modelUser.User, err error) {
	// 判断用户存不住
	userInfo = &modelUser.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 数据库中查询是否存在  传递的是一个指针
	if err = mysql.Login(userInfo); err != nil {
		return nil, err
	}

	// 生成JWT的token
	token, err := jwt.GenToken(userInfo.UserId, userInfo.Username)
	if err != nil {
		return nil, err
	}
	userInfo.Token = token
	return

}
