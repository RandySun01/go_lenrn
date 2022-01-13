package mysql

import (
	"bluebell/models/user"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

/*
@author RandySun
@create 2022-01-12-8:53
*/

// 把每一步数据库操作封装成函数
// 待logic层调用
func QueryUserByUserName() {

}

const secret = "RandySun"

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil

}

// InsertUser 数据库中插入一条新用户
func InsertUser(user *user.User) (err error) {
	// 对密码加密
	password := encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	createUser, err := db.Exec(sqlStr, user.UserId, user.Username, password)
	fmt.Println(createUser)

	return
}

// 对密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))

}
