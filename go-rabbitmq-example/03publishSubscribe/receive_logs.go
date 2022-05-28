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

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// 声明队列
	q, err := ch.QueueDeclare(
		"",    // 空字符串作为队列名称 会随机生成队列名
		false, // 声明为持久队列 必须为新队列
		false, // delete when unused
		true,  // exclusive 独占队列（当前声明队列的连接关闭后即被删除）
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 绑定队列
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	// 获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack false 关闭消息自动确认, true手动确认已经消费完成
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
			log.Printf(" [x] %s", d.Body)
		}
	}(

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
