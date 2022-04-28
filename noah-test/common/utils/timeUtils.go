package utils

import "time"

/*
Author: RandySun
Date: 2022/3/7 5:53 下午
*/

// GetTimeStamp 获取当前时间戳
func GetTimeStamp() int64 {
	nowTime := time.Now()
	return nowTime.Unix()
}
