package main

import "github.com/electricvortex/dar_golang_course/ninth/amqp"

func main() {

	amqp.PublishMessage()
	amqp.ReceiveMessage()

}
