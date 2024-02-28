package main

import (
	"fmt"
	"log"
	"time"
	"os"
	"context"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ Connection URL
const amqpURL = "amqp://guest:guest@localhost:5672/"

// Function to publish message to RabbitMQ
func publishToQueue(ctx context.Context,message []byte) error {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err:= conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Initialize a queue
	q, err:= ch.QueueDeclare(
		"task_queue", // queue name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}
	
	err = ch.PublishWithContext(
		ctx,          // context.Context parameter
		"",           // exchange
		q.Name,       // routing key (queue name)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // make the message persistent
			ContentType:  "text/plain",
			Body:         []byte(message),
		},
	)
	if err!=nil{
		return err
	}

	return nil
}

// Function to consume messages from RabbitMQ
func consumeFromQueue(){
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Declare a durable queue named "task_queue"
	q, err := ch.QueueDeclare(
		"task_queue", // queue name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer
		false,  // auto-acknowledge
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Process messages
	for msg := range msgs {
		fmt.Printf("Received Task: %s\n", msg.Body)

		// Simulate task processing
		
    	// fmt.Println("Simulating task processing...")
		time.Sleep(2 * time.Second)

		// Acknowledge the message after processing
		msg.Ack(true)
	}
}

func main(){

	// Start consuming messages in a goroutine
	go consumeFromQueue()

	// Initialize Fiber
	app := fiber.New()

	// Route to enqueue tasks
	app.Post("/enqueue", func(c *fiber.Ctx) error{
		
		// extracting message from the request body
		taskMessage := c.Body()

		ctx := context.Background()

		// Publishing task to the queue
		err := publishToQueue(ctx, taskMessage)
		if err!=nil{
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Success message
		return c.SendString("Task enqueued successfully")
	})

	// Start Fiber App on port 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}