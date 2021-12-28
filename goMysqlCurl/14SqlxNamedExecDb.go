package main

/*
@author RandySun
@create 2021-08-31-8:28
*/


func InsertUserSqlxDemo() (err error) {
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = DbSqlx.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "RandySun2",
			"age":  18,
		})
	return
}
