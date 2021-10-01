package main

import "fmt"

/*
@author RandySun
@create 2021-08-31-8:04
*/

// 查询单条数据示例
func DeleteRowSqlxDemo(id int) {
	sqlStr := "delete from user where id = ?"
	ret, err := DbSqlx.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
