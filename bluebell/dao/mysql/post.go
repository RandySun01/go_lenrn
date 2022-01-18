package mysql

import "bluebell/models/modelPost"

/*
@author RandySun
@create 2022-01-18-9:11
*/

// CreatePost 创建
func CreatePost(p *modelPost.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.Id, p.Title, p.Content, p.AuthorId, p.CommunityId)
	return err
}

// GetPostDetailById 获取单条数据
func GetPostDetailById(postId int64) (data *modelPost.Post, err error) {
	data = new(modelPost.Post)
	sqlStr := `select post_id, title, content, author_id, community_id from post where post_id = ?`
	err = db.Get(data, sqlStr, postId)
	return
}

// GetPostList 获取所有帖子
func GetPostList(page, size int64) (postList []*modelPost.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post limit ?, ?;`
	postList = make([]*modelPost.Post, 0, 2)
	err = db.Select(&postList, sqlStr, (page-1)*size, size)
	return

}
