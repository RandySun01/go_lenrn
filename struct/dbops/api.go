package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"struct/defs"
)

func GetAllBook()([] *defs.Book, error){
	stmtOut, err := dbConn.Prepare(`select id, name, author from book`)
	var res [] *defs.Book
	rows, err := stmtOut.Query()
	if err != nil{
		return res, err
	}
	for rows.Next(){
		var id int
		var name, author string
		if err := rows.Scan(&id, &name, &author); err!=nil{
			return res, err
		}
		c := &defs.Book{Id: id, Name: name, Author: author}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
