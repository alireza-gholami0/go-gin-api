package config

import (
	"fmt"
	"github.com/streadway/amqp"
)

const (
	LoqQueueName = "log_queue"
)

var connection *amqp.Connection

func RabbitConnection() *amqp.Connection {
	var err error
	if connection == nil {
		connection, err = amqp.Dial("amqp://guest:guest@localhost:5672")
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected to RabbitMQ instance")
	}
	return connection
}

func GetRabbitChannel() (*amqp.Channel, error) {
	var channel *amqp.Channel
	var err error
	channel, err = connection.Channel()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		return nil, err
	}
	_, err = channel.QueueDeclare(
		LoqQueueName, // name
		true,         // durable
		false,        // auto delete
		false,        // exclusive
		false,        // no wait
		nil,          // args
	)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		return nil, err
	}
	return channel, nil
}
