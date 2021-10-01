package main

import "fmt"

/*
@author RandySun
@create 2021-08-30-23:12
*/
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Println("SQL:", sqlStr)
	var u User
	err := Db.QueryRow(sqlStr).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Println("exec failed, err:", err)
		return
	}
	fmt.Println("user: %#v", u)

}
