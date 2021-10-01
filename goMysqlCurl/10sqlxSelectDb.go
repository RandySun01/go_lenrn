package main

import "fmt"

/*
@author RandySun
@create 2021-08-31-8:04
*/

// 查询单条数据示例
func QueryRowSqlxDemo(id int) {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err := DbSqlx.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)

}

// 查询多条数据示例
func QueryMultiRowSqlxDemo(id int) {
	sqlStr := "select id, name, age from user where id > ?"
	var users []User
	err := DbSqlx.Select(&users, sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}
