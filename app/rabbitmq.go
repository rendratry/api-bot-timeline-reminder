package app

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/web"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"net/url"
)

func RabbitMqConn() *amqp.Connection {
	conn, err := amqp.Dial("amqp://admin:Adminmyfin123@api.myfin.id:5672/")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ: ", err)
	}

	return conn
}

func ConsummerDelayRabbitMQ() {

	conn, err := amqp.Dial("amqp://admin:Adminmyfin123@api.myfin.id:5672/")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ: ", err)
	}

	defer conn.Close()

	// Buka channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error opening channel: ", err)
	}
	defer ch.Close()

	// Declare exchange
	err = ch.ExchangeDeclare("delayed_exchange", "x-delayed-message", true, false, false, false, map[string]interface{}{"x-delayed-type": "direct"})
	if err != nil {
		log.Fatal("Error declaring exchange: ", err)
	}

	// Declare queue
	q, err := ch.QueueDeclare("delayed_queue", true, false, false, false, map[string]interface{}{})
	if err != nil {
		log.Fatal("Error declaring queue: ", err)
	}

	// Bind queue to exchange with routing key
	err = ch.QueueBind(q.Name, "", "delayed_exchange", false, map[string]interface{}{"x-delay": 1800000})
	if err != nil {
		log.Fatal("Error binding queue: ", err)
	}

	// Consume message from queue
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Error consuming message: ", err)
	}

	// Process message
	for msg := range msgs {
		_, err := http.Get("https://api.telegram.org/bot5836520934:AAGQ_iQvY7Hbm5goVRPbwH-k57p25dA-Gns/SendMessage?chat_id=1133337434&text=" + url.QueryEscape(string(msg.Body)))
		if err != nil {
			panic(err)
		}
		log.Printf("Received message: %s", msg.Body)
		msg.Ack(false)
	}
}

func ConsummerSendEmail() {
	conn, err := amqp.Dial("amqp://admin:Adminmyfin123@api.myfin.id:5672/")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ: ", err)
	}

	defer conn.Close()

	// Buka channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error opening channel: ", err)
	}
	defer ch.Close()

	// Declare queue
	q, err := ch.QueueDeclare("SendEmail", false, false, false, false, map[string]interface{}{})
	if err != nil {
		log.Fatal("Error declaring queue: ", err)
	}

	// Consume message from queue
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Error consuming message: ", err)
	}

	// Process message
	for msg := range msgs {
		helper.PanicIfError(err)
		sendEmailRequest := web.SendEmailRequest{}
		err := json.Unmarshal(msg.Body, &sendEmailRequest)
		helper.PanicIfError(err)
		log.Println(sendEmailRequest)

		errr := helper.SendEmail(sendEmailRequest.Subject, sendEmailRequest.Email, sendEmailRequest.Message)
		if err != nil {
			helper.PanicIfError(errr)
		}
		msg.Ack(false)
	}
}
