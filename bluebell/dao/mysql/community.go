package mysql

import (
	"bluebell/models/modelCommunity"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

/*
@author RandySun
@create 2022-01-16-16:25
*/

// GetCommunityList 获取社区分类
func GetCommunityList() (data []*modelCommunity.Community, err error) {
	sqlStr := `select community_id, community_name from community;`

	if err = db.Select(&data, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	fmt.Println(data)
	return
}

// GetCommunityDetailById 根据Id查询社区详情
func GetCommunityDetailById(id int64) (communityDetail *modelCommunity.CommunityDetail, err error) {
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?;"
	communityDetail = new(modelCommunity.CommunityDetail)
	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidId
		}
	}
	return
}
