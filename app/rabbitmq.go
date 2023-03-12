package app

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/web"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
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
		publishDelayRequest := web.PublishDelayRequest{}
		errr := json.Unmarshal(msg.Body, &publishDelayRequest)
		helper.PanicIfError(errr)
		receiver, err := helper.GetReceiver(publishDelayRequest.Receiver)
		helper.PanicIfError(err)

		// Defining the map for the platforms
		platforms := map[string]func(){
			"email": func() {
				err := helper.SendEmail(publishDelayRequest.EmailSubject, receiver.Email, publishDelayRequest.EmailMessage)
				helper.PanicIfError(err)
			},
			"telegram": func() {
				helper.SendMessageTelegram(receiver.Telegram, publishDelayRequest.Message)
			},
			"whatsapp": func() {
				helper.SendMessageWhatsapp(receiver.Whatsapp, publishDelayRequest.Message)
			},
		}

		// Executing the commands based on the boolean variables
		var a, b, c = publishDelayRequest.Platform.Email, publishDelayRequest.Platform.Telegram, publishDelayRequest.Platform.Whatsapp
		if a {
			platforms["email"]()
		}
		if b {
			platforms["telegram"]()
		}
		if c {
			platforms["whatsapp"]()
		}
		if a && b && c {
			for _, platform := range platforms {
				platform()
			}
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
