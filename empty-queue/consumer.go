package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var (
	uri   = flag.String("uri", "amqp://guest:guest@localhost:5672//IPROC0001", "AMQP URI")
	queue = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
)

func init() {
	flag.Parse()
}

func main() {
	err := emptyQueue(*uri, *queue)
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("shutting down")
}

func emptyQueue(amqpURI, queueName string) error {
	var conn *amqp.Connection
	var channel *amqp.Channel
	var err error
	var ok bool
	var msg amqp.Delivery

	log.Printf("dialing %q", amqpURI)
	conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}
	log.Printf("got Connection, getting Channel")
	channel, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}
	for {
		if msg, ok, err = channel.Get(
			queueName,   // name of the queue
			false,       // autoAck
		); err != nil {
			return fmt.Errorf("Queue Get: %s", err)
		}
		if ok != true {
			log.Printf("Queue empty")
			break
		}
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(msg.Body),
			msg.DeliveryTag,
			msg.Body,
		)
		msg.Ack(false)
	}
	return nil

}

