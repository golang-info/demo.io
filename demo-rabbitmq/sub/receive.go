package main

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	url := os.Getenv("AMQP_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672"
	}

	connection, err := amqp.Dial(url)
	if err != nil {
		panic("could not establish connection with RabbitMQ: " + err.Error())
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("could not open RabbitMQ channel: " + err.Error())
	}

	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.QueueBind("test", "#", "events", false, nil)

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}

	msgs, err := channel.Consume("test", "", false, false, false, false, nil)
	if err != nil {
		panic("error consumming the queue: " + err.Error())
	}

	for msg := range msgs {
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

	defer connection.Close()
}