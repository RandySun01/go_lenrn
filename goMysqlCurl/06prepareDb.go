package main

import "fmt"

/*
@author RandySun
@create 2021-08-30-22:56
*/

// 预处理查询示例
func prepareQueryDemo(id int) {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("preare failed, err: %v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d, name: %s, age:%d", u.Id, u.Name, u.Age)

	}
}

// 预处理插入示例
func prepareInserDemo(name string, age int) {
	sqlStr := "insert into user(name, age) values(?, ?)"
	stmt, err := Db.Prepare(sqlStr)
	if err != err {
		fmt.Printf("prepare failed, err:%v\n", err)

	}
	defer stmt.Close()
	ret, err := stmt.Exec(name, age)
	if err != nil {
		fmt.Printf("insert failed, err: %d\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed err:%d", err)
	}
	fmt.Printf("insert sucesss n:%d", n)

}
