package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 定义一个全局对象db
var Db *sql.DB

// 定义一个初始化数据库的函数
func InitDb() (err error) {
	// DSN:Data Source Name
	dsn := "root:@tcp(127.0.0.1:3306)/go_mysql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	//  注意!!!,这里不要使用:=,我是是给全局变量赋值,然后在main中使用全局变量db
	//  去初始化全局的db对象而不是心声明一个Db变量
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 昨晚检查之后确保db部位nil，在关闭资源
	//defer Db.Close() // 函数直接完就会被调用,要放在外面
	// 尝试与数据库建立连接(校验dsn是否正确)
	err = Db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed err:%#v", err)
		return err
	}

	Db.SetConnMaxLifetime(time.Second * 10)  // 连接存活时间
	Db.SetMaxOpenConns(200)  // 最大连接数
	Db.SetMaxIdleConns(10)  // 最大空闲连接数

	return nil
}

//func main() {
//	if err:=InitDb(); err !=nil{
//		fmt.Printf("connect to db failed, err:#%V", err)
//	}
//	// 做完检查之后确保db部位nil，在关闭资源
//	defer Db.Close() // 函数直接完就会被调用,要放在外面
//	fmt.Println("connect to db success")
//}