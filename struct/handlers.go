package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"struct/dbops"
)

func GetBook(c *gin.Context){
	books, err := dbops.GetAllBook()
	if err != nil{
		fmt.Println("获取参数失败")
	}
	for i, ele := range books{
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"title":"我是测试", "ce":books})
}
