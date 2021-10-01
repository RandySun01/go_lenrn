package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"logTransfer/core/es"
	"sync"
)

/*
@author RandySun
@create 2021-09-27-22:21
*/

// 连接kafka
func Init(address []string, topic string) (err error){

	// 创建消费者
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	var wg sync.WaitGroup
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者 OffsetOldest 最老的 最新的读OffsetNewest
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		wg.Add(1)
		fmt.Println("start to con sume ...")
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				// 为了将同步流程异步划,所以将取出的日志数据放到channel中
				fmt.Println("in sarama.partitionconsumer")
				fmt.Println(msg.Topic)
				fmt.Println(msg.Value)
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
				var m1 map[string]interface{}
				err = json.Unmarshal(msg.Value, &m1)
				if err != nil{
					fmt.Println("unmarshal msg fail err:%v\n", err)
					continue
				}
				es.PutLogDta(m1)
			}
		}(pc)

	}
	wg.Wait()
	return nil
}