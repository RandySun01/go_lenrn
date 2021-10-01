package main

import "fmt"

/*
@author RandySun
@create 2021-08-31-8:04
*/

// 查询单条数据示例
func UpdateRowSqlxDemo(id, age int) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := DbSqlx.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)

}
