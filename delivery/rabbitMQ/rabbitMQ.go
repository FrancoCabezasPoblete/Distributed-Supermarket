package rabbitMQ

import (
	"bytes"
	"fmt"

	"github.com/streadway/amqp"
)

func RabbitMQ(registroJSON *bytes.Buffer) {

	// call rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Cola",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
	}

	// publish registro to the queue
	err = ch.Publish(
		"",
		"Cola",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        registroJSON.Bytes(),
		},
	)

	if err != nil {
		fmt.Println(err)
	}

}
