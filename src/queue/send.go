package queue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendQueue(body string) {
	fmt.Println("Processando fila")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Fail to connect with RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Fail to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TreasuryDirect.AccountId.Reader", //name
		false,                             //durable
		false,                             //delete when unused
		false,                             //exclusive
		false,                             //no-wait
		nil,                               //arguments
	)
	failOnError(err, "Fail to declare a queue")
	err = ch.Publish(
		"",     //exchange
		q.Name, //routing key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")
}
