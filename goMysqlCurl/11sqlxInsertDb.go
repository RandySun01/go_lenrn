package main

import "fmt"

/*
@author RandySun
@create 2021-08-31-8:04
*/

// 查询单条数据示例
func InsertRowSqlxDemo(name string, age int) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := DbSqlx.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)

}
