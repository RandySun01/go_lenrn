package main

import "fmt"

/*
@author RandySun
@create 2021-08-30-22:29
*/

func insertRowDemo(name string, age int) {
	sqlStr := "insert into user(name, age) values(?,?)"
	ret, err := Db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed , err:%v\n", err)
		return
	}

	theId, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert Id failed, err: %v\n,", err)
	}
	fmt.Println("insert success, the id is:", theId)

}
