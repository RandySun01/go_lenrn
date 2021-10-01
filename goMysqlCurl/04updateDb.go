package main

import "fmt"

/*
@author RandySun
@create 2021-08-30-22:36
*/

func updateRowDemo(age int, id int) {
	sqlStr := "update user set age=? where id=?"
	ret, err := Db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
	}
	n, err := ret.RowsAffected() // 操作收影响的行
	if err != nil {
		fmt.Printf("get RowsAffected failed: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)

}
