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
	failOnError(err, "Falha ao conectar com o Rabbit")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Falha ao abir um canal")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TreasuryDirect.AccountId.Reader", //name
		false,                             //durable
		false,                             //delete when unused
		false,                             //exclusive
		false,                             //no-wait
		nil,                               //arguments
	)
	failOnError(err, "Falha ao declarar a fila")
	err = ch.Publish(
		"",     //exchange
		q.Name, //routing key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	failOnError(err, "Falha ao publicar a mensagem")
}
