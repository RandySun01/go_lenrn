package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("../templates/fileUpload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fileUpload.html", nil)

	})

	// 上传单文件
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("file1") // 从请求中获取参数一样
		// 处理multipart forms提交文件时默认的内存限制是32 MiB

		// 可以通过下面的方式修改
		r.MaxMultipartMemory = 1 << 20 // 8 MiB
		fmt.Println(f.Size, "sdfsd")
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 将读取的文件保存在服务端
			// 拼接保存路径
			//dst := fmt.Sprintf("./%s", f.Filename)
			dst := path.Join("./", f.Filename)
			// 保存文件
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	// 上传多文件
	r.POST("/uploadMany", func(c *gin.Context) {
		// 从请求中读取文件
		fromData, _ := c.MultipartForm()
		files := fromData.File["file1"] // 从请求中获取参数一样
		// 处理multipart forms提交文件时默认的内存限制是32 MiB

		// 可以通过下面的方式修改
		r.MaxMultipartMemory = 1 << 20 // 8 MiB

		for _, fileObj := range files {
			fmt.Printf("fileName:%s", fileObj.Filename)
			dst := path.Join("./", fileObj.Filename)
			// 上传指定目录
			err := c.SaveUploadedFile(fileObj, dst)
			if err != nil {
				// 异常处理
				fmt.Printf("file upload err fileName:%s", fileObj.Filename)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})

	})
	r.Run(":9999")
}
