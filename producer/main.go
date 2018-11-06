package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// Error message helper
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Environment variable helper
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	broker := getEnv("BROKER", "localhost")
	queue := getEnv("QUEUE", "hello")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:5672/", broker))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf("Message: '%s' sent to Queue: %s", body, q.Name)
	failOnError(err, "Failed to deliver message")
}
