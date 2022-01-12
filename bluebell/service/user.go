package service

import "bluebell/dao/mysql"

/*
@author RandySun
@create 2022-01-12-8:50
*/

// 存放业务逻辑的代码

func SignUp() {
	// 判断用户存不住
	mysql.QueryUserByUserName()
	// 生成uid
	// 保存进入数据库

}
