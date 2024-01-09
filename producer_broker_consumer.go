/*
This codebase  illustrates  a basic implentation of a producer-consumer model using rabbitMQ in Go,
where the producer sends a series of messages to a queue, and the consumer receives and processes these messages.
The use ofRabbitMQ allows for asynchronuous, decouipled communication between producer of data and consumer of data.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

/*
The `producer` function is responsible for sending messages to the RabbitMQ queue named "hello".
It uses a for loop to send 5 messages with incremental numbers in the body.
*/

func producer(ctx context.Context, ch *amqp.Channel, totalMessages int) {
	for i := 0; i < totalMessages; i += 2 {
		var body = fmt.Sprintf("message %d", i)
		err := ch.PublishWithContext(
			ctx,
			"",
			"hello",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		failOnError(err, "Failed to publish a message")
		time.Sleep(time.Second)
	}
}

/*
The `consumer` function sets up a consumer that listens for messages from the "hello" queue.
It uses the `ch.Consume` method to establish a message consumption channel.
Messages are automatically acknowledged(`true` for auto-acknowledgment) once received .
The `consumer` function runs in a goroutine that continously prints the received messages bodies.
The program waits for messages to arrive until the program is stopped.
*/

func consumer(ch *amqp.Channel, wg *sync.WaitGroup, totalMessages int) {
	defer wg.Done()
	msgs, err := ch.Consume(
		"hello",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var count = 0
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		count++
		if count == totalMessages {
			break
		}
	}

	log.Printf("Processed %d messages", count)
}

/*
The `main` function establishes a connection to the RabbiMQ server, opens a channel and declares the "hello" queue.
Then it launches both the producer and the consumer functions concurrently with the use of goroutines.
*/

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@0.0.0.0:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a  channel")
	defer ch.Close()

	// The `_` is a blank identifier in Go and used when the syntax requires a variable
	// but the program logic does not .
	_, err = ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	var wg sync.WaitGroup // Declares a WaitGroup .
	wg.Add(2)             // Waits for 2 goroutines .

	// Passing context to producer
	ctx := context.Background()
	go producer(ctx, ch, 50)
	go consumer(ch, &wg, 25)

	// Waits for all messages to be consumed/processed .
	wg.Wait()
}
