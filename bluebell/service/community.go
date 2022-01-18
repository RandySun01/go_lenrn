package service

import (
	"bluebell/dao/mysql"
	"bluebell/models/modelCommunity"
)

/*
@author RandySun
@create 2022-01-16-16:20
*/
func GetCommunityList() ([]*modelCommunity.Community, error) {
	// 查数据库, 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*modelCommunity.CommunityDetail, error) {

	return mysql.GetCommunityDetailById(id)
}
