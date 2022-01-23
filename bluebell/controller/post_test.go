package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

/*
@author RandySun
@create 2022-01-23-15:26
*/

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)
	body := `{
    "community_id": 3,
    "title":"titel4",
    "content":"content4"
}`
	// 发送请求
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	// 创建接收响应
	responseWrite := httptest.NewRecorder()
	r.ServeHTTP(responseWrite, request)
	assert.Equal(t, 200, responseWrite.Code)
	// 判断响应的内容是不是按预期返回了,需要登录密码的错误
	// 方法一: 判断响应内容是不是傲寒指定的字符串
	assert.Contains(t, responseWrite.Body.String(), "需要登录")
	// 方法二: 将响应的内容反序列化到Response 然后判断字段与预期是否一致
	res := new(Response)
	if err := json.Unmarshal(responseWrite.Body.Bytes(), res); err != nil {
		t.Fatalf("json.unmarshal ResponseWrite failed, err:%#v\n", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
