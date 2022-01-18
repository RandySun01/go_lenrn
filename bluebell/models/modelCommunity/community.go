package modelCommunity

import "time"

/*
@author RandySun
@create 2022-01-16-16:28
*/

type Community struct {
	Id   uint64 `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	Id           uint64    `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
