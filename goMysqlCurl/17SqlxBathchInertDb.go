package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

/*
@author RandySun
@create 2021-08-31-8:47
*/

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsersSqlxDemo(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))

	_, err := DbSqlx.Exec(stmt, valueArgs...)
	return err
}

// BatchInsertInUsersSqlxDemo 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

func BatchInsertInUsersSqlxDemo(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err := DbSqlx.Exec(query, args...)
	return err
}

// BatchInsertNamedExecUsersSqlxDemo 使用NamedExec实现批量插入
func BatchInsertNamedExecUsersSqlxDemo(users []*User) error {
	_, err := DbSqlx.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}
