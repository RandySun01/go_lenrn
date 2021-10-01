package main

import "fmt"

/*
@author RandySun
@create 2021-08-30-8:32
*/

func QueryRowDemo(id int) {
	// 获取单条数据
	sqlStr := "select id, name, age from user where id=?"
	var u User
	// 非常重要:确保QueryRow之后调用Scan方法,否则持有数据的连接不会被释放
	err := Db.QueryRow(sqlStr, id).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d, name:%s, age:%d\n", u.Id, u.Name, u.Age)
}

func QueryMultiRowDemo(id int) {
	sqlStr := "select id, name, age from user where id > ?"
	userList := make([]User, 0, 120)
	//userList := []User{1,"s", 3}

	rows, err := Db.Query(sqlStr, id)
	if err != nil {
		return
	}
	// 非常重要,关闭rows释放持有的数据库连接
	defer rows.Close()
	fmt.Println(rows, 5555555555555)
	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		//fmt.Println(u)
		userList = append(userList, u)
		fmt.Printf("id: %d, name: %s, age: %d\n", u.Id, u.Name, u.Age)

	}
	fmt.Printf("%#v", userList)
}
