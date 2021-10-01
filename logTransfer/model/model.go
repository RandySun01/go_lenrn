package model

/*
@author RandySun
@create 2021-09-27-8:56
*/

type Config struct {
	KafkaConf `ini:"kafka"`
	EsConf `ini:"es"`
}

// KafkaConf kafka配置
type KafkaConf struct {
	Address string `ini:"address"` // kafka地址
	Topic string `ini:"topic"` // kafka-topic
}

// EsConf es配置
type EsConf struct {
	Address string`ini:"address"` // es地址,要以http://开头
	Index string `ini:"index"` //  es索引
	MaxChanSize int `ini:"max_chan_size"` // 存储日志的信息
	GoroutineNum int `ini:"goroutine_num"` // 往es写数据的goroutine数量
}
