package mypackage

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consume(topics []string) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "144.76.104.26:19092",
		"group.id":          "odoo-consumer-prod",
		"auto.offset.reset": "latest",
		"sasl.mechanisms":    "PLAIN",
		"sasl.username":      "KaFkApR0dAdM1nUz3r41288f9e",
		"sasl.password":      "c72bFVNEy4d34kgu8EjM8d28yvTK6TKU1ca68631aae8",
		"security.protocol":  "SASL_PLAINTEXT",
		"enable.auto.commit": true,
		"debug":              "consumer",
	})

	if err != nil {
		panic(err)
	}

  consumer.SubscribeTopics(topics, nil)
  rq_channel := GetRQChannel()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
      RabbitMQSender(*rq_channel, "hello", []byte(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
    break
	}

	consumer.Close()
}
