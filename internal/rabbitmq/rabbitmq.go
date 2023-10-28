package rabbitmq

import (
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	onceConn sync.Once
	onceChan sync.Once
	conn     *amqp.Connection
	channel  *amqp.Channel
	err      error
)

func getRabbitConnection() *amqp.Connection {
	onceConn.Do(func() {
		username := "guest"
		password := "guest"
		host := "localhost"
		port := "5672"
		vhost := "/"

		amqpURI := "amqp://" + username + ":" + password + "@" + host + ":" + port + vhost

		conn, err = amqp.Dial(amqpURI)
		failOnError(err, "Failed to connect to RabbitMQ")
	})
	return conn
}

func GetChannel() *amqp.Channel {
	onceChan.Do(func() {
		connection := getRabbitConnection()
		channel, err = connection.Channel()
		failOnError(err, "Failed to open a channel")
	})
	createQueue(channel)
	return channel
}

func createQueue(ch *amqp.Channel) *amqp.Queue {
	queue, err := ch.QueueDeclare(
		"meow",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	return &queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func CloseChannel() {
	if channel != nil {
		channel.Close()
	}
}

func CloseConnection() {
	if conn != nil {
		conn.Close()
	}
}
