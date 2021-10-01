package main

import (
	"log"
	"logTransfer/core/es"
	"logTransfer/core/kafka"
	"logTransfer/model"
)

/*
@author RandySun
@create 2021-09-27-8:53
*/
import "gopkg.in/ini.v1"
// 1. 从kafka消费日志数据, 写入es中




func main() {
	
	// 1. 加载配置文件
	var cfg = new(model.Config)
	err := ini.MapTo(cfg, "G:\\goproject\\go\\logTransfer\\conf\\logtransfer.ini")
	if err != nil{
		log.Fatalf("load config failed, err: %v\n", err)
		panic(err)
	}

	log.Println("load config success", cfg)
	// 2.连接es
	err = es.Init(cfg.EsConf.Address, cfg.EsConf.Index, cfg.EsConf.GoroutineNum, cfg.EsConf.MaxChanSize)
	if err != nil{
		log.Fatalf("connect to es failed, err: %v\n", err)
	}
	log.Println("Init es success")

	// 3. 连接kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil{
		log.Fatalf("connect to kafaka failed, err: %v\n", err)
	}
	log.Println("Init kafka success")



	// 让程序在这停顿
	select {

	}
}