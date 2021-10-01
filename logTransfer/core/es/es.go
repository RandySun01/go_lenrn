package es

import (
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"context"
)

/*
@author RandySun
@create 2021-09-27-22:21
*/

// 将日志写入Elasticsearch
type EsClient struct {
	client *elastic.Client // es client
	index string // index 数据库索引
	logDataChan chan interface{} // 接收日志数据的channel
}
var (
	esClient  = &EsClient{}
)
// 连接es
func Init( address, index string, goroutineNum, maxSize int)(err error){
	client, err := elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	esClient.client = client
	esClient.index = index
	esClient.logDataChan = make(chan interface{}, maxSize)
	// 从通道中获取数据


	// 从通道中取数据写入到kafka中去
	for i:=0; i< goroutineNum;i++{
		go sendToEs()
	}


	return nil
}

// 向kafka中写入数据
func sendToEs() {
	for m1 := range esClient.logDataChan{
		//b, err := json.Marshal(m1)
		//if err != nil{
		//	log.Fatalf("marsh1 m1 failed continue, err:%v\n", err)
		//	continue
		//}
		put1, err := esClient.client.Index(). // Index获取所有数据库
			Index(esClient.index).
			BodyJson(m1).
			Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}

}
// 通过首字母大写的函数从包外调用传递信息
func PutLogDta(msg interface{})  {
	esClient.logDataChan <- msg

}