package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

/*
@author RandySun
@create 2021-08-31-9:10
*/
// QueryByIds 根据给定ID查询
func QueryByIds(ids []int) (users []User, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT id, name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	fmt.Println(query, args, err)
	query = DbSqlx.Rebind(query) // SELECT name, age FROM user WHERE id IN (?, ?, ?)
	fmt.Println(query)

	err = DbSqlx.Select(&users, query, args...)
	return
}

// QueryAndOrderByIds 按照指定id查询并维护顺序
func QueryAndOrderByIds(ids []int) (users []User, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	fmt.Println(strIDs)
	query, args, err := sqlx.In("SELECT id, name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}
	fmt.Println(query, args, err)

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = DbSqlx.Rebind(query)

	err = DbSqlx.Select(&users, query, args...)
	return
}
