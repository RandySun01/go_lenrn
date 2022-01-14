package mysql

import (
	"bluebell/models/user"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
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

var (
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorUserNotExist    = errors.New("用户不存在, 请注册")
	ErrorInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return

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

func Login(user *user.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	// 用户不存在
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	// 查询数据库错误
	if err != nil {
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword

	}
	return
}
