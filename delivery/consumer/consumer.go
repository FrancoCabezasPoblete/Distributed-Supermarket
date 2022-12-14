package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"tarea2/delivery/entity"
	"tarea2/delivery/storage"
)

func main() {

	// call rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	DB := storage.ConnectDatabase()

	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}

	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"Cola",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {

			var registro entity.Registro

			// decode d.Body []byte to entity
			if err := json.NewDecoder(bytes.NewBuffer(d.Body)).Decode(&registro); err != nil {
				log.Fatal("Error")
			}

			// send registro to the database
			storage.AddRegister(registro)

			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()
	<-forever

	DB.Close()

}

func strings(b []byte) {

	panic("unimplemented")

}
