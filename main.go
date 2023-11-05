package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // 队列名称
		false,   // 持久性
		false,   // 自动删除
		false,   // 排他性
		false,   // 无等待
		nil,     // 参数
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello, RabbitMQ!"
	err = ch.Publish(
		"",     // 交换机名称
		q.Name, // 队列名称
		false,  // 强制性
		false,  // 立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")

	log.Printf("Sent: %s", body)
}
