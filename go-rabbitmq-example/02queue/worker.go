package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

/*
@author RandySun
@create 2022-05-28-16:48
*/

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// 1. 尝试连接RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://admin:admin@192.168.12.55:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 2. 接下来，我们创建一个通道，大多数API都是用过该通道操作的。
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	// 声明队列
	q, err := ch.QueueDeclare(
		//"hello",      // name
		"task_queue", // name
		true,         // 声明为持久队列 必须为新队列
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")
	// 获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack false 关闭消息自动确认, true手动确认已经消费完成
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	// 开启循环不断地消费消息
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte(".")) // 数一下有几个.
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second) // 模拟耗时的任务
			log.Printf("Done")
			d.Ack(false) // 手动传递消息确认
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
