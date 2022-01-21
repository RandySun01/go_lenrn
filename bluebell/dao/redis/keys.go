package redis

/*
@author RandySun
@create 2022-01-19-8:51
*/
// redis key
// redis kye 注意使用命名空间的方式,方便查询和拆分

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"  // zset:帖子及发帖的时间
	KeyPostScoreZSet       = "post:score" // zset:帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted" // zset记录用户及投票的类型 参数是post id
	KeyCommunitySetPrefix  = "community:" // set保存每个分区下帖子的id
)

// redis前缀凭借上key
func getRedisKey(key string) string {
	return KeyPrefix + key

}
