package main

import "fmt"

/*
@author RandySun
@create 2021-08-30-22:36
*/

func deleteRowDemo(id int) {
	sqlStr := "delete from user where id=?"
	ret, err := Db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 获取操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected fail, err:%v\n", err)
	}
	fmt.Printf("delete success, affected rows: %v \n", n)

}
