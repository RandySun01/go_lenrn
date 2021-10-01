package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

/*
@author RandySun
@create 2021-08-31-8:01
*/
var DbSqlx *sqlx.DB

func InitDbSqlx() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_mysql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	DbSqlx, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	DbSqlx.SetMaxOpenConns(20)
	DbSqlx.SetMaxIdleConns(10)
	return
}
