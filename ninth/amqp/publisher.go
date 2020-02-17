package amqp

import (
	"fmt"
	"github.com/streadway/amqp"
)

func PublishMessage() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"router",     // exchange
		"send-message", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			CorrelationId: "123",
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

}


func failOnError(err error, msg string) {
	if err != nil {
		fmt.Errorf("%s: %s", msg, err)
	}
}