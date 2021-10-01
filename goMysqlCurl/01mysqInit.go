package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var Db *sql.DB

// 定义一个初始化数据库的函数
func InitDb() (err error) {
	// DSN:Data Source Name
	dsn := "root:@tcp(127.0.0.1:3306)/go_mysql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	//  注意!!!,这里不要使用:=,我是是给全局变量赋值,然后在main中使用全局变量db
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接(校验dsn是否正确)
	err = Db.Ping()
	if err != nil {
		return err
	}

	return nil
}
