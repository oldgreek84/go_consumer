package mypackage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getConnection() (*amqp.Connection) {
	fmt.Println("Go RabbitMQ Starts")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed connect to RabbitMQ")
	// defer conn.Close()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
  return conn
}

func getChannel(conn amqp.Connection) (*amqp.Channel) {
	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")
	// defer ch.Close()
  return ch
}

func declareQueue(queue_name string, ch amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		queue_name, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	fmt.Println(q)
	return q

}

func GetRQChannel() *amqp.Channel{
  conn := getConnection()
  // defer conn.Close()

  ch := getChannel(*conn)
	// defer ch.Close()
  return ch

}

func RunRMQConsumer() {
  conn := getConnection()
  defer conn.Close()

  ch := getChannel(*conn)
	defer ch.Close()

  q := declareQueue("hello", *ch)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %T", msg.Body)
			log.Printf("Received a message: %s", msg.Body)
      sendMsg(string(msg.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func RabbitMQSender(ch amqp.Channel,queue_name string, body []byte) {

  err := ch.Publish(
		"",     // exchange
		queue_name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

const (
  serverPort = 8000
  eventEndpoint = "headers"
)


func sendMsg(body string) {
  fmt.Printf("\n%T\n", body)

  postBody, _ := json.Marshal(body)
  responseBody := bytes.NewBuffer(postBody)

  requestURL := fmt.Sprintf("http://localhost:%d/%s", serverPort, eventEndpoint)
	// res, err := http.Post(requestURL, "application/json", responseBody)
  res, _ := http.NewRequest("POST", requestURL, responseBody)
  res.Header.Set("X-Custom-Header", "myvalue")
  res.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(res)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		// os.Exit(1)
    panic(err)
	}
  defer resp.Body.Close()

  fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", resp.StatusCode)
	fmt.Printf("client: status code: %d\n", resp.Body)

  bodyData, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(bodyData))
}
