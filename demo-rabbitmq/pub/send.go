package main

import (
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

	message := amqp.Publishing{
		Body: []byte("Hello World!"),
	}

	err = channel.Publish("events", "random-key", false, false, message)
	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}

	_, err = channel.QueueDeclare("test", true, false, false, false, nil)
	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	err = channel.QueueBind("test", "#", "events", false, nil)
	if err !=   nil {
		panic("error binding to the queue: " + err.Error())
	}



}
