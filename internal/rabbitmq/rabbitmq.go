package rabbitmq

import (
	"log"
	"os"
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
		username := os.Getenv("RABBITMQ_USERNAME")
		password := os.Getenv("RABBITMQ_PASSWORD")
		host := os.Getenv("RABBITMQ_HOST")
		port := os.Getenv("RABBITMQ_PORT")
		vhost := os.Getenv("RABBITMQ_VHOST")

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
		true,
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
