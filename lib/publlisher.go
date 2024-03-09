package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/antal0x11/blockchat-cli/dst"
	amqp "github.com/rabbitmq/amqp091-go"
)

func TransactionPublisher(_transaction dst.Transaction) {

	connectionURL := os.Getenv("CONNECTION_URL")

	conn, err := amqp.Dial(connectionURL)
	LogError(err, "# [TransactionProducerExchange]Couldn't establish a connection with RabbitMQ server.")
	defer conn.Close()

	channel, err := conn.Channel()
	LogError(err, "# [TransactionProducerExchange]Couldn't create a channel.")
	defer channel.Close()

	err = channel.ExchangeDeclare(
		"transactions",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	LogError(err, "# [TransactionProducerExchange]Failed to declare the transactions exchange.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(_transaction)
	LogError(err, "# [TransactionProducerExchange]Failed to create json message from transaction.")
	err = channel.PublishWithContext(ctx,
		"transactions",
		"",
		false,
		false, amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	LogError(err, "# [TransactionProducerExchange]Failed to publish transaction.")
	fmt.Println("# [TransactionProducerExchange]Transaction Sent.")
}

func LogError(err error, msg string) {
	if err != nil {
		log.Panicf("%v: %v", msg, err)
	}
}
