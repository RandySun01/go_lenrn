package mysql

import (
	"bluebell/models/modelPost"
	"bluebell/settings"
	"testing"
)

/*
@author RandySun
@create 2022-01-23-16:17
*/

func init() {
	dbConfig := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "",
		DbName:       "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbConfig)
	if err != nil {
		panic(err)
	}
}
func TestCreatePost(t *testing.T) {
	post := modelPost.Post{
		Id:          10,
		AuthorId:    1,
		CommunityId: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("createpost insert record into mysql failed, err:%#v", err)
	}
	t.Logf("createPost insert record into mysql success")
}
