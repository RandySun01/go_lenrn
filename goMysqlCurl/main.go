package main

import "fmt"

/*
@author RandySun
@create 2021-08-30-8:41
*/
func main() {
	//database/sql
	//err := InitDb() // 调用输出数据库的函数
	//if err != nil{
	//	fmt.Printf("init db failed, err: %v\n", err)
	//	return
	//}
	//fmt.Printf("init db success db%v\n", Db)

	//查询单挑数据
	//QueryRowDemo(1)
	// 查询多条数据
	//QueryMultiRowDemo(1)
	// 插入数据
	//insertRowDemo("RandySun", 18)
	//updateRowDemo(19, 5)
	//deleteRowDemo(5)
	//预处理查询示例
	//prepareQueryDemo(3)
	//prepareInserDemo("RandySun", 18)
	// sql注入
	//sqlInjectDemo("xxx' or 1=1#")
	//sqlInjectDemo("xxx' union select * from user #")
	//sqlInjectDemo("xxx' and (select count(*) from user) <4 #")
	// 事务操作
	//transactionDemo()

	// sqlx使用
	err := InitDbSqlx()
	if err != nil {
		fmt.Printf("init db failed, err: %v\n", err)
		return
	}
	fmt.Printf("init db success db%v\n", DbSqlx)
	// 查询单条数据示例
	//QueryRowSqlxDemo(1)
	// 查询多条数据示例
	//QueryMultiRowSqlxDemo(1)
	// 插入数据
	//InsertRowSqlxDemo("RandySunSqlx", 18)
	// 更新数据
	//UpdateRowSqlxDemo(7, 20)
	// 删除数据
	//DeleteRowSqlxDemo(7)
	//InsertUserSqlxDemo()
	//NamedQuerySqlxDemo()
	//TransactionSqlxDemo()
	//u1 := User{Name: "RandySun1", Age: 18}
	//u2 := User{Name: "RandySun2", Age: 28}
	//u3 := User{Name: "RandySun3", Age: 38}
	//
	//// 方法1
	//users := []*User{&u1, &u2, &u3}
	//err = BatchInsertUsersSqlxDemo(users)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	//}
	//
	//// 方法2
	//users2 := []interface{}{u1, u2, u3}
	//err = BatchInsertInUsersSqlxDemo(users2)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	//}
	//
	//// 方法3
	//users3 := []*User{&u1, &u2, &u3}
	//err = BatchInsertNamedExecUsersSqlxDemo(users3)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	//}
	Ids := []int{3,1,2}
	//userList, err := QueryByIds(Ids)
	//fmt.Println(userList)

	userList, err :=QueryAndOrderByIds(Ids)
	fmt.Println(userList)
}
